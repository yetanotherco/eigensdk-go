// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractStrategyManager

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

// ContractStrategyManagerMetaData contains all meta data concerning the ContractStrategyManager contract.
var ContractStrategyManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_BURN_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPOSIT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addStrategiesToDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnShares\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"sharesToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateStrategyDepositDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositIntoStrategy\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"depositedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositIntoStrategyWithSignature\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"depositedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeposits\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"depositSharesToRemove\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeStrategiesFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategiesToRemoveFromWhitelist\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStrategyWhitelister\",\"inputs\":[{\"name\":\"newStrategyWhitelister\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakerDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyList\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"strategies\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerStrategyListLength\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyIsWhitelistedForDeposit\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"whitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyWhitelister\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyAddedToDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyRemovedFromDepositWhitelist\",\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyWhitelisterChanged\",\"inputs\":[{\"name\":\"previousAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxStrategiesExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyWhitelister\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesAmountZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakerAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotWhitelisted\",\"inputs\":[]}]",
	Bin: "0x610100604052348015610010575f5ffd5b5060405161276538038061276583398101604081905261002f91610204565b81816001600160a01b038116610058576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b039081166080521660a0524660c052610076610088565b60e052610081610131565b505061023c565b5f60c051461461012a5750604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b5060e05190565b5f54610100900460ff161561019c5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146101eb575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610201575f5ffd5b50565b5f5f60408385031215610215575f5ffd5b8251610220816101ed565b6020840151909250610231816101ed565b809150509250929050565b60805160a05160c05160e0516124be6102a75f395f6118df01525f61181f01525f818161046a0152818161071a01528181610a7101528181610e0f01528181610f7f01526114dc01525f818161032a015281816105290152818161087c015261111c01526124be5ff3fe608060405234801561000f575f5ffd5b50600436106101e7575f3560e01c806394f649dd11610109578063df5cf7231161009e578063f3b4a0001161006e578063f3b4a000146104c5578063f698da25146104cf578063fabc1cbc146104d7578063fe243a17146104ea575f5ffd5b8063df5cf72314610465578063e7a050aa1461048c578063ee7a7c041461049f578063f2fde38b146104b2575f5ffd5b8063c4623ea1116100d9578063c4623ea1146103f7578063c66567021461041f578063cbc2bd6214610432578063de44acb614610445575f5ffd5b806394f649dd1461039d578063967fc0d2146103be5780639ac01d61146103d1578063b5d8b5b8146103e4575f5ffd5b80635de08ff21161017f5780637ecebe001161014f5780637ecebe0014610306578063886f1195146103255780638b8aac3c146103645780638da5cb5b1461038c575f5ffd5b80635de08ff2146102b6578063663c1de4146102c9578063715018a6146102eb578063724af423146102f3575f5ffd5b806348825e94116101ba57806348825e941461024c578063595c6a67146102735780635ac86ab71461027b5780635c975abb146102ae575f5ffd5b8063136439dd146101eb5780631794bb3c146102005780632eae418c1461021357806332e89ace14610226575b5f5ffd5b6101fe6101f9366004611f6c565b610514565b005b6101fe61020e366004611f97565b6105e9565b6101fe610221366004611fd5565b61070f565b610239610234366004612037565b6107c3565b6040519081526020015b60405180910390f35b6102397f4337f82d142e41f2a8c10547cd8c859bddb92262a61058e77842e24d9dea922481565b6101fe610867565b61029e610289366004612132565b609854600160ff9092169190911b9081161490565b6040519015158152602001610243565b609854610239565b6101fe6102c4366004612152565b610916565b61029e6102d73660046121c3565b60d16020525f908152604090205460ff1681565b6101fe610a55565b6101fe610301366004611f97565b610a66565b6102396103143660046121c3565b60ca6020525f908152604090205481565b61034c7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610243565b6102396103723660046121c3565b6001600160a01b03165f90815260ce602052604090205490565b6033546001600160a01b031661034c565b6103b06103ab3660046121c3565b610aba565b604051610243929190612221565b60cb5461034c906001600160a01b031681565b6102396103df366004612279565b610c32565b6101fe6103f2366004612152565b610cc3565b61040a610405366004611fd5565b610e02565b60408051928352602083019190915201610243565b6101fe61042d3660046121c3565b610e67565b61034c6104403660046122da565b610e7b565b6104586104533660046121c3565b610eaf565b6040516102439190612304565b61034c7f000000000000000000000000000000000000000000000000000000000000000081565b61023961049a366004611f97565b610f22565b6101fe6104ad3660046122da565b610f74565b6101fe6104c03660046121c3565b611096565b61034c620e16e481565b61023961110c565b6101fe6104e5366004611f6c565b61111a565b6102396104f8366004612316565b60cd60209081525f928352604080842090915290825290205481565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610576573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061059a919061234d565b6105b757604051631d77d47760e21b815260040160405180910390fd5b60985481811681146105dc5760405163c61dca5d60e01b815260040160405180910390fd5b6105e582611230565b5050565b5f54610100900460ff161580801561060757505f54600160ff909116105b806106205750303b15801561062057505f5460ff166001145b6106885760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff1916600117905580156106a9575f805461ff0019166101001790555b6106b282611230565b6106bb8461126d565b6106c4836112be565b8015610709575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107585760405163f739589b60e01b815260040160405180910390fd5b604051636ce5768960e11b81526001600160a01b03858116600483015283811660248301526044820183905284169063d9caed12906064015f604051808303815f87803b1580156107a7575f5ffd5b505af11580156107b9573d5f5f3e3d5ffd5b5050505050505050565b6098545f9081906001908116036107ed5760405163840a48d560e01b815260040160405180910390fd5b6107f5611327565b6001600160a01b0385165f90815260ca60205260409020546108268661081f818c8c8c878c610c32565b8688611380565b6001600160a01b0386165f90815260ca6020526040902060018201905561084f868a8a8a6113d2565b92505061085c6001606555565b509695505050505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156108c9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108ed919061234d565b61090a57604051631d77d47760e21b815260040160405180910390fd5b6109145f19611230565b565b60cb546001600160a01b03163314610941576040516320ba3ff960e21b815260040160405180910390fd5b805f5b818110156107095760d15f8585848181106109615761096161236c565b905060200201602081019061097691906121c3565b6001600160a01b0316815260208101919091526040015f205460ff16610a4d57600160d15f8686858181106109ad576109ad61236c565b90506020020160208101906109c291906121c3565b6001600160a01b0316815260208101919091526040015f20805460ff19169115159190911790557f0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe848483818110610a1c57610a1c61236c565b9050602002016020810190610a3191906121c3565b6040516001600160a01b03909116815260200160405180910390a15b600101610944565b610a5d611540565b6109145f61126d565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610aaf5760405163f739589b60e01b815260040160405180910390fd5b61070983838361159a565b6001600160a01b0381165f90815260ce60205260408120546060918291908167ffffffffffffffff811115610af157610af1612023565b604051908082528060200260200182016040528015610b1a578160200160208202803683370190505b5090505f5b82811015610ba8576001600160a01b0386165f90815260cd6020908152604080832060ce9092528220805491929184908110610b5d57610b5d61236c565b5f9182526020808320909101546001600160a01b031683528201929092526040019020548251839083908110610b9557610b9561236c565b6020908102919091010152600101610b1f565b5060ce5f866001600160a01b03166001600160a01b031681526020019081526020015f208181805480602002602001604051908101604052809291908181526020018280548015610c2057602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610c02575b50505050509150935093505050915091565b604080517f4337f82d142e41f2a8c10547cd8c859bddb92262a61058e77842e24d9dea922460208201526001600160a01b03808916928201929092528187166060820152908516608082015260a0810184905260c0810183905260e081018290525f90610cb890610100016040516020818303038152906040528051906020012061165e565b979650505050505050565b60cb546001600160a01b03163314610cee576040516320ba3ff960e21b815260040160405180910390fd5b805f5b818110156107095760d15f858584818110610d0e57610d0e61236c565b9050602002016020810190610d2391906121c3565b6001600160a01b0316815260208101919091526040015f205460ff1615610dfa575f60d15f868685818110610d5a57610d5a61236c565b9050602002016020810190610d6f91906121c3565b6001600160a01b0316815260208101919091526040015f20805460ff19169115159190911790557f4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030848483818110610dc957610dc961236c565b9050602002016020810190610dde91906121c3565b6040516001600160a01b03909116815260200160405180910390a15b600101610cf1565b5f80336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e4d5760405163f739589b60e01b815260040160405180910390fd5b610e59868587866116a4565b915091505b94509492505050565b610e6f611540565b610e78816112be565b50565b60ce602052815f5260405f208181548110610e94575f80fd5b5f918252602090912001546001600160a01b03169150829050565b6001600160a01b0381165f90815260ce6020908152604091829020805483518184028101840190945280845260609392830182828015610f1657602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610ef8575b50505050509050919050565b6098545f908190600190811603610f4c5760405163840a48d560e01b815260040160405180910390fd5b610f54611327565b610f60338686866113d2565b9150610f6c6001606555565b509392505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610fbd5760405163f739589b60e01b815260040160405180910390fd5b816001600160a01b031663d9caed12620e16e4846001600160a01b0316632495a5996040518163ffffffff1660e01b8152600401602060405180830381865afa15801561100c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110309190612380565b6040516001600160e01b031960e085901b1681526001600160a01b03928316600482015291166024820152604481018490526064015f604051808303815f87803b15801561107c575f5ffd5b505af192505050801561108d575060015b156105e5575050565b61109e611540565b6001600160a01b0381166111035760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161067f565b610e788161126d565b5f61111561181c565b905090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611176573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061119a9190612380565b6001600160a01b0316336001600160a01b0316146111cb5760405163794821ff60e01b815260040160405180910390fd5b609854801982198116146111f25760405163c61dca5d60e01b815260040160405180910390fd5b609882905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b609881905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b60cb54604080516001600160a01b03928316815291831660208301527f4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29910160405180910390a160cb80546001600160a01b0319166001600160a01b0392909216919091179055565b6002606554036113795760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161067f565b6002606555565b428110156113a157604051630819bdcd60e01b815260040160405180910390fd5b6113b56001600160a01b0385168484611901565b61070957604051638baa579f60e01b815260040160405180910390fd5b6001600160a01b0383165f90815260d16020526040812054849060ff1661140c57604051632efd965160e11b815260040160405180910390fd5b6114216001600160a01b03851633878661195f565b6040516311f9fbc960e21b81526001600160a01b038581166004830152602482018590528616906347e7ef24906044016020604051808303815f875af115801561146d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611491919061239b565b91505f5f6114a1888789876116a4565b604051631e328e7960e11b81526001600160a01b038b811660048301528a8116602483015260448201849052606482018390529294509092507f000000000000000000000000000000000000000000000000000000000000000090911690633c651cf2906084015f604051808303815f87803b15801561151f575f5ffd5b505af1158015611531573d5f5f3e3d5ffd5b50505050505050949350505050565b6033546001600160a01b031633146109145760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161067f565b5f815f036115bb576040516342061b2560e11b815260040160405180910390fd5b6001600160a01b038085165f90815260cd60209081526040808320938716835292905220548083111561160157604051634b18b19360e01b815260040160405180910390fd5b61160b83826123c6565b6001600160a01b038087165f90815260cd6020908152604080832093891683529290529081208290559091508190036116525761164885856119b9565b6001915050611657565b5f9150505b9392505050565b5f61166761181c565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b5f806001600160a01b0386166116cd576040516316f2ccc960e01b815260040160405180910390fd5b825f036116ed576040516342061b2560e11b815260040160405180910390fd5b6001600160a01b038087165f90815260cd6020908152604080832093881683529290529081205490819003611793576001600160a01b0387165f90815260ce602090815260409091205410611755576040516301a1443960e31b815260040160405180910390fd5b6001600160a01b038781165f90815260ce602090815260408220805460018101825590835291200180546001600160a01b0319169187169190911790555b61179d84826123df565b6001600160a01b038881165f81815260cd602090815260408083208b861680855290835292819020959095558451928352928a169282019290925291820152606081018590527f7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a969060800160405180910390a196929550919350505050565b5f7f000000000000000000000000000000000000000000000000000000000000000046146118dc5750604080518082018252600a81526922b4b3b2b72630bcb2b960b11b60209182015281517f8cad95687ba82c2ce50e74f7b754645e5117c3a5bec8151c0726d5857980a866818301527f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea81840152466060820152306080808301919091528351808303909101815260a0909101909252815191012090565b507f000000000000000000000000000000000000000000000000000000000000000090565b5f5f5f61190e8585611b37565b90925090505f816004811115611926576119266123f2565b1480156119445750856001600160a01b0316826001600160a01b0316145b806119555750611955868686611b79565b9695505050505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052610709908590611c60565b6001600160a01b0382165f90815260ce6020526040812054905b81811015611acb576001600160a01b038481165f90815260ce6020526040902080549185169183908110611a0957611a0961236c565b5f918252602090912001546001600160a01b031603611ac3576001600160a01b0384165f90815260ce602052604090208054611a47906001906123c6565b81548110611a5757611a5761236c565b5f9182526020808320909101546001600160a01b03878116845260ce9092526040909220805491909216919083908110611a9357611a9361236c565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550611acb565b6001016119d3565b818103611aeb57604051632df15a4160e11b815260040160405180910390fd5b6001600160a01b0384165f90815260ce60205260409020805480611b1157611b11612406565b5f8281526020902081015f1990810180546001600160a01b031916905501905550505050565b5f5f8251604103611b6b576020830151604084015160608501515f1a611b5f87828585611d38565b94509450505050611b72565b505f905060025b9250929050565b5f5f5f856001600160a01b0316631626ba7e60e01b8686604051602401611ba1929190612448565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051611bdf9190612460565b5f60405180830381855afa9150503d805f8114611c17576040519150601f19603f3d011682016040523d82523d5f602084013e611c1c565b606091505b5091509150818015611c3057506020815110155b801561195557508051630b135d3f60e11b90611c55908301602090810190840161239b565b149695505050505050565b5f611cb4826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611df29092919063ffffffff16565b905080515f1480611cd4575080806020019051810190611cd4919061234d565b611d335760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161067f565b505050565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611d6d57505f90506003610e5e565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611dbe573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116611de6575f60019250925050610e5e565b965f9650945050505050565b6060611e0084845f85611e08565b949350505050565b606082471015611e695760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161067f565b5f5f866001600160a01b03168587604051611e849190612460565b5f6040518083038185875af1925050503d805f8114611ebe576040519150601f19603f3d011682016040523d82523d5f602084013e611ec3565b606091505b5091509150610cb88783838760608315611f3d5782515f03611f36576001600160a01b0385163b611f365760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161067f565b5081611e00565b611e008383815115611f525781518083602001fd5b8060405162461bcd60e51b815260040161067f9190612476565b5f60208284031215611f7c575f5ffd5b5035919050565b6001600160a01b0381168114610e78575f5ffd5b5f5f5f60608486031215611fa9575f5ffd5b8335611fb481611f83565b92506020840135611fc481611f83565b929592945050506040919091013590565b5f5f5f5f60808587031215611fe8575f5ffd5b8435611ff381611f83565b9350602085013561200381611f83565b9250604085013561201381611f83565b9396929550929360600135925050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f5f5f5f60c0878903121561204c575f5ffd5b863561205781611f83565b9550602087013561206781611f83565b945060408701359350606087013561207e81611f83565b92506080870135915060a087013567ffffffffffffffff8111156120a0575f5ffd5b8701601f810189136120b0575f5ffd5b803567ffffffffffffffff8111156120ca576120ca612023565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156120f9576120f9612023565b6040528181528282016020018b1015612110575f5ffd5b816020840160208301375f602083830101528093505050509295509295509295565b5f60208284031215612142575f5ffd5b813560ff81168114611657575f5ffd5b5f5f60208385031215612163575f5ffd5b823567ffffffffffffffff811115612179575f5ffd5b8301601f81018513612189575f5ffd5b803567ffffffffffffffff81111561219f575f5ffd5b8560208260051b84010111156121b3575f5ffd5b6020919091019590945092505050565b5f602082840312156121d3575f5ffd5b813561165781611f83565b5f8151808452602084019350602083015f5b828110156122175781516001600160a01b03168652602095860195909101906001016121f0565b5093949350505050565b604081525f61223360408301856121de565b82810360208401528084518083526020830191506020860192505f5b8181101561226d57835183526020938401939092019160010161224f565b50909695505050505050565b5f5f5f5f5f5f60c0878903121561228e575f5ffd5b863561229981611f83565b955060208701356122a981611f83565b945060408701356122b981611f83565b959894975094956060810135955060808101359460a0909101359350915050565b5f5f604083850312156122eb575f5ffd5b82356122f681611f83565b946020939093013593505050565b602081525f61165760208301846121de565b5f5f60408385031215612327575f5ffd5b823561233281611f83565b9150602083013561234281611f83565b809150509250929050565b5f6020828403121561235d575f5ffd5b81518015158114611657575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215612390575f5ffd5b815161165781611f83565b5f602082840312156123ab575f5ffd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b818103818111156123d9576123d96123b2565b92915050565b808201808211156123d9576123d96123b2565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52603160045260245ffd5b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b828152604060208201525f611e00604083018461241a565b5f82518060208501845e5f920191825250919050565b602081525f611657602083018461241a56fea26469706673582212207144109046638e552ada8cc695370f34693f671743b5db7c703e16bd2db1b06064736f6c634300081b0033",
}

// ContractStrategyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractStrategyManagerMetaData.ABI instead.
var ContractStrategyManagerABI = ContractStrategyManagerMetaData.ABI

// ContractStrategyManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractStrategyManagerMetaData.Bin instead.
var ContractStrategyManagerBin = ContractStrategyManagerMetaData.Bin

// DeployContractStrategyManager deploys a new Ethereum contract, binding an instance of ContractStrategyManager to it.
func DeployContractStrategyManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _pauserRegistry common.Address) (common.Address, *types.Transaction, *ContractStrategyManager, error) {
	parsed, err := ContractStrategyManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractStrategyManagerBin), backend, _delegation, _pauserRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractStrategyManager{ContractStrategyManagerCaller: ContractStrategyManagerCaller{contract: contract}, ContractStrategyManagerTransactor: ContractStrategyManagerTransactor{contract: contract}, ContractStrategyManagerFilterer: ContractStrategyManagerFilterer{contract: contract}}, nil
}

// ContractStrategyManagerMethods is an auto generated interface around an Ethereum contract.
type ContractStrategyManagerMethods interface {
	ContractStrategyManagerCalls
	ContractStrategyManagerTransacts
	ContractStrategyManagerFilters
}

// ContractStrategyManagerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractStrategyManagerCalls interface {
	DEFAULTBURNADDRESS(opts *bind.CallOpts) (common.Address, error)

	DEPOSITTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	CalculateStrategyDepositDigestHash(opts *bind.CallOpts, staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error)

	Delegation(opts *bind.CallOpts) (common.Address, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error)

	GetStakerStrategyList(opts *bind.CallOpts, staker common.Address) ([]common.Address, error)

	Nonces(opts *bind.CallOpts, signer common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	StakerDepositShares(opts *bind.CallOpts, staker common.Address, strategy common.Address) (*big.Int, error)

	StakerStrategyList(opts *bind.CallOpts, staker common.Address, arg1 *big.Int) (common.Address, error)

	StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error)

	StrategyIsWhitelistedForDeposit(opts *bind.CallOpts, strategy common.Address) (bool, error)

	StrategyWhitelister(opts *bind.CallOpts) (common.Address, error)
}

// ContractStrategyManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractStrategyManagerTransacts interface {
	AddShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error)

	AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error)

	BurnShares(opts *bind.TransactOpts, strategy common.Address, sharesToBurn *big.Int) (*types.Transaction, error)

	DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error)

	DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveDepositShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error)

	RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetStrategyWhitelister(opts *bind.TransactOpts, newStrategyWhitelister common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error)
}

// ContractStrategyManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractStrategyManagerFilters interface {
	FilterDeposit(opts *bind.FilterOpts) (*ContractStrategyManagerDepositIterator, error)
	WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerDeposit) (event.Subscription, error)
	ParseDeposit(log types.Log) (*ContractStrategyManagerDeposit, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractStrategyManagerInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractStrategyManagerInitialized, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractStrategyManagerOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractStrategyManagerOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyManagerPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractStrategyManagerPaused, error)

	FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyAddedToDepositWhitelistIterator, error)
	WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyAddedToDepositWhitelist) (event.Subscription, error)
	ParseStrategyAddedToDepositWhitelist(log types.Log) (*ContractStrategyManagerStrategyAddedToDepositWhitelist, error)

	FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator, error)
	WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyRemovedFromDepositWhitelist) (event.Subscription, error)
	ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*ContractStrategyManagerStrategyRemovedFromDepositWhitelist, error)

	FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyWhitelisterChangedIterator, error)
	WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyWhitelisterChanged) (event.Subscription, error)
	ParseStrategyWhitelisterChanged(log types.Log) (*ContractStrategyManagerStrategyWhitelisterChanged, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyManagerUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractStrategyManagerUnpaused, error)
}

