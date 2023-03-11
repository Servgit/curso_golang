package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("hola, ", nombre)
}

func sumar(a, b int) int {
	return a + b
}

func main() {
	saludar("Juan")
	r := sumar(10, 20)
	fmt.Println("La suma es: ", r)
}
