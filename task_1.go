package main

import (
	"math"
	"sort"
)

func Solution6(A []int) int {
	longestDistance := 0

	sort.Ints(A)
	for i, distanceFromCenter := range A {
		if i+1 < len(A) {
			distanceInBetween := math.Abs(float64(distanceFromCenter - A[i+1]))
			currentDistance := int(math.Floor(distanceInBetween / 2))

			if currentDistance > longestDistance {
				longestDistance = currentDistance
			}
		}
	}

	return longestDistance
}
