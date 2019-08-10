package main

import (
	"testing"
)

func BenchmarkTestSomaUmMAisUm(t *testing.B) {
	x := soma(1, 1)
	if x != 2 {
		t.Error("1+1 nao é igual a 2, obtive ", x)
	}
}

func BenchmarkTestSomaUmMaisDois(t *testing.B) {
	x := soma(1, 2)
	if x != 3 {
		t.Error("1+1 nao é igual a 2, obtive ", x)
	}
}

func BenchmarkTesteDivide(t *testing.B) {
	x, err := divide(4, 2)
	if err != nil {
		t.Error(err)
	} else {
		if x != 2 {
			t.Error("Opa 4 / 2 e igual a 2, obtive ", x)
		}
	}

}

func BenchmarkTestDividePorZero(t *testing.B) {
	_, err := divide(4, 0)
	if err != nil {
		t.Error(err)
	}
}
