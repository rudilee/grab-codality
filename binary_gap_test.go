package main

import "testing"

func TestSolution(t *testing.T) {
	if Solution(9) != 2 {
		t.Fatalf("Binary Gap for 9 should be 2")
	}

	expected := Solution(529)
	if expected != 4 {
		t.Fatalf("Binary Gap for 529 should be 4 not %v", expected)
	}

	if Solution(1041) != 5 {
		t.Fatalf("Binary Gap for 1041 should be 5")
	}

	if Solution(32) != 0 {
		t.Fatalf("Binary Gap for 32 should be 0")
	}

	expected1 := Solution(15)
	if expected1 != 0 {
		t.Fatalf("Binary Gap for 15 should be 0 no %v", expected1)
	}

	actual := Solution(561892)
	if actual != 3 {
		t.Fatalf("Binary Gap for 561892 should be 3 not %v", actual)
	}

	actual2 := Solution(74901729)
	if actual2 != 4 {
		t.Fatalf("Binary Gap for 74901729 should be 4 not %v", actual2)
	}

	actual3 := Solution(1376796946)
	if actual3 != 5 {
		t.Fatalf("Binary Gap for 1376796946 should be 5 not %v", actual3)
	}
}