// ContractStrategyManager is an auto generated Go binding around an Ethereum contract.
type ContractStrategyManager struct {
	ContractStrategyManagerCaller     // Read-only binding to the contract
	ContractStrategyManagerTransactor // Write-only binding to the contract
	ContractStrategyManagerFilterer   // Log filterer for contract events
}

// ContractStrategyManager implements the ContractStrategyManagerMethods interface.
var _ ContractStrategyManagerMethods = (*ContractStrategyManager)(nil)

// ContractStrategyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractStrategyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyManagerCaller implements the ContractStrategyManagerCalls interface.
var _ ContractStrategyManagerCalls = (*ContractStrategyManagerCaller)(nil)

// ContractStrategyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractStrategyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyManagerTransactor implements the ContractStrategyManagerTransacts interface.
var _ ContractStrategyManagerTransacts = (*ContractStrategyManagerTransactor)(nil)

// ContractStrategyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractStrategyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractStrategyManagerFilterer implements the ContractStrategyManagerFilters interface.
var _ ContractStrategyManagerFilters = (*ContractStrategyManagerFilterer)(nil)

// ContractStrategyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractStrategyManagerSession struct {
	Contract     *ContractStrategyManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractStrategyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractStrategyManagerCallerSession struct {
	Contract *ContractStrategyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ContractStrategyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractStrategyManagerTransactorSession struct {
	Contract     *ContractStrategyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ContractStrategyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractStrategyManagerRaw struct {
	Contract *ContractStrategyManager // Generic contract binding to access the raw methods on
}

// ContractStrategyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractStrategyManagerCallerRaw struct {
	Contract *ContractStrategyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractStrategyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractStrategyManagerTransactorRaw struct {
	Contract *ContractStrategyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractStrategyManager creates a new instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManager(address common.Address, backend bind.ContractBackend) (*ContractStrategyManager, error) {
	contract, err := bindContractStrategyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManager{ContractStrategyManagerCaller: ContractStrategyManagerCaller{contract: contract}, ContractStrategyManagerTransactor: ContractStrategyManagerTransactor{contract: contract}, ContractStrategyManagerFilterer: ContractStrategyManagerFilterer{contract: contract}}, nil
}

// NewContractStrategyManagerCaller creates a new read-only instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractStrategyManagerCaller, error) {
	contract, err := bindContractStrategyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerCaller{contract: contract}, nil
}

// NewContractStrategyManagerTransactor creates a new write-only instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractStrategyManagerTransactor, error) {
	contract, err := bindContractStrategyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerTransactor{contract: contract}, nil
}

// NewContractStrategyManagerFilterer creates a new log filterer instance of ContractStrategyManager, bound to a specific deployed contract.
func NewContractStrategyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractStrategyManagerFilterer, error) {
	contract, err := bindContractStrategyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerFilterer{contract: contract}, nil
}

