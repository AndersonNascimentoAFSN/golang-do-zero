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
