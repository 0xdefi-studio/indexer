package dice

import (
	"github.com/ethereum/go-ethereum/common"
)

const DICE = "dice"
const DICE_CONTRACT1 = "0x2760D70eC200380ce4C9e08a4ea287e9388ea9E2"
const DICE_CONTRACT2 = "0x8E646CF6490530bf10E55f22abf846BcD9352475"

var DiceContractHex = common.HexToAddress(DICE_CONTRACT1)
var DiceContractV2Hex = common.HexToAddress(DICE_CONTRACT2)