// bindContractStrategyManager binds a generic wrapper to an already deployed contract.
func bindContractStrategyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractStrategyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractStrategyManager *ContractStrategyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractStrategyManager.Contract.ContractStrategyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractStrategyManager *ContractStrategyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.ContractStrategyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractStrategyManager *ContractStrategyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.ContractStrategyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractStrategyManager *ContractStrategyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractStrategyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractStrategyManager *ContractStrategyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractStrategyManager *ContractStrategyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTBURNADDRESS is a free data retrieval call binding the contract method 0xf3b4a000.
//
// Solidity: function DEFAULT_BURN_ADDRESS() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) DEFAULTBURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "DEFAULT_BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEFAULTBURNADDRESS is a free data retrieval call binding the contract method 0xf3b4a000.
//
// Solidity: function DEFAULT_BURN_ADDRESS() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) DEFAULTBURNADDRESS() (common.Address, error) {
	return _ContractStrategyManager.Contract.DEFAULTBURNADDRESS(&_ContractStrategyManager.CallOpts)
}

// DEFAULTBURNADDRESS is a free data retrieval call binding the contract method 0xf3b4a000.
//
// Solidity: function DEFAULT_BURN_ADDRESS() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) DEFAULTBURNADDRESS() (common.Address, error) {
	return _ContractStrategyManager.Contract.DEFAULTBURNADDRESS(&_ContractStrategyManager.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCaller) DEPOSITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "DEPOSIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DEPOSITTYPEHASH(&_ContractStrategyManager.CallOpts)
}

// DEPOSITTYPEHASH is a free data retrieval call binding the contract method 0x48825e94.
//
// Solidity: function DEPOSIT_TYPEHASH() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) DEPOSITTYPEHASH() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DEPOSITTYPEHASH(&_ContractStrategyManager.CallOpts)
}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCaller) CalculateStrategyDepositDigestHash(opts *bind.CallOpts, staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "calculateStrategyDepositDigestHash", staker, strategy, token, amount, nonce, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) CalculateStrategyDepositDigestHash(staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	return _ContractStrategyManager.Contract.CalculateStrategyDepositDigestHash(&_ContractStrategyManager.CallOpts, staker, strategy, token, amount, nonce, expiry)
}

// CalculateStrategyDepositDigestHash is a free data retrieval call binding the contract method 0x9ac01d61.
//
// Solidity: function calculateStrategyDepositDigestHash(address staker, address strategy, address token, uint256 amount, uint256 nonce, uint256 expiry) view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) CalculateStrategyDepositDigestHash(staker common.Address, strategy common.Address, token common.Address, amount *big.Int, nonce *big.Int, expiry *big.Int) ([32]byte, error) {
	return _ContractStrategyManager.Contract.CalculateStrategyDepositDigestHash(&_ContractStrategyManager.CallOpts, staker, strategy, token, amount, nonce, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) Delegation() (common.Address, error) {
	return _ContractStrategyManager.Contract.Delegation(&_ContractStrategyManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractStrategyManager.Contract.Delegation(&_ContractStrategyManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerSession) DomainSeparator() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DomainSeparator(&_ContractStrategyManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractStrategyManager.Contract.DomainSeparator(&_ContractStrategyManager.CallOpts)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_ContractStrategyManager *ContractStrategyManagerCaller) GetDeposits(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "getDeposits", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_ContractStrategyManager *ContractStrategyManagerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractStrategyManager.Contract.GetDeposits(&_ContractStrategyManager.CallOpts, staker)
}

// GetDeposits is a free data retrieval call binding the contract method 0x94f649dd.
//
// Solidity: function getDeposits(address staker) view returns(address[], uint256[])
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) GetDeposits(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractStrategyManager.Contract.GetDeposits(&_ContractStrategyManager.CallOpts, staker)
}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_ContractStrategyManager *ContractStrategyManagerCaller) GetStakerStrategyList(opts *bind.CallOpts, staker common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "getStakerStrategyList", staker)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_ContractStrategyManager *ContractStrategyManagerSession) GetStakerStrategyList(staker common.Address) ([]common.Address, error) {
	return _ContractStrategyManager.Contract.GetStakerStrategyList(&_ContractStrategyManager.CallOpts, staker)
}

