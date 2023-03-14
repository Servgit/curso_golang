package main

import "fmt"

func main() {

	a := 2

	a = a + 2

	fmt.Println(a)

	//Operadores de asignacion

	//Suma en asignacion
	a += 2
	fmt.Println(a)

	//Resta en asignacion
	a -= 2
	fmt.Println(a)

	//Multiplicacion en asignacion
	a *= 2
	fmt.Println(a)

	//Division en asignacion
	a /= 2
	fmt.Println(a)

	//Modulo en asignacion
	a %= 2
	fmt.Println(a)
}
