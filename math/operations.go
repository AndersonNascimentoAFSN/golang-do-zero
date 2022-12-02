package math

import "errors"

// Sum é uma função que recebe dois números inteiros e retorna a soma deles
func Sum(a int, b int) (int, error) {
	res := a + b

	if a < 0 || b < 0 {
		return 0, errors.New("numbers must be positive")
	}

	return res, nil
}

// Multiply é uma função que recebe dois números inteiros e retorna a multiplicação deles
func Multiply(a int, b int) (result int) { // retorno nomeado
	result = a * b
	return
}

// SumTotal é uma função que recebe "n" números inteiros e retorna a soma de todos eles
func SumTotal(numbers ...int) (result int) {
	result = 0

	for _, number := range numbers {
		result += number
	}
	return
}

func Anonymous(x ...int) func() int {
	res := 0

	for _, v := range x {
		res += v
	}

	return func() int {
		return res * res
	}
}

func Operations(x ...int) func(op string) int {
	res := 0
	return func(op string) int {
		if op == "sum" {
			for _, v := range x {
				res += v
			}
			return res
		}
		if op == "multi" {
			for _, v := range x {
				res *= v
			}
			return res
		}
		return 0
	}
}
