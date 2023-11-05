package main

import (
	"testing"
)

func TestSolution5(t *testing.T) {
	actual := Solution5("011100")
	if actual != 7 {
		t.Fatalf("Expected 7 got %v", actual)
	}

	actual = Solution5("111")
	if actual != 5 {
		t.Fatalf("Expected 5 got %v", actual)
	}

	actual = Solution5("1111010101111")
	if actual != 22 {
		t.Fatalf("Expected 22 got %v", actual)
	}

	//actual = Solution5(strings.Repeat("1111", 100000))
	//if actual != 799999 {
	//	t.Fatalf("Expected 799999 got %v", actual)
	//}
}
