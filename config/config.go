package config

import (
	"time"

	"github.com/urfave/cli/v2"

	"github.com/cozyxiong/VRF-Node/flags"
)

type Config struct {
	Migrations     string
	Chain          ChainConfig
	MasterDB       DBConfig
	SlaveDB        DBConfig
	SlaveDbEnable  bool
	ApiCacheEnable bool
}

type ChainConfig struct {
	ChainId                   uint
	ChainRpcUrl               string
	StartBlockNumber          uint64
	Confirmations             uint64
	BlockStep                 uint64
	SynInterval               time.Duration
	ParseInterval             time.Duration
	CallInterval              time.Duration
	PrivateKey                string
	VrfContractAddress        string
	VrfFactoryContractAddress string
	CallerAddress             string
	NumConfirmations          uint64
	SafeAbortNonceTooLowCount uint64
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		Migrations: ctx.String(flags.MigrationsFlag.Name),
		Chain: ChainConfig{
			ChainId:                   ctx.Uint(flags.ChainIdFlag.Name),
			ChainRpcUrl:               ctx.String(flags.ChainRpcFlag.Name),
			StartBlockNumber:          ctx.Uint64(flags.StartBlockNumberFlag.Name),
			Confirmations:             ctx.Uint64(flags.ConfirmationsFlag.Name),
			BlockStep:                 ctx.Uint64(flags.BlockStepFlag.Name),
			SynInterval:               ctx.Duration(flags.SynIntervalFlag.Name),
			ParseInterval:             ctx.Duration(flags.ParseIntervalFlag.Name),
			CallInterval:              ctx.Duration(flags.CallIntervalFlag.Name),
			PrivateKey:                ctx.String(flags.PrivateKeyFlag.Name),
			VrfContractAddress:        ctx.String(flags.VrfContractAddressFlag.Name),
			VrfFactoryContractAddress: ctx.String(flags.VrfFactoryContractAddressFlag.Name),
			CallerAddress:             ctx.String(flags.CallerAddressFlag.Name),
			NumConfirmations:          ctx.Uint64(flags.NumConfirmationsFlag.Name),
			SafeAbortNonceTooLowCount: ctx.Uint64(flags.SafeAbortNonceTooLowCountFlag.Name),
		},
		MasterDB: DBConfig{
			Host:     ctx.String(flags.MasterDbHostFlag.Name),
			Port:     ctx.Int(flags.MasterDbPortFlag.Name),
			Name:     ctx.String(flags.MasterDbNameFlag.Name),
			User:     ctx.String(flags.MasterDbUserFlag.Name),
			Password: ctx.String(flags.MasterDbPasswordFlag.Name),
		},
		SlaveDB: DBConfig{
			Host:     ctx.String(flags.SlaveDbHostFlag.Name),
			Port:     ctx.Int(flags.SlaveDbPortFlag.Name),
			Name:     ctx.String(flags.SlaveDbNameFlag.Name),
			User:     ctx.String(flags.SlaveDbUserFlag.Name),
			Password: ctx.String(flags.SlaveDbPasswordFlag.Name),
		},
		SlaveDbEnable: ctx.Bool(flags.SlaveDbEnableFlag.Name),
	}
}