// GetStakerStrategyList is a free data retrieval call binding the contract method 0xde44acb6.
//
// Solidity: function getStakerStrategyList(address staker) view returns(address[])
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) GetStakerStrategyList(staker common.Address) ([]common.Address, error) {
	return _ContractStrategyManager.Contract.GetStakerStrategyList(&_ContractStrategyManager.CallOpts, staker)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Nonces(opts *bind.CallOpts, signer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "nonces", signer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_ContractStrategyManager *ContractStrategyManagerSession) Nonces(signer common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.Nonces(&_ContractStrategyManager.CallOpts, signer)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Nonces(signer common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.Nonces(&_ContractStrategyManager.CallOpts, signer)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) Owner() (common.Address, error) {
	return _ContractStrategyManager.Contract.Owner(&_ContractStrategyManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Owner() (common.Address, error) {
	return _ContractStrategyManager.Contract.Owner(&_ContractStrategyManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerSession) Paused(index uint8) (bool, error) {
	return _ContractStrategyManager.Contract.Paused(&_ContractStrategyManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractStrategyManager.Contract.Paused(&_ContractStrategyManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) Paused0() (*big.Int, error) {
	return _ContractStrategyManager.Contract.Paused0(&_ContractStrategyManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractStrategyManager.Contract.Paused0(&_ContractStrategyManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractStrategyManager.Contract.PauserRegistry(&_ContractStrategyManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractStrategyManager.Contract.PauserRegistry(&_ContractStrategyManager.CallOpts)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StakerDepositShares(opts *bind.CallOpts, staker common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "stakerDepositShares", staker, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerSession) StakerDepositShares(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerDepositShares(&_ContractStrategyManager.CallOpts, staker, strategy)
}

// StakerDepositShares is a free data retrieval call binding the contract method 0xfe243a17.
//
// Solidity: function stakerDepositShares(address staker, address strategy) view returns(uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StakerDepositShares(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerDepositShares(&_ContractStrategyManager.CallOpts, staker, strategy)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StakerStrategyList(opts *bind.CallOpts, staker common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "stakerStrategyList", staker, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_ContractStrategyManager *ContractStrategyManagerSession) StakerStrategyList(staker common.Address, arg1 *big.Int) (common.Address, error) {
	return _ContractStrategyManager.Contract.StakerStrategyList(&_ContractStrategyManager.CallOpts, staker, arg1)
}

// StakerStrategyList is a free data retrieval call binding the contract method 0xcbc2bd62.
//
// Solidity: function stakerStrategyList(address staker, uint256 ) view returns(address strategies)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StakerStrategyList(staker common.Address, arg1 *big.Int) (common.Address, error) {
	return _ContractStrategyManager.Contract.StakerStrategyList(&_ContractStrategyManager.CallOpts, staker, arg1)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StakerStrategyListLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "stakerStrategyListLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerStrategyListLength(&_ContractStrategyManager.CallOpts, staker)
}

// StakerStrategyListLength is a free data retrieval call binding the contract method 0x8b8aac3c.
//
// Solidity: function stakerStrategyListLength(address staker) view returns(uint256)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StakerStrategyListLength(staker common.Address) (*big.Int, error) {
	return _ContractStrategyManager.Contract.StakerStrategyListLength(&_ContractStrategyManager.CallOpts, staker)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StrategyIsWhitelistedForDeposit(opts *bind.CallOpts, strategy common.Address) (bool, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "strategyIsWhitelistedForDeposit", strategy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_ContractStrategyManager *ContractStrategyManagerSession) StrategyIsWhitelistedForDeposit(strategy common.Address) (bool, error) {
	return _ContractStrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_ContractStrategyManager.CallOpts, strategy)
}

// StrategyIsWhitelistedForDeposit is a free data retrieval call binding the contract method 0x663c1de4.
//
// Solidity: function strategyIsWhitelistedForDeposit(address strategy) view returns(bool whitelisted)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StrategyIsWhitelistedForDeposit(strategy common.Address) (bool, error) {
	return _ContractStrategyManager.Contract.StrategyIsWhitelistedForDeposit(&_ContractStrategyManager.CallOpts, strategy)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCaller) StrategyWhitelister(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractStrategyManager.contract.Call(opts, &out, "strategyWhitelister")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerSession) StrategyWhitelister() (common.Address, error) {
	return _ContractStrategyManager.Contract.StrategyWhitelister(&_ContractStrategyManager.CallOpts)
}

// StrategyWhitelister is a free data retrieval call binding the contract method 0x967fc0d2.
//
// Solidity: function strategyWhitelister() view returns(address)
func (_ContractStrategyManager *ContractStrategyManagerCallerSession) StrategyWhitelister() (common.Address, error) {
	return _ContractStrategyManager.Contract.StrategyWhitelister(&_ContractStrategyManager.CallOpts)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address strategy, address token, uint256 shares) returns(uint256, uint256)
func (_ContractStrategyManager *ContractStrategyManagerTransactor) AddShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "addShares", staker, strategy, token, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address strategy, address token, uint256 shares) returns(uint256, uint256)
func (_ContractStrategyManager *ContractStrategyManagerSession) AddShares(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.AddShares(&_ContractStrategyManager.TransactOpts, staker, strategy, token, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0xc4623ea1.
//
// Solidity: function addShares(address staker, address strategy, address token, uint256 shares) returns(uint256, uint256)
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) AddShares(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.AddShares(&_ContractStrategyManager.TransactOpts, staker, strategy, token, shares)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) AddStrategiesToDepositWhitelist(opts *bind.TransactOpts, strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "addStrategiesToDepositWhitelist", strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToWhitelist)
}

// AddStrategiesToDepositWhitelist is a paid mutator transaction binding the contract method 0x5de08ff2.
//
// Solidity: function addStrategiesToDepositWhitelist(address[] strategiesToWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) AddStrategiesToDepositWhitelist(strategiesToWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.AddStrategiesToDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToWhitelist)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address strategy, uint256 sharesToBurn) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) BurnShares(opts *bind.TransactOpts, strategy common.Address, sharesToBurn *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "burnShares", strategy, sharesToBurn)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address strategy, uint256 sharesToBurn) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) BurnShares(strategy common.Address, sharesToBurn *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.BurnShares(&_ContractStrategyManager.TransactOpts, strategy, sharesToBurn)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address strategy, uint256 sharesToBurn) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) BurnShares(strategy common.Address, sharesToBurn *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.BurnShares(&_ContractStrategyManager.TransactOpts, strategy, sharesToBurn)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerTransactor) DepositIntoStrategy(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "depositIntoStrategy", strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategy(&_ContractStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategy is a paid mutator transaction binding the contract method 0xe7a050aa.
//
// Solidity: function depositIntoStrategy(address strategy, address token, uint256 amount) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) DepositIntoStrategy(strategy common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategy(&_ContractStrategyManager.TransactOpts, strategy, token, amount)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerTransactor) DepositIntoStrategyWithSignature(opts *bind.TransactOpts, strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "depositIntoStrategyWithSignature", strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategyWithSignature(&_ContractStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// DepositIntoStrategyWithSignature is a paid mutator transaction binding the contract method 0x32e89ace.
//
// Solidity: function depositIntoStrategyWithSignature(address strategy, address token, uint256 amount, address staker, uint256 expiry, bytes signature) returns(uint256 depositedShares)
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) DepositIntoStrategyWithSignature(strategy common.Address, token common.Address, amount *big.Int, staker common.Address, expiry *big.Int, signature []byte) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.DepositIntoStrategyWithSignature(&_ContractStrategyManager.TransactOpts, strategy, token, amount, staker, expiry, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "initialize", initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Initialize(&_ContractStrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address initialStrategyWhitelister, uint256 initialPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) Initialize(initialOwner common.Address, initialStrategyWhitelister common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Initialize(&_ContractStrategyManager.TransactOpts, initialOwner, initialStrategyWhitelister, initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Pause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Pause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.PauseAll(&_ContractStrategyManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.PauseAll(&_ContractStrategyManager.TransactOpts)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) RemoveDepositShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "removeDepositShares", staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RemoveDepositShares(&_ContractStrategyManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveDepositShares is a paid mutator transaction binding the contract method 0x724af423.
//
// Solidity: function removeDepositShares(address staker, address strategy, uint256 depositSharesToRemove) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) RemoveDepositShares(staker common.Address, strategy common.Address, depositSharesToRemove *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RemoveDepositShares(&_ContractStrategyManager.TransactOpts, staker, strategy, depositSharesToRemove)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) RemoveStrategiesFromDepositWhitelist(opts *bind.TransactOpts, strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "removeStrategiesFromDepositWhitelist", strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RemoveStrategiesFromDepositWhitelist is a paid mutator transaction binding the contract method 0xb5d8b5b8.
//
// Solidity: function removeStrategiesFromDepositWhitelist(address[] strategiesToRemoveFromWhitelist) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) RemoveStrategiesFromDepositWhitelist(strategiesToRemoveFromWhitelist []common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RemoveStrategiesFromDepositWhitelist(&_ContractStrategyManager.TransactOpts, strategiesToRemoveFromWhitelist)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RenounceOwnership(&_ContractStrategyManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.RenounceOwnership(&_ContractStrategyManager.TransactOpts)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) SetStrategyWhitelister(opts *bind.TransactOpts, newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "setStrategyWhitelister", newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SetStrategyWhitelister(&_ContractStrategyManager.TransactOpts, newStrategyWhitelister)
}

// SetStrategyWhitelister is a paid mutator transaction binding the contract method 0xc6656702.
//
// Solidity: function setStrategyWhitelister(address newStrategyWhitelister) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) SetStrategyWhitelister(newStrategyWhitelister common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.SetStrategyWhitelister(&_ContractStrategyManager.TransactOpts, newStrategyWhitelister)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.TransferOwnership(&_ContractStrategyManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.TransferOwnership(&_ContractStrategyManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Unpause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.Unpause(&_ContractStrategyManager.TransactOpts, newPausedStatus)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.contract.Transact(opts, "withdrawSharesAsTokens", staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_ContractStrategyManager *ContractStrategyManagerSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.WithdrawSharesAsTokens(&_ContractStrategyManager.TransactOpts, staker, strategy, token, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address token, uint256 shares) returns()
func (_ContractStrategyManager *ContractStrategyManagerTransactorSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, token common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractStrategyManager.Contract.WithdrawSharesAsTokens(&_ContractStrategyManager.TransactOpts, staker, strategy, token, shares)
}

// ContractStrategyManagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ContractStrategyManager contract.
type ContractStrategyManagerDepositIterator struct {
	Event *ContractStrategyManagerDeposit // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerDeposit)
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
		it.Event = new(ContractStrategyManagerDeposit)
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
func (it *ContractStrategyManagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerDeposit represents a Deposit event raised by the ContractStrategyManager contract.
type ContractStrategyManagerDeposit struct {
	Staker   common.Address
	Token    common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*ContractStrategyManagerDepositIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerDepositIterator{contract: _ContractStrategyManager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerDeposit) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerDeposit)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7cfff908a4b583f36430b25d75964c458d8ede8a99bd61be750e97ee1b2f3a96.
//
// Solidity: event Deposit(address staker, address token, address strategy, uint256 shares)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseDeposit(log types.Log) (*ContractStrategyManagerDeposit, error) {
	event := new(ContractStrategyManagerDeposit)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractStrategyManager contract.
type ContractStrategyManagerInitializedIterator struct {
	Event *ContractStrategyManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerInitialized)
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
		it.Event = new(ContractStrategyManagerInitialized)
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
func (it *ContractStrategyManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerInitialized represents a Initialized event raised by the ContractStrategyManager contract.
type ContractStrategyManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractStrategyManagerInitializedIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerInitializedIterator{contract: _ContractStrategyManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerInitialized)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseInitialized(log types.Log) (*ContractStrategyManagerInitialized, error) {
	event := new(ContractStrategyManagerInitialized)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractStrategyManager contract.
type ContractStrategyManagerOwnershipTransferredIterator struct {
	Event *ContractStrategyManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerOwnershipTransferred)
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
		it.Event = new(ContractStrategyManagerOwnershipTransferred)
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
func (it *ContractStrategyManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractStrategyManager contract.
type ContractStrategyManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractStrategyManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerOwnershipTransferredIterator{contract: _ContractStrategyManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerOwnershipTransferred)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractStrategyManagerOwnershipTransferred, error) {
	event := new(ContractStrategyManagerOwnershipTransferred)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractStrategyManager contract.
type ContractStrategyManagerPausedIterator struct {
	Event *ContractStrategyManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerPaused)
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
		it.Event = new(ContractStrategyManagerPaused)
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
func (it *ContractStrategyManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerPaused represents a Paused event raised by the ContractStrategyManager contract.
type ContractStrategyManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerPausedIterator{contract: _ContractStrategyManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerPaused)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParsePaused(log types.Log) (*ContractStrategyManagerPaused, error) {
	event := new(ContractStrategyManagerPaused)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerStrategyAddedToDepositWhitelistIterator is returned from FilterStrategyAddedToDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyAddedToDepositWhitelist events raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyAddedToDepositWhitelistIterator struct {
	Event *ContractStrategyManagerStrategyAddedToDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerStrategyAddedToDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
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
		it.Event = new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
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
func (it *ContractStrategyManagerStrategyAddedToDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerStrategyAddedToDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerStrategyAddedToDepositWhitelist represents a StrategyAddedToDepositWhitelist event raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyAddedToDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToDepositWhitelist is a free log retrieval operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterStrategyAddedToDepositWhitelist(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyAddedToDepositWhitelistIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerStrategyAddedToDepositWhitelistIterator{contract: _ContractStrategyManager.contract, event: "StrategyAddedToDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToDepositWhitelist is a free log subscription operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchStrategyAddedToDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyAddedToDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "StrategyAddedToDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
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

// ParseStrategyAddedToDepositWhitelist is a log parse operation binding the contract event 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe.
//
// Solidity: event StrategyAddedToDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseStrategyAddedToDepositWhitelist(log types.Log) (*ContractStrategyManagerStrategyAddedToDepositWhitelist, error) {
	event := new(ContractStrategyManagerStrategyAddedToDepositWhitelist)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyAddedToDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator is returned from FilterStrategyRemovedFromDepositWhitelist and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromDepositWhitelist events raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator struct {
	Event *ContractStrategyManagerStrategyRemovedFromDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
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
		it.Event = new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
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
func (it *ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerStrategyRemovedFromDepositWhitelist represents a StrategyRemovedFromDepositWhitelist event raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyRemovedFromDepositWhitelist struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromDepositWhitelist is a free log retrieval operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterStrategyRemovedFromDepositWhitelist(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerStrategyRemovedFromDepositWhitelistIterator{contract: _ContractStrategyManager.contract, event: "StrategyRemovedFromDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromDepositWhitelist is a free log subscription operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchStrategyRemovedFromDepositWhitelist(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyRemovedFromDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "StrategyRemovedFromDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
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

// ParseStrategyRemovedFromDepositWhitelist is a log parse operation binding the contract event 0x4074413b4b443e4e58019f2855a8765113358c7c72e39509c6af45fc0f5ba030.
//
// Solidity: event StrategyRemovedFromDepositWhitelist(address strategy)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseStrategyRemovedFromDepositWhitelist(log types.Log) (*ContractStrategyManagerStrategyRemovedFromDepositWhitelist, error) {
	event := new(ContractStrategyManagerStrategyRemovedFromDepositWhitelist)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyRemovedFromDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerStrategyWhitelisterChangedIterator is returned from FilterStrategyWhitelisterChanged and is used to iterate over the raw logs and unpacked data for StrategyWhitelisterChanged events raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyWhitelisterChangedIterator struct {
	Event *ContractStrategyManagerStrategyWhitelisterChanged // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerStrategyWhitelisterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerStrategyWhitelisterChanged)
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
		it.Event = new(ContractStrategyManagerStrategyWhitelisterChanged)
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
func (it *ContractStrategyManagerStrategyWhitelisterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerStrategyWhitelisterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerStrategyWhitelisterChanged represents a StrategyWhitelisterChanged event raised by the ContractStrategyManager contract.
type ContractStrategyManagerStrategyWhitelisterChanged struct {
	PreviousAddress common.Address
	NewAddress      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStrategyWhitelisterChanged is a free log retrieval operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterStrategyWhitelisterChanged(opts *bind.FilterOpts) (*ContractStrategyManagerStrategyWhitelisterChangedIterator, error) {

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerStrategyWhitelisterChangedIterator{contract: _ContractStrategyManager.contract, event: "StrategyWhitelisterChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyWhitelisterChanged is a free log subscription operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchStrategyWhitelisterChanged(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerStrategyWhitelisterChanged) (event.Subscription, error) {

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "StrategyWhitelisterChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerStrategyWhitelisterChanged)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
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

// ParseStrategyWhitelisterChanged is a log parse operation binding the contract event 0x4264275e593955ff9d6146a51a4525f6ddace2e81db9391abcc9d1ca48047d29.
//
// Solidity: event StrategyWhitelisterChanged(address previousAddress, address newAddress)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseStrategyWhitelisterChanged(log types.Log) (*ContractStrategyManagerStrategyWhitelisterChanged, error) {
	event := new(ContractStrategyManagerStrategyWhitelisterChanged)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "StrategyWhitelisterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractStrategyManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractStrategyManager contract.
type ContractStrategyManagerUnpausedIterator struct {
	Event *ContractStrategyManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractStrategyManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStrategyManagerUnpaused)
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
		it.Event = new(ContractStrategyManagerUnpaused)
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
func (it *ContractStrategyManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractStrategyManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractStrategyManagerUnpaused represents a Unpaused event raised by the ContractStrategyManager contract.
type ContractStrategyManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractStrategyManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractStrategyManagerUnpausedIterator{contract: _ContractStrategyManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractStrategyManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractStrategyManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractStrategyManagerUnpaused)
				if err := _ContractStrategyManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractStrategyManager *ContractStrategyManagerFilterer) ParseUnpaused(log types.Log) (*ContractStrategyManagerUnpaused, error) {
	event := new(ContractStrategyManagerUnpaused)
	if err := _ContractStrategyManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
