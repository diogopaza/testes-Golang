package main

import "errors"

var errorDivisaoZero = errors.New("O divisor da operacao nao pode ser 0")

func soma(x, y int) int {
	return x + y
}

func divide(x, y int) (int, error) {

	var err error
	var resultado int

	if y == 0 {
		err = errorDivisaoZero
	} else {
		resultado = x / y
	}
	return resultado, err
}
