package main

import "sort"

func Solution4(A []int) int {
	sort.Ints(A)

	positivePos := sort.Search(len(A), func(i int) bool {
		return A[i] > 0
	})

	positives := A[positivePos:]

	if len(positives) > 1 {
		if positives[0] <= 1 {
			for i, val := range positives {
				lessVal := val - 1
				moreVal := val + 1

				if i-1 < 0 && lessVal > 1 {
					return lessVal
				}

				if i-1 >= 0 {
					if positives[i-1] < lessVal {
						return lessVal
					}
				}

				if i+1 >= len(positives) {
					return moreVal
				}

				if positives[i+1] > moreVal {
					return moreVal
				}
			}
		}
	}

	return 1
}
