package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProxyContract struct {
	GUID         uuid.UUID      `gorm:"primaryKey"`
	ProxyAddress common.Address `gorm:"serializer:bytes"`
	Timestamp    uint64
}

type ProxyContractsView interface {
	ProxyContractsAddress() ([]common.Address, error)
}

type ProxyContractsDB interface {
	ProxyContractsView
	StoreProxyContractsAddress([]ProxyContract) error
}

type proxyContractsDB struct {
	gorm *gorm.DB
}

func NewProxyContractsDB(db *gorm.DB) ProxyContractsDB {
	return &proxyContractsDB{gorm: db}
}

func (db *proxyContractsDB) ProxyContractsAddress() ([]common.Address, error) {
	var proxies []ProxyContract
	err := db.gorm.Find(&proxies).Error
	if err != nil {
		return nil, fmt.Errorf("get proxy contracts address failed: %w", err)
	}
	var addresses []common.Address
	for _, proxy := range proxies {
		addresses = append(addresses, proxy.ProxyAddress)
	}
	return addresses, nil
}

func (db *proxyContractsDB) StoreProxyContractsAddress(proxies []ProxyContract) error {
	return db.gorm.CreateInBatches(&proxies, len(proxies)).Error
}
