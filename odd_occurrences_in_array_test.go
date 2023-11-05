package main

import "testing"

func TestSolution21(t *testing.T) {
	expected := 7
	actual := Solution2([]int{9, 3, 9, 3, 9, 7, 9})
	assert2(t, expected, actual)
}

func TestSolution22(t *testing.T) {
	actual := Solution2([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 7, 7})
	assert2(t, 5, actual)
}

func TestSolution23(t *testing.T) {
	actual := Solution2([]int{2, 2, 2, 2, 2, 1, 1, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2})
	assert2(t, 1, actual)
}

func TestSolution24(t *testing.T) {
	actual := Solution2([]int{1, 1, 2})
	assert2(t, 2, actual)
}

func assert2(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Fatalf("Expected %v not equal actual %v", expected, actual)
	}
}
