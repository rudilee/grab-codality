package main

import "strconv"

//func Solution(N int) (longestGap int) {
//	longestGap = 0
//	currentGap := 0
//
//	binary := strconv.FormatInt(int64(N), 2)
//	println(binary)
//	for pos, char := range binary {
//		if pos > 0 {
//			if char == 49 {
//				if currentGap > longestGap {
//					longestGap = currentGap
//				}
//
//				currentGap = 0
//			} else {
//				currentGap++
//			}
//		}
//	}
//
//	return
//}

func countTrailingBits(N uint) uint {
	var _lookup = [...]uint{
		32, 0, 1, 26, 2, 23, 27, 0, 3, 16, 24, 30, 28, 11, 0, 13, 4,
		7, 17, 0, 25, 22, 31, 15, 29, 10, 12, 6, 0, 21, 14, 9, 5,
		20, 8, 19, 18}

	return _lookup[(N&-N)%37]
}

func Solution(N int) (maxLength int) {
	maxLength = 0
	currentLength := 0
	trailing := countTrailingBits(uint(N))
	binary := strconv.FormatInt(int64(N), 2)
	println(binary)

	if trailing > 0 {
		N = N >> trailing
	}

	for N > 0 {
		// N % 2 gets the bit at the rightmost part and if its 1 then reset the counter
		if N%2 == 1 {
			currentLength = 0
		} else {
			currentLength += 1
		}

		if currentLength > maxLength {
			maxLength = currentLength
		}

		N >>= 1
	}

	return
}
