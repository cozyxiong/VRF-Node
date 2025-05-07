// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSApkRegistrySignatureCheckParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistrySignatureCheckParams struct {
	NonSignerPubKeysG1 []BN254G1Point
	ApkG2              BN254G2Point
	Sigma              BN254G1Point
	TotalEthStake      *big.Int
	TotalTokenStake    *big.Int
}

// VRFManagerMetaData contains all meta data concerning the VRFManager contract.
var VRFManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blsRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfillRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.SignatureCheckParams\",\"components\":[{\"name\":\"nonSignerPubKeysG1\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalEthStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalTokenStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRequestStatus\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"isFulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nodeAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_blsRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastRequestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"isFulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_numWords\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWhitelist\",\"inputs\":[{\"name\":\"_nodeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FillRandomWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestSent\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"numWords\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f5ffd5b5061001e61002360201b60201c565b61019e565b5f61003261012160201b60201c565b9050805f0160089054906101000a900460ff161561007c576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff161461011e5767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff6040516101159190610185565b60405180910390a15b50565b5f5f61013161013a60201b60201c565b90508091505090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b5f67ffffffffffffffff82169050919050565b61017f81610163565b82525050565b5f6020820190506101985f830184610176565b92915050565b611756806101ab5f395ff3fe608060405234801561000f575f5ffd5b50600436106100cd575f3560e01c80638da5cb5b1161008a578063c0c53b8b11610064578063c0c53b8b146101e9578063d8a4676f14610205578063f2fde38b14610236578063fc2a88c314610252576100cd565b80638da5cb5b1461019157806393e4b880146101af578063be8e11df146101cd576100cd565b80631b739ef1146100d15780632aa8481f146100ed578063715018a61461010b57806382e215ab14610115578063854cff2f146101455780638796ba8c14610161575b5f5ffd5b6100eb60048036038101906100e69190610d03565b610270565b005b6100f5610391565b6040516101029190610d80565b60405180910390f35b6101136103b5565b005b61012f600480360381019061012a9190610d99565b6103c8565b60405161013c9190610dde565b60405180910390f35b61015f600480360381019061015a9190610e21565b6103ee565b005b61017b60048036038101906101769190610d99565b610438565b6040516101889190610e5b565b60405180910390f35b610199610458565b6040516101a69190610d80565b60405180910390f35b6101b761048d565b6040516101c49190610ecf565b60405180910390f35b6101e760048036038101906101e2919061108e565b6104b2565b005b61020360048036038101906101fe919061113d565b610685565b005b61021f600480360381019061021a9190610d99565b610887565b60405161022d929190611244565b60405180910390f35b610250600480360381019061024b9190610e21565b610917565b005b61025a61099b565b6040516102679190610e5b565b60405180910390f35b6102786109a1565b60405180604001604052805f151581526020015f67ffffffffffffffff8111156102a5576102a4610efc565b5b6040519080825280602002602001820160405280156102d35781602001602082028036833780820191505090505b5081525060025f8481526020019081526020015f205f820151815f015f6101000a81548160ff0219169083151502179055506020820151816001019080519060200190610321929190610c59565b50905050600382908060018154018082558091505060019003905f5260205f20015f9091909190915055816004819055507fe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f01682823060405161038593929190611272565b60405180910390a15050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6103bd6109a1565b6103c65f610a28565b565b6002602052805f5260405f205f91509050805f015f9054906101000a900460ff16905081565b6103f66109a1565b805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60038181548110610447575f80fd5b905f5260205f20015f915090505481565b5f5f610462610af9565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610540576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053790611301565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ab44321b8484846040518463ffffffff1660e01b815260040161059e9392919061159a565b606060405180830381865afa1580156105b9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105dd919061164f565b505060405180604001604052806001151581526020018581525060025f8781526020019081526020015f205f820151815f015f6101000a81548160ff0219169083151502179055506020820151816001019080519060200190610641929190610c59565b509050507ff3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a858560405161067692919061168d565b60405180910390a15050505050565b5f61068e610b20565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff161480156106d65750825b90505f60018367ffffffffffffffff1614801561070957505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610717575080155b1561074e576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550831561079b576001855f0160086101000a81548160ff0219169083151502179055505b6107a488610b33565b865f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550831561087d575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516108749190611707565b60405180910390a15b5050505050505050565b5f606060025f8481526020019081526020015f205f015f9054906101000a900460ff1660025f8581526020019081526020015f206001018080548060200260200160405190810160405280929190818152602001828054801561090757602002820191905f5260205f20905b8154815260200190600101908083116108f3575b5050505050905091509150915091565b61091f6109a1565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361098f575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016109869190610d80565b60405180910390fd5b61099881610a28565b50565b60045481565b6109a9610b47565b73ffffffffffffffffffffffffffffffffffffffff166109c7610458565b73ffffffffffffffffffffffffffffffffffffffff1614610a26576109ea610b47565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610a1d9190610d80565b60405180910390fd5b565b5f610a31610af9565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f5f610b2a610b4e565b90508091505090565b610b3b610b77565b610b4481610bb7565b50565b5f33905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b610b7f610c3b565b610bb5576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610bbf610b77565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610c2f575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610c269190610d80565b60405180910390fd5b610c3881610a28565b50565b5f610c44610b20565b5f0160089054906101000a900460ff16905090565b828054828255905f5260205f20908101928215610c93579160200282015b82811115610c92578251825591602001919060010190610c77565b5b509050610ca09190610ca4565b5090565b5b80821115610cbb575f815f905550600101610ca5565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610ce281610cd0565b8114610cec575f5ffd5b50565b5f81359050610cfd81610cd9565b92915050565b5f5f60408385031215610d1957610d18610cc8565b5b5f610d2685828601610cef565b9250506020610d3785828601610cef565b9150509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d6a82610d41565b9050919050565b610d7a81610d60565b82525050565b5f602082019050610d935f830184610d71565b92915050565b5f60208284031215610dae57610dad610cc8565b5b5f610dbb84828501610cef565b91505092915050565b5f8115159050919050565b610dd881610dc4565b82525050565b5f602082019050610df15f830184610dcf565b92915050565b610e0081610d60565b8114610e0a575f5ffd5b50565b5f81359050610e1b81610df7565b92915050565b5f60208284031215610e3657610e35610cc8565b5b5f610e4384828501610e0d565b91505092915050565b610e5581610cd0565b82525050565b5f602082019050610e6e5f830184610e4c565b92915050565b5f819050919050565b5f610e97610e92610e8d84610d41565b610e74565b610d41565b9050919050565b5f610ea882610e7d565b9050919050565b5f610eb982610e9e565b9050919050565b610ec981610eaf565b82525050565b5f602082019050610ee25f830184610ec0565b92915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610f3282610eec565b810181811067ffffffffffffffff82111715610f5157610f50610efc565b5b80604052505050565b5f610f63610cbf565b9050610f6f8282610f29565b919050565b5f67ffffffffffffffff821115610f8e57610f8d610efc565b5b602082029050602081019050919050565b5f5ffd5b5f610fb5610fb084610f74565b610f5a565b90508083825260208201905060208402830185811115610fd857610fd7610f9f565b5b835b818110156110015780610fed8882610cef565b845260208401935050602081019050610fda565b5050509392505050565b5f82601f83011261101f5761101e610ee8565b5b813561102f848260208601610fa3565b91505092915050565b5f5ffd5b5f610120828403121561105257611051611038565b5b81905092915050565b5f819050919050565b61106d8161105b565b8114611077575f5ffd5b50565b5f8135905061108881611064565b92915050565b5f5f5f5f5f60a086880312156110a7576110a6610cc8565b5b5f6110b488828901610cef565b955050602086013567ffffffffffffffff8111156110d5576110d4610ccc565b5b6110e18882890161100b565b94505060406110f288828901610cef565b935050606086013567ffffffffffffffff81111561111357611112610ccc565b5b61111f8882890161103c565b92505060806111308882890161107a565b9150509295509295909350565b5f5f5f6060848603121561115457611153610cc8565b5b5f61116186828701610e0d565b935050602061117286828701610e0d565b925050604061118386828701610e0d565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6111bf81610cd0565b82525050565b5f6111d083836111b6565b60208301905092915050565b5f602082019050919050565b5f6111f28261118d565b6111fc8185611197565b9350611207836111a7565b805f5b8381101561123757815161121e88826111c5565b9750611229836111dc565b92505060018101905061120a565b5085935050505092915050565b5f6040820190506112575f830185610dcf565b818103602083015261126981846111e8565b90509392505050565b5f6060820190506112855f830186610e4c565b6112926020830185610e4c565b61129f6040830184610d71565b949350505050565b5f82825260208201905092915050565b7f5652462e6f6e6c7957686974654c6973740000000000000000000000000000005f82015250565b5f6112eb6011836112a7565b91506112f6826112b7565b602082019050919050565b5f6020820190508181035f830152611318816112df565b9050919050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f8335600160200384360303811261134757611346611327565b5b83810192508235915060208301925067ffffffffffffffff82111561136f5761136e61131f565b5b60408202360383131561138557611384611323565b5b509250929050565b5f82825260208201905092915050565b5f819050919050565b5f6113b46020840184610cef565b905092915050565b604082016113cc5f8301836113a6565b6113d85f8501826111b6565b506113e660208301836113a6565b6113f360208501826111b6565b50505050565b5f61140483836113bc565b60408301905092915050565b5f82905092915050565b5f604082019050919050565b5f611431838561138d565b935061143c8261139d565b805f5b85811015611474576114518284611410565b61145b88826113f9565b97506114668361141a565b92505060018101905061143f565b5085925050509392505050565b5f82905092915050565b5f82905092915050565b82818337505050565b6114aa60408383611495565b5050565b608082016114be5f83018361148b565b6114ca5f85018261149e565b506114d8604083018361148b565b6114e5604085018261149e565b50505050565b5f61012083016114fd5f84018461132b565b8583035f87015261150f838284611426565b925050506115206020840184611481565b61152d60208601826114ae565b5061153b60a0840184611410565b61154860a08601826113bc565b5061155660e08401846113a6565b61156360e08601826111b6565b506115726101008401846113a6565b6115806101008601826111b6565b508091505092915050565b6115948161105b565b82525050565b5f6060820190506115ad5f830186610e4c565b81810360208301526115bf81856114eb565b90506115ce604083018461158b565b949350505050565b5f5ffd5b5f815190506115e881610cd9565b92915050565b5f60408284031215611603576116026115d6565b5b61160d6040610f5a565b90505f61161c848285016115da565b5f83015250602061162f848285016115da565b60208301525092915050565b5f8151905061164981611064565b92915050565b5f5f6060838503121561166557611664610cc8565b5b5f611672858286016115ee565b92505060406116838582860161163b565b9150509250929050565b5f6040820190506116a05f830185610e4c565b81810360208301526116b281846111e8565b90509392505050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f6116f16116ec6116e7846116bb565b610e74565b6116c4565b9050919050565b611701816116d7565b82525050565b5f60208201905061171a5f8301846116f8565b9291505056fea26469706673582212207c2ece28dc4bc901d26030ec65b9cc96b41a779dd5ea48d580836bdeecfeeca564736f6c634300081d0033",
}

// VRFManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFManagerMetaData.ABI instead.
var VRFManagerABI = VRFManagerMetaData.ABI

// VRFManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VRFManagerMetaData.Bin instead.
var VRFManagerBin = VRFManagerMetaData.Bin

// DeployVRFManager deploys a new Ethereum contract, binding an instance of VRFManager to it.
func DeployVRFManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRFManager, error) {
	parsed, err := VRFManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFManager{VRFManagerCaller: VRFManagerCaller{contract: contract}, VRFManagerTransactor: VRFManagerTransactor{contract: contract}, VRFManagerFilterer: VRFManagerFilterer{contract: contract}}, nil
}

// VRFManager is an auto generated Go binding around an Ethereum contract.
type VRFManager struct {
	VRFManagerCaller     // Read-only binding to the contract
	VRFManagerTransactor // Write-only binding to the contract
	VRFManagerFilterer   // Log filterer for contract events
}

// VRFManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFManagerSession struct {
	Contract     *VRFManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFManagerCallerSession struct {
	Contract *VRFManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// VRFManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFManagerTransactorSession struct {
	Contract     *VRFManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VRFManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFManagerRaw struct {
	Contract *VRFManager // Generic contract binding to access the raw methods on
}

// VRFManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFManagerCallerRaw struct {
	Contract *VRFManagerCaller // Generic read-only contract binding to access the raw methods on
}

// VRFManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFManagerTransactorRaw struct {
	Contract *VRFManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFManager creates a new instance of VRFManager, bound to a specific deployed contract.
func NewVRFManager(address common.Address, backend bind.ContractBackend) (*VRFManager, error) {
	contract, err := bindVRFManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFManager{VRFManagerCaller: VRFManagerCaller{contract: contract}, VRFManagerTransactor: VRFManagerTransactor{contract: contract}, VRFManagerFilterer: VRFManagerFilterer{contract: contract}}, nil
}

// NewVRFManagerCaller creates a new read-only instance of VRFManager, bound to a specific deployed contract.
func NewVRFManagerCaller(address common.Address, caller bind.ContractCaller) (*VRFManagerCaller, error) {
	contract, err := bindVRFManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFManagerCaller{contract: contract}, nil
}

// NewVRFManagerTransactor creates a new write-only instance of VRFManager, bound to a specific deployed contract.
func NewVRFManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFManagerTransactor, error) {
	contract, err := bindVRFManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFManagerTransactor{contract: contract}, nil
}

// NewVRFManagerFilterer creates a new log filterer instance of VRFManager, bound to a specific deployed contract.
func NewVRFManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFManagerFilterer, error) {
	contract, err := bindVRFManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFManagerFilterer{contract: contract}, nil
}

