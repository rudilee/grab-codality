package main

func Solution2(A []int) int {
	bag := make(map[int]int)

	for i := 0; i < len(A); i++ {
		if _, ok := bag[A[i]]; ok {
			bag[A[i]]++
		} else {
			bag[A[i]] = 1
		}
	}

	for key, val := range bag {
		if val%2 == 1 {
			return key
		}
	}

	return 0
}
