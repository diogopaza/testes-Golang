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