// bindVRFManager binds a generic wrapper to an already deployed contract.
func bindVRFManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFManager *VRFManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFManager.Contract.VRFManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFManager *VRFManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFManager.Contract.VRFManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFManager *VRFManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFManager.Contract.VRFManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFManager *VRFManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFManager *VRFManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFManager *VRFManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFManager.Contract.contract.Transact(opts, method, params...)
}

// BlsRegistry is a free data retrieval call binding the contract method 0x93e4b880.
//
// Solidity: function blsRegistry() view returns(address)
func (_VRFManager *VRFManagerCaller) BlsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFManager.contract.Call(opts, &out, "blsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsRegistry is a free data retrieval call binding the contract method 0x93e4b880.
//
// Solidity: function blsRegistry() view returns(address)
func (_VRFManager *VRFManagerSession) BlsRegistry() (common.Address, error) {
	return _VRFManager.Contract.BlsRegistry(&_VRFManager.CallOpts)
}

// BlsRegistry is a free data retrieval call binding the contract method 0x93e4b880.
//
// Solidity: function blsRegistry() view returns(address)
func (_VRFManager *VRFManagerCallerSession) BlsRegistry() (common.Address, error) {
	return _VRFManager.Contract.BlsRegistry(&_VRFManager.CallOpts)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool isFulfilled, uint256[] randomWords)
func (_VRFManager *VRFManagerCaller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (struct {
	IsFulfilled bool
	RandomWords []*big.Int
}, error) {
	var out []interface{}
	err := _VRFManager.contract.Call(opts, &out, "getRequestStatus", _requestId)

	outstruct := new(struct {
		IsFulfilled bool
		RandomWords []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsFulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RandomWords = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool isFulfilled, uint256[] randomWords)
func (_VRFManager *VRFManagerSession) GetRequestStatus(_requestId *big.Int) (struct {
	IsFulfilled bool
	RandomWords []*big.Int
}, error) {
	return _VRFManager.Contract.GetRequestStatus(&_VRFManager.CallOpts, _requestId)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool isFulfilled, uint256[] randomWords)
func (_VRFManager *VRFManagerCallerSession) GetRequestStatus(_requestId *big.Int) (struct {
	IsFulfilled bool
	RandomWords []*big.Int
}, error) {
	return _VRFManager.Contract.GetRequestStatus(&_VRFManager.CallOpts, _requestId)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VRFManager *VRFManagerCaller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFManager.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VRFManager *VRFManagerSession) LastRequestId() (*big.Int, error) {
	return _VRFManager.Contract.LastRequestId(&_VRFManager.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_VRFManager *VRFManagerCallerSession) LastRequestId() (*big.Int, error) {
	return _VRFManager.Contract.LastRequestId(&_VRFManager.CallOpts)
}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_VRFManager *VRFManagerCaller) NodeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFManager.contract.Call(opts, &out, "nodeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_VRFManager *VRFManagerSession) NodeAddress() (common.Address, error) {
	return _VRFManager.Contract.NodeAddress(&_VRFManager.CallOpts)
}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_VRFManager *VRFManagerCallerSession) NodeAddress() (common.Address, error) {
	return _VRFManager.Contract.NodeAddress(&_VRFManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VRFManager *VRFManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VRFManager *VRFManagerSession) Owner() (common.Address, error) {
	return _VRFManager.Contract.Owner(&_VRFManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VRFManager *VRFManagerCallerSession) Owner() (common.Address, error) {
	return _VRFManager.Contract.Owner(&_VRFManager.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VRFManager *VRFManagerCaller) RequestIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VRFManager.contract.Call(opts, &out, "requestIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VRFManager *VRFManagerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _VRFManager.Contract.RequestIds(&_VRFManager.CallOpts, arg0)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_VRFManager *VRFManagerCallerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _VRFManager.Contract.RequestIds(&_VRFManager.CallOpts, arg0)
}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool isFulfilled)
func (_VRFManager *VRFManagerCaller) RequestMapping(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _VRFManager.contract.Call(opts, &out, "requestMapping", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool isFulfilled)
func (_VRFManager *VRFManagerSession) RequestMapping(arg0 *big.Int) (bool, error) {
	return _VRFManager.Contract.RequestMapping(&_VRFManager.CallOpts, arg0)
}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool isFulfilled)
func (_VRFManager *VRFManagerCallerSession) RequestMapping(arg0 *big.Int) (bool, error) {
	return _VRFManager.Contract.RequestMapping(&_VRFManager.CallOpts, arg0)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xbe8e11df.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords, uint256 blockNumber, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) params, bytes32 msgHash) returns()
func (_VRFManager *VRFManagerTransactor) FulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int, blockNumber *big.Int, params IBLSApkRegistrySignatureCheckParams, msgHash [32]byte) (*types.Transaction, error) {
	return _VRFManager.contract.Transact(opts, "fulfillRandomWords", _requestId, _randomWords, blockNumber, params, msgHash)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xbe8e11df.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords, uint256 blockNumber, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) params, bytes32 msgHash) returns()
func (_VRFManager *VRFManagerSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int, blockNumber *big.Int, params IBLSApkRegistrySignatureCheckParams, msgHash [32]byte) (*types.Transaction, error) {
	return _VRFManager.Contract.FulfillRandomWords(&_VRFManager.TransactOpts, _requestId, _randomWords, blockNumber, params, msgHash)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xbe8e11df.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords, uint256 blockNumber, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) params, bytes32 msgHash) returns()
func (_VRFManager *VRFManagerTransactorSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int, blockNumber *big.Int, params IBLSApkRegistrySignatureCheckParams, msgHash [32]byte) (*types.Transaction, error) {
	return _VRFManager.Contract.FulfillRandomWords(&_VRFManager.TransactOpts, _requestId, _randomWords, blockNumber, params, msgHash)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _nodeAddress, address _blsRegistry) returns()
func (_VRFManager *VRFManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _nodeAddress common.Address, _blsRegistry common.Address) (*types.Transaction, error) {
	return _VRFManager.contract.Transact(opts, "initialize", initialOwner, _nodeAddress, _blsRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _nodeAddress, address _blsRegistry) returns()
func (_VRFManager *VRFManagerSession) Initialize(initialOwner common.Address, _nodeAddress common.Address, _blsRegistry common.Address) (*types.Transaction, error) {
	return _VRFManager.Contract.Initialize(&_VRFManager.TransactOpts, initialOwner, _nodeAddress, _blsRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _nodeAddress, address _blsRegistry) returns()
func (_VRFManager *VRFManagerTransactorSession) Initialize(initialOwner common.Address, _nodeAddress common.Address, _blsRegistry common.Address) (*types.Transaction, error) {
	return _VRFManager.Contract.Initialize(&_VRFManager.TransactOpts, initialOwner, _nodeAddress, _blsRegistry)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VRFManager *VRFManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VRFManager *VRFManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _VRFManager.Contract.RenounceOwnership(&_VRFManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VRFManager *VRFManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VRFManager.Contract.RenounceOwnership(&_VRFManager.TransactOpts)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_VRFManager *VRFManagerTransactor) RequestRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _VRFManager.contract.Transact(opts, "requestRandomWords", _requestId, _numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_VRFManager *VRFManagerSession) RequestRandomWords(_requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _VRFManager.Contract.RequestRandomWords(&_VRFManager.TransactOpts, _requestId, _numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_VRFManager *VRFManagerTransactorSession) RequestRandomWords(_requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _VRFManager.Contract.RequestRandomWords(&_VRFManager.TransactOpts, _requestId, _numWords)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _nodeAddress) returns()
func (_VRFManager *VRFManagerTransactor) SetWhitelist(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _VRFManager.contract.Transact(opts, "setWhitelist", _nodeAddress)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _nodeAddress) returns()
func (_VRFManager *VRFManagerSession) SetWhitelist(_nodeAddress common.Address) (*types.Transaction, error) {
	return _VRFManager.Contract.SetWhitelist(&_VRFManager.TransactOpts, _nodeAddress)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _nodeAddress) returns()
func (_VRFManager *VRFManagerTransactorSession) SetWhitelist(_nodeAddress common.Address) (*types.Transaction, error) {
	return _VRFManager.Contract.SetWhitelist(&_VRFManager.TransactOpts, _nodeAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VRFManager *VRFManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VRFManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VRFManager *VRFManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VRFManager.Contract.TransferOwnership(&_VRFManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VRFManager *VRFManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VRFManager.Contract.TransferOwnership(&_VRFManager.TransactOpts, newOwner)
}

// VRFManagerFillRandomWordsIterator is returned from FilterFillRandomWords and is used to iterate over the raw logs and unpacked data for FillRandomWords events raised by the VRFManager contract.
type VRFManagerFillRandomWordsIterator struct {
	Event *VRFManagerFillRandomWords // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VRFManagerFillRandomWordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFManagerFillRandomWords)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VRFManagerFillRandomWords)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VRFManagerFillRandomWordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFManagerFillRandomWordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFManagerFillRandomWords represents a FillRandomWords event raised by the VRFManager contract.
type VRFManagerFillRandomWords struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFillRandomWords is a free log retrieval operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_VRFManager *VRFManagerFilterer) FilterFillRandomWords(opts *bind.FilterOpts) (*VRFManagerFillRandomWordsIterator, error) {

	logs, sub, err := _VRFManager.contract.FilterLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return &VRFManagerFillRandomWordsIterator{contract: _VRFManager.contract, event: "FillRandomWords", logs: logs, sub: sub}, nil
}

// WatchFillRandomWords is a free log subscription operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_VRFManager *VRFManagerFilterer) WatchFillRandomWords(opts *bind.WatchOpts, sink chan<- *VRFManagerFillRandomWords) (event.Subscription, error) {

	logs, sub, err := _VRFManager.contract.WatchLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFManagerFillRandomWords)
				if err := _VRFManager.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFillRandomWords is a log parse operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_VRFManager *VRFManagerFilterer) ParseFillRandomWords(log types.Log) (*VRFManagerFillRandomWords, error) {
	event := new(VRFManagerFillRandomWords)
	if err := _VRFManager.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the VRFManager contract.
type VRFManagerInitializedIterator struct {
	Event *VRFManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VRFManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VRFManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VRFManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFManagerInitialized represents a Initialized event raised by the VRFManager contract.
type VRFManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VRFManager *VRFManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*VRFManagerInitializedIterator, error) {

	logs, sub, err := _VRFManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VRFManagerInitializedIterator{contract: _VRFManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VRFManager *VRFManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VRFManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _VRFManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFManagerInitialized)
				if err := _VRFManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_VRFManager *VRFManagerFilterer) ParseInitialized(log types.Log) (*VRFManagerInitialized, error) {
	event := new(VRFManagerInitialized)
	if err := _VRFManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VRFManager contract.
type VRFManagerOwnershipTransferredIterator struct {
	Event *VRFManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VRFManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VRFManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VRFManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFManagerOwnershipTransferred represents a OwnershipTransferred event raised by the VRFManager contract.
type VRFManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VRFManager *VRFManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VRFManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VRFManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VRFManagerOwnershipTransferredIterator{contract: _VRFManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VRFManager *VRFManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VRFManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFManagerOwnershipTransferred)
				if err := _VRFManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VRFManager *VRFManagerFilterer) ParseOwnershipTransferred(log types.Log) (*VRFManagerOwnershipTransferred, error) {
	event := new(VRFManagerOwnershipTransferred)
	if err := _VRFManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFManagerRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the VRFManager contract.
type VRFManagerRequestSentIterator struct {
	Event *VRFManagerRequestSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VRFManagerRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFManagerRequestSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VRFManagerRequestSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VRFManagerRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFManagerRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFManagerRequestSent represents a RequestSent event raised by the VRFManager contract.
type VRFManagerRequestSent struct {
	RequestId *big.Int
	NumWords  *big.Int
	Current   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 numWords, address current)
func (_VRFManager *VRFManagerFilterer) FilterRequestSent(opts *bind.FilterOpts) (*VRFManagerRequestSentIterator, error) {

	logs, sub, err := _VRFManager.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &VRFManagerRequestSentIterator{contract: _VRFManager.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 numWords, address current)
func (_VRFManager *VRFManagerFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *VRFManagerRequestSent) (event.Subscription, error) {

	logs, sub, err := _VRFManager.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFManagerRequestSent)
				if err := _VRFManager.contract.UnpackLog(event, "RequestSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestSent is a log parse operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 numWords, address current)
func (_VRFManager *VRFManagerFilterer) ParseRequestSent(log types.Log) (*VRFManagerRequestSent, error) {
	event := new(VRFManagerRequestSent)
	if err := _VRFManager.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
