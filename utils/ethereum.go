package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

func HexToBigInt(hexStr string) *big.Int {
	if hexStr[0:2] != "0x" {
		hexStr = "0x" + hexStr
	}

	return hexutil.MustDecodeBig(hexStr)
}

func FullyValidAddress(hexStr string) bool {
	if hexStr[0:2] != "0x" {
		hexStr = "0x" + hexStr
	}

	if !common.IsHexAddress(hexStr) {
		return false
	}

	return common.HexToAddress(hexStr).Hex() != hexStr
}

func IntOfUnit(amount *big.Int, unit string) *big.Int {
	return big.NewInt(0)
}
