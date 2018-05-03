package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("1 + 1 =", sum(1, 1))

	a := 1
	b := 0

	div, err := div(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("%d / %d = %d", a, b, div))
	}
}

func sum(a, b int) int {
	return a + b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return a / b, nil
}
