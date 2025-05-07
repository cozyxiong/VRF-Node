package block

import (
	"errors"
	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	"github.com/cozyxiong/VRF-Node/database/utils"
)

type BlockHeader struct {
	GUID       uuid.UUID   `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','')"`
	Hash       common.Hash `gorm:"serializer:bytes"`
	ParentHash common.Hash `gorm:"serializer:bytes"`
	Number     *big.Int    `gorm:"serializer:u256"`
	Timestamp  uint64
	RLPHeader  *utils.RLPHeader `gorm:"serializer:rlp;column:rlp_bytes"`
}

func (BlockHeader) TableName() string {
	return "block_headers"
}

type BlockHeadersView interface {
	BlockHeaderByHash(common.Hash) (*BlockHeader, error)
	BlockHeaderByNumber(*big.Int) (*BlockHeader, error)
	LastBlockHeader() (*BlockHeader, error)
}

type BlockHeadersDB interface {
	BlockHeadersView
	StoreBlockHeaders([]BlockHeader) error
}

type blockHeadersDB struct {
	gorm *gorm.DB
}

func NewBlockHeadersDB(db *gorm.DB) BlockHeadersDB {
	return &blockHeadersDB{gorm: db}
}

func (db *blockHeadersDB) BlockHeaderByHash(hash common.Hash) (*BlockHeader, error) {
	return db.BlockHeaderWithFilter(BlockHeader{Hash: hash})
}

func (db *blockHeadersDB) BlockHeaderByNumber(number *big.Int) (*BlockHeader, error) {
	return db.BlockHeaderWithFilter(BlockHeader{Number: number})
}

func (db *blockHeadersDB) BlockHeaderWithFilter(header BlockHeader) (*BlockHeader, error) {
	var block BlockHeader
	result := db.gorm.Where(&header).Take(&block)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &block, nil
}

func (db *blockHeadersDB) LastBlockHeader() (*BlockHeader, error) {
	var block BlockHeader
	result := db.gorm.Order("number DESC").Take(&block)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &block, nil
}

func (db *blockHeadersDB) StoreBlockHeaders(headers []BlockHeader) error {
	result := db.gorm.Omit("guid").Create(headers)
	return result.Error
}
