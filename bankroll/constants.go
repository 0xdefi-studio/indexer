package bankroll

import "github.com/ethereum/go-ethereum/common"

const BANKROLL = "bankroll"
const BANKROLL_WMNT_ADDRESS = "0xB591D637cFd989A21e31873dbE64AFa4BF18f169"
const BANKROLL_MINU_ADDRESS = "0x872dEF0be6A91B212e67BbD56D37B6Cc9513B7B7"

const BANKROLL_WMNT = "WMNT"
const BANKROLL_MINU = "MINU"

var BankrollWmntContractHex = common.HexToAddress(BANKROLL_WMNT_ADDRESS)
var BankrollMinuContractHex = common.HexToAddress(BANKROLL_MINU_ADDRESS)
