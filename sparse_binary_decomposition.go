package main

import (
	"math"
)

func Solution3(N int) int {
	if N <= 2 {
		return N
	}

	mask := calculateAlternatingMask(N)
	return N & mask
}

func isSparseBinary(num int) bool {
	return (num & (num >> 1)) == 0
}

func calculateAlternatingMask(value int) int {
	binaryIndex := math.Ceil(math.Log2(float64(value)))
	alternatingMask := math.Pow(2, binaryIndex-1)
	zeroBit := 1

	for binaryIndex >= 0 {
		if zeroBit == 1 {
			alternatingMask -= math.Pow(2, binaryIndex)
		}

		zeroBit ^= 1
		binaryIndex--
	}

	return int(alternatingMask)
}
