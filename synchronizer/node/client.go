package node

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/cozyxiong/VRF-Node/common/global_const"
	"github.com/cozyxiong/VRF-Node/synchronizer/retry"
)

const (
	defaultDialTimeout    = 5 * time.Second
	defaultDialAttempts   = 5
	defaultRequestTimeout = 100 * time.Second
)

type Logs struct {
	Logs          []types.Log
	ToBlockHeader *types.Header
}

type EthClient interface {
	BlockHeaderByNumber(*big.Int) (*types.Header, error)
	BlockHeaderByHash(common.Hash) (*types.Header, error)
	BlockHeaderByRange(*big.Int, *big.Int, uint) ([]types.Header, error)
	LatestSafeBlockHeader() (*types.Header, error)
	LatestFinalizedBlockHeader() (*types.Header, error)

	TxByHash(common.Hash) (*types.Transaction, error)

	FilterLogs(query ethereum.FilterQuery) (Logs, error)
	StorageHash(common.Address, *big.Int) (common.Hash, error)

	Close()
}

type client struct {
	rpc RPC
}

func DialEthClient(ctx context.Context, rpcUrl string) (EthClient, error) {
	ctx, cancle := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancle()

	strategy := retry.Exponential()
	rpcClient, err := retry.Do(ctx, defaultDialAttempts, strategy, func() (*rpc.Client, error) {
		if !IsURLAvailable(rpcUrl) {
			return nil, fmt.Errorf("address unavailable (%s)", rpcUrl)
		}

		client, err := rpc.DialContext(ctx, rpcUrl)
		if err != nil {
			return nil, fmt.Errorf("dial eth client: %w", err)
		}

		return client, nil
	})
	if err != nil {
		return nil, err
	}
	return &client{rpc: NewRPC(rpcClient)}, nil
}

func (c *client) BlockHeaderByNumber(number *big.Int) (*types.Header, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var header *types.Header
	err := c.rpc.CallContext(ctx, &header, "eth_blockNumber", toBlockNumberArg(number), false)
	if err != nil {
		log.Fatalln("call eth_blockNumber failed:", "err", err)
		return nil, err
	} else if header == nil {
		log.Println("header not found")
		return nil, ethereum.NotFound
	}
	return header, nil
}

func (c *client) BlockHeaderByHash(hash common.Hash) (*types.Header, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var header *types.Header
	err := c.rpc.CallContext(ctx, &header, "eth_getBlockByHash", hash, false)
	if err != nil {
		log.Fatalln("call eth_getBlockByHash failed:", "err", err)
		return nil, err
	} else if header == nil {
		log.Println("header not found")
		return nil, ethereum.NotFound
	}
	if header.Hash() != hash {
		return nil, errors.New("header mismatch")
	}
	return header, nil
}

func (c *client) BlockHeaderByRange(startHeight *big.Int, endHeight *big.Int, chainId uint) ([]types.Header, error) {
	if startHeight.Cmp(endHeight) == 0 {
		header, err := c.BlockHeaderByNumber(startHeight)
		if err != nil {
			return nil, err
		}
		return []types.Header{*header}, nil
	}

	count := new(big.Int).Sub(endHeight, startHeight).Uint64() + 1
	headers := make([]types.Header, count)
	batchElems := make([]rpc.BatchElem, count)

	if chainId != uint(global_const.PolygonChainId) {
		for i := uint64(0); i < count; i++ {
			height := new(big.Int).Add(startHeight, new(big.Int).SetUint64(i))
			batchElems[i] = rpc.BatchElem{Method: "eth_getBlockNumber", Args: []interface{}{toBlockNumberArg(height)}, Result: &headers[i]}
		}
		ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
		defer cancel()

		err := c.rpc.BatchCallContext(ctx, batchElems)
		if err != nil {
			return nil, err
		}
	} else {
		groupSize := 100
		var wg sync.WaitGroup
		numGroups := (int(count)-1)/groupSize + 1
		wg.Add(numGroups)

		for i := 0; i < int(count); i += groupSize {
			start := i
			end := i + groupSize - 1
			if end > int(count) {
				end = int(count) - 1
			}
			go func(start, end int) {
				defer wg.Done()
				for j := start; j <= end; j++ {
					ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
					defer cancel()
					height := new(big.Int).Add(startHeight, new(big.Int).SetUint64(uint64(j)))
					batchElems[j] = rpc.BatchElem{
						Method: "eth_getBlockByNumber",
						Result: new(types.Header),
						Error:  nil,
					}
					header := new(types.Header)
					batchElems[j].Error = c.rpc.CallContext(ctx, header, batchElems[j].Method, toBlockNumberArg(height), false)
					batchElems[j].Result = header
				}
			}(start, end)
		}
		wg.Wait()
	}
	size := 0
	for i, batchElem := range batchElems {
		header, ok := batchElem.Result.(*types.Header)
		if !ok {
			return nil, fmt.Errorf("unable to transform rpc response %v into utils.Header", batchElem.Result)
		}
		headers[i] = *header
		size = size + 1
	}
	headers = headers[:size]
	return headers, nil
}

