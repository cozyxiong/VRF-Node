package event

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContractEvent struct {
	GUID            uuid.UUID      `gorm:"primaryKey"`
	BlockHash       common.Hash    `gorm:"serializer:bytes"`
	ContractAddress common.Address `gorm:"serializer:bytes"`
	TransactionHash common.Hash    `gorm:"serializer:bytes"`
	LogIndex        uint64
	EventSignature  common.Hash `gorm:"serializer:bytes"`
	Timestamp       uint64
	RLPLog          *types.Log `gorm:"serializer:rlp;column:rlp_bytes"`
}

type ContractEventsView interface {
}

type ContractEventsDB interface {
	ContractEventsView
	StoreContractEvents([]ContractEvent) error
}

type contractEventsDB struct {
	gorm *gorm.DB
}

func NewContractEventDB(db *gorm.DB) ContractEventsDB {
	return &contractEventsDB{gorm: db}
}

func NewContractEventFromLog(log *types.Log, timestamp uint64) ContractEvent {
	eventSig := common.Hash{}
	if len(log.Topics) > 0 {
		eventSig = log.Topics[0]
	}
	return ContractEvent{
		GUID:            uuid.New(),
		BlockHash:       log.BlockHash,
		TransactionHash: log.TxHash,
		ContractAddress: log.Address,
		EventSignature:  eventSig,
		LogIndex:        uint64(log.Index),
		Timestamp:       timestamp,
		RLPLog:          log,
	}
}

func (db *contractEventsDB) StoreContractEvents(events []ContractEvent) error {
	result := db.gorm.CreateInBatches(&events, len(events))
	return result.Error
}
