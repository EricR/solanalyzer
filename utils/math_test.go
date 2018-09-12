package utils

import (
	"math/big"
	"testing"
)

func TestNewMaxIntSize(t *testing.T) {
	if NewMaxIntSize(8).Cmp(big.NewInt(255)) != 0 {
		t.Error("Max int size calculation is wrong")
	}
}

func TestMustParseBigInt(t *testing.T) {
	defer testPanic(t)

	MustParseBigInt("a")
}

func TestSmallestIntSize(t *testing.T) {
	if SmallestIntSize(big.NewInt(255), false) != 8 {
		t.Error("Smallest int size calculation is wrong")
	}
	if SmallestIntSize(big.NewInt(256), false) != 16 {
		t.Error("Smallest int size calculation is wrong")
	}
	if SmallestIntSize(big.NewInt(128), true) != 8 {
		t.Error("Smallest int size calculation is wrong")
	}
	if SmallestIntSize(big.NewInt(129), true) != 16 {
		t.Error("Smallest int size calculation is wrong")
	}
	if SmallestIntSize(big.NewInt(0), false) != 8 {
		t.Error("Smallest int size calculation is wrong")
	}
	if SmallestIntSize(big.NewInt(0), true) != 8 {
		t.Error("Smallest int size calculation is wrong")
	}
}

func testPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Error("Did not panic")
	}
}
