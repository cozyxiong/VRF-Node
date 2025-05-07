package synchronizer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cozyxiong/VRF-Node/common/tasks"
	"github.com/cozyxiong/VRF-Node/config"
	"github.com/cozyxiong/VRF-Node/database"
	"github.com/cozyxiong/VRF-Node/database/block"
	"github.com/cozyxiong/VRF-Node/database/event"
	"github.com/cozyxiong/VRF-Node/database/utils"
	"github.com/cozyxiong/VRF-Node/synchronizer/node"
	"github.com/cozyxiong/VRF-Node/synchronizer/retry"
)

type Synchronizer struct {
	ethClient       node.EthClient
	db              *database.DB
	chainCfg        *config.ChainConfig
	headerTraversal *node.HeaderTraversal
	headers         []types.Header
	lastHeader      *types.Header
	ctx             context.Context
	cancel          context.CancelFunc
	tasks           tasks.Group
}

func NewSynchronizer(client node.EthClient, db *database.DB, cfg *config.Config, shutdown context.CancelCauseFunc) (*Synchronizer, error) {
	lastHeader, err := db.Blocks.LastBlockHeader()
	if err != nil {
		log.Error("get last block header error", "err", err)
		return nil, err
	}
	var fromHeader *types.Header
	if lastHeader != nil {
		fromHeader = lastHeader.RLPHeader.Header()
	} else if cfg.Chain.StartBlockNumber > 0 {
		header, err := client.BlockHeaderByNumber(big.NewInt(int64(cfg.Chain.StartBlockNumber)))
		if err != nil {
			log.Error("get header by block number error", "err", err)
			return nil, err
		}
		fromHeader = header
	} else {
		log.Info("block number not exist")
	}
	headerTraversal := node.NewHeaderTraversal(client, cfg.Chain.ChainId, fromHeader, big.NewInt(int64(cfg.Chain.Confirmations)))

	ctx, cancel := context.WithCancel(context.Background())

	return &Synchronizer{
		ethClient:       client,
		db:              db,
		chainCfg:        &cfg.Chain,
		headerTraversal: headerTraversal,
		lastHeader:      fromHeader,
		ctx:             ctx,
		cancel:          cancel,
		tasks: tasks.Group{
			HandleCrit: func(err error) {
				shutdown(fmt.Errorf("critical error in Synchronizer %w", err))
			},
		},
	}, err
}

func (sync *Synchronizer) Start() error {
	ticker := time.NewTicker(sync.chainCfg.SynInterval)
	sync.tasks.Go(func() error {
		for range ticker.C {
			newHeaders, err := sync.headerTraversal.NextHeaders(sync.chainCfg.BlockStep)
			if err != nil {
				log.Error("get new headers error", "err", err)
				continue
			} else if len(newHeaders) == 0 {
				log.Warn("no new headers")
			} else {
				sync.headers = newHeaders
			}

			lastHeader := sync.headerTraversal.LastTraversedHeader()
			if lastHeader != nil {
				log.Info("the last synchronized header is %v", lastHeader)
			}

			err = sync.processHeaders(sync.headers)
			if err == nil {
				sync.headers = nil
			}
		}
		return nil
	})
	return nil
}

func (sync *Synchronizer) processHeaders(headers []types.Header) error {
	if len(headers) == 0 {
		return nil
	}
	firstHeader, lastHeader := headers[0], headers[len(headers)-1]
	log.Info("synchronizing header from", "first", firstHeader, " to last", lastHeader)

	addresses, err := sync.db.Proxies.ProxyContractsAddress()
	if err != nil {
		log.Error("get proxy contract addresses error", "err", err)
		return err
	}

	log.Info("getting contract logs...")
	filterQuery := ethereum.FilterQuery{FromBlock: firstHeader.Number, ToBlock: lastHeader.Number, Addresses: addresses}
	logs, err := sync.ethClient.FilterLogs(filterQuery)
	if err != nil {
		log.Error("filter logs error", "err", err)
		return err
	}
	if logs.ToBlockHeader.Number.Cmp(lastHeader.Number) != 0 {
		return errors.New("block number not match")
	} else if logs.ToBlockHeader.Hash() != lastHeader.Hash() {
		return errors.New("block hash not match")
	}
	if len(logs.Logs) > 0 {
		log.Info("detected logs", "count", len(logs.Logs))
	}

	blockHeaders := make([]block.BlockHeader, 0, len(headers))
	for _, header := range headers {
		if header.Number == nil {
			continue
		}

		blockHeader := block.BlockHeader{
			Number:     header.Number,
			Hash:       header.Hash(),
			ParentHash: header.ParentHash,
			Timestamp:  header.Time,
			RLPHeader:  (*utils.RLPHeader)(&header),
		}
		blockHeaders = append(blockHeaders, blockHeader)
	}

	headerMap := make(map[common.Hash]*types.Header, len(headers))
	for _, header := range headers {
		headerMap[header.Hash()] = &header
	}
	contractEvents := make([]event.ContractEvent, len(logs.Logs))
	for _, log := range logs.Logs {
		if _, ok := headerMap[log.BlockHash]; ok {
			continue
		}
		contractEvent := event.NewContractEventFromLog(&log, headerMap[log.BlockHash].Time)
		contractEvents = append(contractEvents, contractEvent)
	}

	strategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](sync.ctx, 10, strategy, func() (interface{}, error) {
		if err := sync.db.Transaction(func(tx *database.DB) error {
			if err := tx.Blocks.StoreBlockHeaders(blockHeaders); err != nil {
				return err
			}
			if err := tx.Events.StoreContractEvents(contractEvents); err != nil {
				return err
			}
			return nil
		}); err != nil {
			log.Info("store block headers and contract events error", "err", err)
			return nil, fmt.Errorf("store block headers and contract events error", "err", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

func (sync *Synchronizer) Close() error {
	sync.cancel()
	return nil
}