func (c *client) LatestSafeBlockHeader() (*types.Header, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var header *types.Header
	err := c.rpc.CallContext(ctx, &header, "eth_getBlockNumber", "safe", false)
	if err != nil {
		return nil, err
	} else if header == nil {
		return nil, ethereum.NotFound
	}
	return header, nil
}

func (c *client) LatestFinalizedBlockHeader() (*types.Header, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var header *types.Header
	err := c.rpc.CallContext(ctx, &header, "eth_getBlockByNumber", "finalized", false)
	if err != nil {
		return nil, err
	} else if header == nil {
		return nil, ethereum.NotFound
	}
	return header, nil
}

func (c *client) TxByHash(hash common.Hash) (*types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	var tx *types.Transaction
	err := c.rpc.CallContext(ctx, &tx, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, err
	} else if tx == nil {
		return nil, ethereum.NotFound
	}
	return tx, nil
}

func (c *client) FilterLogs(query ethereum.FilterQuery) (Logs, error) {
	arg, err := toFilterArg(query)
	if err != nil {
		return Logs{}, err
	}

	var logs []types.Log
	var header types.Header

	batchElems := make([]rpc.BatchElem, 2)
	batchElems[0] = rpc.BatchElem{Method: "eth_getBlockByNumber", Args: []interface{}{toBlockNumberArg(query.ToBlock), false}, Result: &header}
	batchElems[1] = rpc.BatchElem{Method: "eth_getLogs", Args: []interface{}{arg}, Result: &logs}

	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()
	err = c.rpc.BatchCallContext(ctx, batchElems)
	if err != nil {
		return Logs{}, err
	}

	if batchElems[0].Error != nil {
		return Logs{}, fmt.Errorf("unable to query for the `FilterQuery#ToBlock` header: %w", batchElems[0].Error)
	}

	if batchElems[1].Error != nil {
		return Logs{}, fmt.Errorf("unable to query logs: %w", batchElems[1].Error)
	}

	return Logs{Logs: logs, ToBlockHeader: &header}, nil
}

// StorageHash 32 Bytes - SHA3 of the StorageRoot. All storage will deliver a MerkleProof starting with this rootHash
func (c *client) StorageHash(address common.Address, blockNumber *big.Int) (common.Hash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	proof := struct{ StorageHash common.Hash }{}

	// ARRAY 32 Bytes - array of storage-keys which should be proofed and included.
	err := c.rpc.CallContext(ctx, &proof, "eth_getProof", address, nil, toBlockNumberArg(blockNumber))
	if err != nil {
		return common.Hash{}, err
	}

	return proof.StorageHash, nil
}

func (c *client) Close() {
	c.rpc.Close()
}

func toBlockNumberArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	if number.Sign() >= 0 {
		return hexutil.EncodeBig(number)
	}
	return rpc.BlockNumber(number.Int64()).String()
}

func toFilterArg(query ethereum.FilterQuery) (interface{}, error) {
	arg := map[string]interface{}{"address": query.Addresses, "topics": query.Topics}
	if query.BlockHash != nil {
		arg["blockHash"] = *query.BlockHash
		if query.FromBlock != nil || query.ToBlock != nil {
			return nil, errors.New("cannot filter by both BlockHash and FromBlock/ToBlock")
		}
	} else {
		if query.FromBlock == nil {
			arg["fromBlock"] = "0x0"
		} else {
			arg["fromBlock"] = toBlockNumberArg(query.FromBlock)
		}
		arg["toBlock"] = toBlockNumberArg(query.ToBlock)
	}
	return arg, nil
}

func IsURLAvailable(urlAddress string) bool {
	u, err := url.Parse(urlAddress)
	if err != nil {
		return false
	}
	addr := u.Host
	if u.Port() != "" {
		switch u.Scheme {
		case "http", "ws":
			addr += ":80"
		case "https", "wss":
			addr += ":443"
		default:
			return true
		}
	}
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return false
	}
	err = conn.Close()
	if err != nil {
		return false
	}
	return true
}

type RPC interface {
	CallContext(ctx context.Context, result any, method string, args ...any) error
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	Close()
}

type rpcClient struct {
	rpc *rpc.Client
}

func NewRPC(client *rpc.Client) RPC {
	return &rpcClient{client}
}

func (c *rpcClient) CallContext(ctx context.Context, result any, method string, args ...any) error {
	err := c.rpc.CallContext(ctx, result, method, args...)
	return err
}

func (c *rpcClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	err := c.rpc.BatchCallContext(ctx, b)
	return err
}

func (c *rpcClient) Close() {
	c.rpc.Close()
}
