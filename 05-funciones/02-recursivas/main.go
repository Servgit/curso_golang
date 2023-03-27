package main

import "fmt"

//la siguiente funcion calcula el factorial de un numero
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(3))
	// 3  1 2 3
}
