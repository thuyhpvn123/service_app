package common

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

const (
	MAX_VALIDATOR = 101
	// TODO MOVE TO CONFIG
	TRANSFER_GAS_COST = 20000
	PUNISH_GAS_COST   = 10000

	BLOCK_GAS_LIMIT                       = 1000000000
	BASE_FEE_INCREASE_GAS_USE_THRESH_HOLD = 500000000
	MINIMUM_BASE_FEE                      = 1000000000
	BASE_FEE_CHANGE_PERCENTAGE            = 12.5
)

var (
	VALIDATOR_STAKE_POOL_ADDRESS   = common.Address{}
	SLASH_VALIDATOR_AMOUNT         = uint256.NewInt(0).SetBytes(common.FromHex("8ac7230489e80000"))
	MINIMUM_VALIDATOR_STAKE_AMOUNT = uint256.NewInt(0).SetBytes(common.FromHex("8ac7230489e80000"))
)
