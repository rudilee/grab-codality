package main

import "testing"

func TestSolution6(t *testing.T) {
	actual := Solution6([]int{10, 0, 8, 2, -1, 12, 11, 3})
	if actual != 2 {
		t.Fatalf("Expected 2 got %v", actual)
	}

	actual = Solution6([]int{5, 5})
	if actual != 0 {
		t.Fatalf("Expected 0 got %v", actual)
	}

	actual = Solution6([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5})
	if actual != 0 {
		t.Fatalf("Expected 0 got %v", actual)
	}

	actual = Solution6([]int{-10, -2, 2, 3, 6})
	if actual != 4 {
		t.Fatalf("Expected 4 got %v", actual)
	}
}
