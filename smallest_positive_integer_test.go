package main

import "testing"

func TestSolution4(t *testing.T) {
	actual := Solution4([]int{1, 3, 6, 4, 1, 2})
	if actual != 5 {
		t.Fatalf("Expected 5 got %v", actual)
	}

	actual = Solution4([]int{1, 2, 3})
	if actual != 4 {
		t.Fatalf("Expected 4 got %v", actual)
	}

	actual = Solution4([]int{-1, -3})
	if actual != 1 {
		t.Fatalf("Expected 1 got %v", actual)
	}

	actual = Solution4([]int{})
	if actual != 1 {
		t.Fatalf("Expected 1 got %v", actual)
	}

	actual = Solution4([]int{-10, -33, 23, 64, 68, -4, 100, 123, 70, 70, 0, 42})
	if actual != 1 {
		t.Fatalf("Expected 1 got %v", actual)
	}

	actual = Solution4([]int{10, 32, 72, 12, 1, 2, 3, 4, 5, 6, 90, -15})
	if actual != 7 {
		t.Fatalf("Expected 7 got %v", actual)
	}

	actual = Solution4([]int{-1000000})
	if actual != 1 {
		t.Fatalf("Expected 1 got %v", actual)
	}

	actual = Solution4([]int{1000000})
	if actual != 1 {
		t.Fatalf("Expected 1 got %v", actual)
	}

	actual = Solution4([]int{1000000, 1, 2, 3, 4, 5})
	if actual != 6 {
		t.Fatalf("Expected 6 got %v", actual)
	}

	actual = Solution4([]int{1, 3, 5, 7, 9})
	if actual != 2 {
		t.Fatalf("Expected 2 got %v", actual)
	}

	actual = Solution4([]int{1, 2, 3, 4, 5, 6, 7, 8})
	if actual != 9 {
		t.Fatalf("Expected 9 got %v", actual)
	}

	actual = Solution4([]int{-1000001, 2, 3, 4, 5, 0, 1234})
	if actual != 1 {
		t.Fatalf("Expected 1 got %v", actual)
	}

	actual = Solution4([]int{-1000001, 3, 4, 5, 0, 1234, 1})
	if actual != 2 {
		t.Fatalf("Expected 2 got %v", actual)
	}
}
