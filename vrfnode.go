package VRF_Node

import (
	"context"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cozyxiong/VRF-Node/config"
	"github.com/cozyxiong/VRF-Node/database"
	"github.com/cozyxiong/VRF-Node/synchronizer"
	"github.com/cozyxiong/VRF-Node/synchronizer/node"
)

type VrfNode struct {
	db *database.DB

	synchronizer *synchronizer.Synchronizer

	shutdown context.CancelCauseFunc
	stopped  atomic.Bool
}

func NewVrfNode(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*VrfNode, error) {
	db, err := database.NewDB(ctx, cfg.MasterDB)
	if err != nil {
		log.Error("new database error", "err", err)
		return nil, err
	}

	ethClient, err := node.DialEthClient(ctx, cfg.Chain.ChainRpcUrl)
	if err != nil {
		log.Error("new eth client error", "err", err)
		return nil, err
	}

	sync, err := synchronizer.NewSynchronizer(ethClient, db, cfg, shutdown)
	if err != nil {
		log.Error("new synchronizer error", "err", err)
		return nil, err
	}

	return &VrfNode{
		db:           db,
		synchronizer: sync,
		shutdown:     shutdown,
	}, nil
}

func (node *VrfNode) Start(ctx context.Context) error {
	err := node.synchronizer.Start()
	if err != nil {
		return err
	}

	return nil
}

func (node *VrfNode) Stop(ctx context.Context) error {
	err := node.synchronizer.Close()
	if err != nil {
		return err
	}

	return nil
}

func (node *VrfNode) Stopped() bool {
	return node.stopped.Load()
}
