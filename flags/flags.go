package flags

import (
	"github.com/urfave/cli/v2"
	"time"
)

const envVarPrefix = "VRF_NODE"

func prefixEnvVars(name string) []string { return []string{envVarPrefix + "_" + name} }

var (
	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Usage:   "path to migrations directory",
		EnvVars: prefixEnvVars("MIGRATIONS_DIR"),
		Value:   "./migrations",
	}
	ChainIdFlag = &cli.UintFlag{
		Name:     "chain-id",
		Usage:    "The port of the api",
		EnvVars:  prefixEnvVars("CHAIN_ID"),
		Value:    1,
		Required: true,
	}
	ChainRpcFlag = &cli.StringFlag{
		Name:     "chain-rpc",
		Usage:    "rpc node's url",
		EnvVars:  prefixEnvVars("CHAIN_RPC"),
		Required: true,
	}
	StartBlockNumberFlag = &cli.Uint64Flag{
		Name:    "start-block-number",
		Usage:   "The start block number of the chain",
		EnvVars: prefixEnvVars("START_BLOCK_NUMBER"),
		Value:   0,
	}
	ConfirmationsFlag = &cli.Uint64Flag{
		Name:    "confirmations",
		Usage:   "The confirmations depth of the chain",
		EnvVars: prefixEnvVars("CONFIRMATIONS"),
		Value:   64,
	}
	SynIntervalFlag = &cli.DurationFlag{
		Name:    "syn-loop-interval",
		Usage:   "The interval of synchronization",
		EnvVars: prefixEnvVars("SYN_LOOP_INTERVAL"),
		Value:   time.Second * 5,
	}
	BlockStepFlag = &cli.UintFlag{
		Name:    "block-step",
		Usage:   "Scanner blocks step",
		EnvVars: prefixEnvVars("BLOCKS_STEP"),
		Value:   5,
	}
	ParseIntervalFlag = &cli.DurationFlag{
		Name:    "parse-loop-interval",
		Usage:   "The interval of event parse",
		EnvVars: prefixEnvVars("PARSE_LOOP_INTERVAL"),
		Value:   time.Second * 5,
	}
	CallIntervalFlag = &cli.DurationFlag{
		Name:    "call-loop-interval",
		Usage:   "The interval of contact call",
		EnvVars: prefixEnvVars("CALL_LOOP_INTERVAL"),
		Value:   time.Second * 5,
	}
	PrivateKeyFlag = &cli.StringFlag{
		Name:     "private-key",
		Usage:    "The private key of caller contracts",
		EnvVars:  prefixEnvVars("PRIVATE_KEY"),
		Required: true,
	}
	VrfContractAddressFlag = &cli.StringFlag{
		Name:     "vrf-address",
		Usage:    "The VRF contract address",
		EnvVars:  prefixEnvVars("VRF_ADDRESS"),
		Required: true,
	}
	VrfFactoryContractAddressFlag = &cli.StringFlag{
		Name:     "vrf-factory-address",
		Usage:    "The VRF factory contract address",
		EnvVars:  prefixEnvVars("VRF_FACTORY_ADDRESS"),
		Required: true,
	}
	CallerAddressFlag = &cli.StringFlag{
		Name:     "caller-address",
		Usage:    "The contract caller address for vrf",
		EnvVars:  prefixEnvVars("CALLER_ADDRESS"),
		Required: true,
	}
	NumConfirmationsFlag = &cli.Uint64Flag{
		Name:    "num-confirmations",
		Usage:   "The number of confirmations for waiting after appending a new batch",
		EnvVars: prefixEnvVars("NUM_CONFIRMATIONS"),
		Value:   1,
	}
	SafeAbortNonceTooLowCountFlag = &cli.Uint64Flag{
		Name:    "safe-abort-nonce-too-low-count",
		Usage:   "The number of try in particle nonce to abort the transaction without receiving confirmation",
		EnvVars: prefixEnvVars("SAFE_ABORT_NONCE_TOO_LOW_COUNT"),
		Value:   3,
	}

	MasterDbHostFlag = &cli.StringFlag{
		Name:     "master-db-host",
		Usage:    "The hostname of the master db",
		EnvVars:  prefixEnvVars("MASTER_DB_HOST"),
		Required: true,
	}
	MasterDbPortFlag = &cli.UintFlag{
		Name:     "master-db-port",
		Usage:    "The port of the master db",
		EnvVars:  prefixEnvVars("MASTER_DB_PORT"),
		Required: true,
	}
	MasterDbUserFlag = &cli.StringFlag{
		Name:     "master-db-user",
		Usage:    "The username of the master db",
		EnvVars:  prefixEnvVars("MASTER_DB_USER"),
		Required: true,
	}
	MasterDbPasswordFlag = &cli.StringFlag{
		Name:     "master-db-password",
		Usage:    "The password of the master db",
		EnvVars:  prefixEnvVars("MASTER_DB_PASSWORD"),
		Required: true,
	}
	MasterDbNameFlag = &cli.StringFlag{
		Name:     "master-db-name",
		Usage:    "The name of the master db",
		EnvVars:  prefixEnvVars("MASTER_DB_NAME"),
		Required: true,
	}

	SlaveDbEnableFlag = &cli.BoolFlag{
		Name:     "slave-db-enable",
		Usage:    "enable to use slave db",
		EnvVars:  prefixEnvVars("SLAVE_DB_ENABLE"),
		Required: true,
	}
	SlaveDbHostFlag = &cli.StringFlag{
		Name:     "slave-db-host",
		Usage:    "The hostname of the slave db",
		EnvVars:  prefixEnvVars("SLAVE_DB_HOST"),
		Required: true,
	}
	SlaveDbPortFlag = &cli.UintFlag{
		Name:     "slave-db-port",
		Usage:    "The port of the slave db",
		EnvVars:  prefixEnvVars("SLAVE_DB_PORT"),
		Required: true,
	}
	SlaveDbUserFlag = &cli.StringFlag{
		Name:     "slave-db-user",
		Usage:    "The username of the salve db",
		EnvVars:  prefixEnvVars("SLAVE_DB_USER"),
		Required: true,
	}
	SlaveDbPasswordFlag = &cli.StringFlag{
		Name:     "slave-db-password",
		Usage:    "The password of the slave db",
		EnvVars:  prefixEnvVars("SLAVE_DB_PASSWORD"),
		Required: true,
	}
	SlaveDbNameFlag = &cli.StringFlag{
		Name:     "salve-db-name",
		Usage:    "The name of the salve db",
		EnvVars:  prefixEnvVars("SLAVE_DB_NAME"),
		Required: true,
	}
)

var requiredFlags = []cli.Flag{
	MigrationsFlag,
	ChainIdFlag,
	ChainRpcFlag,
	SynIntervalFlag,
	BlockStepFlag,
	ParseIntervalFlag,
	CallIntervalFlag,
	PrivateKeyFlag,
	VrfContractAddressFlag,
	VrfFactoryContractAddressFlag,
	CallerAddressFlag,
	NumConfirmationsFlag,
	SafeAbortNonceTooLowCountFlag,
	MasterDbHostFlag,
	MasterDbPortFlag,
	MasterDbUserFlag,
	MasterDbPasswordFlag,
	MasterDbNameFlag,
	SlaveDbEnableFlag,
}

var optionalFlags = []cli.Flag{
	StartBlockNumberFlag,
	ConfirmationsFlag,
	SlaveDbHostFlag,
	SlaveDbPortFlag,
	SlaveDbUserFlag,
	SlaveDbPasswordFlag,
	SlaveDbNameFlag,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

var Flags []cli.Flag
