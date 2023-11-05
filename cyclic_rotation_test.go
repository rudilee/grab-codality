package main

import (
	"reflect"
	"testing"
)

func TestSolution10(t *testing.T) {
	actual := Solution1([]int{}, 10)
	expected := []int{}
	assert(t, actual, expected)
}

func TestSolution11(t *testing.T) {
	actual := Solution1([]int{3, 8, 9, 7, 6}, 3)
	expected := []int{9, 7, 6, 3, 8}
	assert(t, actual, expected)
}

func TestSolution12(t *testing.T) {
	actual := Solution1([]int{0, 0, 0}, 1)
	expected := []int{0, 0, 0}
	assert(t, actual, expected)
}

func TestSolution13(t *testing.T) {
	actual := Solution1([]int{1, 2, 3, 4}, 4)
	expected := []int{1, 2, 3, 4}
	assert(t, actual, expected)
}

func TestSolution14(t *testing.T) {
	actual := Solution1([]int{1, 2, 3, 4, 1, 2, 3}, 4)
	expected := []int{4, 1, 2, 3, 1, 2, 3}
	assert(t, actual, expected)
}

func TestSolution15(t *testing.T) {
	actual := Solution1([]int{1, 2, 3, 4}, 6)
	expected := []int{3, 4, 1, 2}
	assert(t, actual, expected)
}

func TestSolution16(t *testing.T) {
	actual := Solution1([]int{1, 2, 3, 4}, 8)
	expected := []int{1, 2, 3, 4}
	assert(t, actual, expected)
}

func TestSolution17(t *testing.T) {
	actual := Solution1([]int{1, 2, 3, 4}, 0)
	expected := []int{1, 2, 3, 4}
	assert(t, actual, expected)
}

func assert(t *testing.T, actual []int, expected []int) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
