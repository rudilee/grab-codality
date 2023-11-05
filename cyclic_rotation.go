package main

func Solution1(A []int, K int) []int {
	lenA := len(A)

	if lenA == 0 {
		return A
	}

	if lenA < K {
		K = K % lenA

		if K == 0 {
			return A
		}
	}

	return append(A[lenA-K:lenA], A[0:lenA-K]...)
}
