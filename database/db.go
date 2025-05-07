package database

import (
	"context"
	"fmt"
	"github.com/cozyxiong/VRF-Node/database/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/cozyxiong/VRF-Node/config"
	"github.com/cozyxiong/VRF-Node/database/block"
	"github.com/cozyxiong/VRF-Node/database/event"
	"github.com/cozyxiong/VRF-Node/synchronizer/retry"
)

type DB struct {
	gorm *gorm.DB

	Blocks  block.BlockHeadersDB
	Events  event.ContractEventsDB
	Proxies service.ProxyContractsDB
}

func NewDB(ctx context.Context, dbConfig config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Name)
	if dbConfig.Port != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConfig.Port)
	}
	if dbConfig.User != "" {
		dsn += fmt.Sprintf(" user=%s", dbConfig.User)
	}
	if dbConfig.Password != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.Password)
	}

	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		CreateBatchSize:        3_000,
	}

	strategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	gorm, err := retry.Do[*gorm.DB](context.Background(), 10, strategy, func() (*gorm.DB, error) {
		gorm, err := gorm.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %v", err)
		}
		return gorm, nil
	})
	if err != nil {
		return nil, err
	}

	return &DB{
		gorm:    gorm,
		Blocks:  block.NewBlockHeadersDB(gorm),
		Events:  event.NewContractEventDB(gorm),
		Proxies: service.NewProxyContractsDB(gorm),
	}, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(tx *gorm.DB) error {
		txDB := &DB{
			gorm:    tx,
			Blocks:  block.NewBlockHeadersDB(tx),
			Events:  event.NewContractEventDB(tx),
			Proxies: service.NewProxyContractsDB(tx),
		}
		return fn(txDB)
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to process migration file: %s", path))
		}
		if info.IsDir() {
			return nil
		}
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("failed to read SQL file: %s", path))
		}
		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("failed to execute SQL script: %s", path))
		}
		return nil
	})
	return err
}
