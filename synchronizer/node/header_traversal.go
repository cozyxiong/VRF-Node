package node

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/cozyxiong/VRF-Node/common/bigint"
)

type HeaderTraversal struct {
	ethClient EthClient
	chainId   uint

	latestHeader        *types.Header
	lastTraversedHeader *types.Header

	blockConfirmationDepth *big.Int
}

func NewHeaderTraversal(ethClient EthClient, chainId uint, fromHeader *types.Header, confDepth *big.Int) *HeaderTraversal {
	return &HeaderTraversal{
		ethClient:              ethClient,
		chainId:                chainId,
		lastTraversedHeader:    fromHeader,
		blockConfirmationDepth: confDepth,
	}
}

func (h *HeaderTraversal) LatestHeader() *types.Header {
	return h.latestHeader
}

func (h *HeaderTraversal) LastTraversedHeader() *types.Header {
	return h.lastTraversedHeader
}

func (h *HeaderTraversal) NextHeaders(maxSize uint64) ([]types.Header, error) {
	latestHeader, err := h.ethClient.BlockHeaderByNumber(nil)
	if err != nil {
		return nil, fmt.Errorf("error getting latest header: %v", err)
	} else if latestHeader == nil {
		return nil, fmt.Errorf("latest header unreported")
	} else {
		h.latestHeader = latestHeader
	}

	endHeight := new(big.Int).Sub(latestHeader.Number, h.blockConfirmationDepth)
	if endHeight.Sign() < 0 {
		return nil, nil
	}

	nextHeight := bigint.Zero
	if h.lastTraversedHeader != nil {
		cmp := h.lastTraversedHeader.Number.Cmp(endHeight)
		if cmp == 0 {
			return nil, nil
		} else if cmp > 0 {
			return nil, errors.New("the HeaderTraversal's internal state is ahead of the provider")
		}
	} else {
		nextHeight = new(big.Int).Add(h.lastTraversedHeader.Number, bigint.One)
	}

	endHeight = bigint.Clamp(nextHeight, endHeight, maxSize)
	headers, err := h.ethClient.BlockHeaderByRange(nextHeight, endHeight, h.chainId)
	if err != nil {
		return nil, fmt.Errorf("error querying blocks by range: %w", err)
	}

	numHeaders := len(headers)
	if numHeaders == 0 {
		return nil, nil
	} else if h.lastTraversedHeader != nil && headers[0].Hash() != h.lastTraversedHeader.Hash() {
		fmt.Println(h.lastTraversedHeader.Number)
		fmt.Println(headers[0].Number)
		fmt.Println(len(headers))
		return nil, errors.New("the HeaderTraversal and provider have diverged in state")
	}
	h.lastTraversedHeader = &headers[numHeaders-1]
	return headers, nil
}
