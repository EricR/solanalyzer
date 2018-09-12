package utils

import (
	"math/big"
	"strconv"
)

func NewMaxIntSize(e int) *big.Int {
	if e == 0 {
		return big.NewInt(0)
	}

	expResult := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(e)), big.NewInt(0))
	return big.NewInt(0).Sub(expResult, big.NewInt(1))
}

func MustParseBigInt(str string) *big.Int {
	i := big.NewInt(0)

	i, ok := i.SetString(str, 10)
	if !ok {
		panic("Could not parse big int")
	}

	return i
}

func MustParseUint(str string) uint {
	i, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		panic("Could not parse uint")
	}

	return uint(i)
}

func SmallestIntSize(i *big.Int, negative bool) int {
	max := big.NewInt(0)

	for n := 8; n <= 256; n += 8 {
		if negative {
			max = big.NewInt(0).Add(NewMaxIntSize(n-1), big.NewInt(1))
		} else {
			max = NewMaxIntSize(n)
		}

		if i.Cmp(max) < 1 {
			return n
		}
	}

	return 0
}
