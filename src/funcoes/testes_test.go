package main

import (
	"testing"
)

func TestSomaUmMAisUm(t *testing.T) {
	x := soma(1, 1)
	if x != 2 {
		t.Error("1+1 nao é igual a 2, obtive ", x)
	}
}

func TestSomaUmMaisDois(t *testing.T) {
	x := soma(1, 2)
	if x != 3 {
		t.Error("1+1 nao é igual a 2, obtive ", x)
	}
}

func TesteDivide(t *testing.T) {
	x, err := divide(4, 2)
	if err != nil {
		t.Error(err)
	} else {
		if x != 2 {
			t.Error("Opa 4 / 2 e igual a 2, obtive ", x)
		}
	}

}

func TestDividePorZero(t *testing.T) {
	_, err := divide(4, 0)
	if err != nil {
		t.Error(err)
	}
}
