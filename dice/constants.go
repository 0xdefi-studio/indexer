package dice

import (
	"github.com/ethereum/go-ethereum/common"
)

const DICE = "dice"
const DICE_OUTCOME = "dice_outcome"
const DICE_CONTRACT1 = "0xbCf0BDaAB5731e1CDF9ef03B8b65003a8a5c4Ccf"
const DICE_CONTRACT2 = "0x8E646CF6490530bf10E55f22abf846BcD9352475"
const DICE_MINU_CONTRACT = "0xc5b55B90F423ab00eF174F4479772cBa8c3bA421"

var DiceContractHex = common.HexToAddress(DICE_CONTRACT1)
var DiceContractV2Hex = common.HexToAddress(DICE_CONTRACT2)
var DiceMinuContractHex = common.HexToAddress(DICE_MINU_CONTRACT)
