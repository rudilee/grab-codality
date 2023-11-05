package main

import "math/big"

func Solution5(S string) int {
	steps := 0
	V := new(big.Int)
	V, ok := V.SetString(S, 2)
	if ok {
		for ; V.Cmp(&big.Int{}) == 1; steps++ {
			temp := new(big.Int).Set(V)
			if temp.Mod(V, big.NewInt(2)).Cmp(&big.Int{}) == 0 {
				V.Div(V, big.NewInt(2))
			} else {
				V.Sub(V, big.NewInt(1))
			}
		}
	}

	return steps
}
