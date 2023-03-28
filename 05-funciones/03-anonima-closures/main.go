package main

import (
	"fmt"
	"strings"
)

// closures son funciones que retornan funciones
func repeat(n int) func(cadena string) string {

	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	repeat3 := repeat(3)
	fmt.Println(repeat3("Hola "))
	fmt.Println(repeat3("Mundo "))

	repeat5 := repeat(5)
	fmt.Println(repeat5("Juan "))
	fmt.Println(repeat5("Maria "))

	/*
			// funcion anonima
			func() {
				fmt.Println("Hola desde la funcion anonima")
			}() //se le colocan los () para que se autoejecute la funcion

		//Tambien se puede asignar una funcion anomima a una variable
		myfunc := func(nombre string) string {
			return fmt.Sprintf("Hola, %s desde la funcion anonima", nombre)
		}
		fmt.Println(myfunc("Juan"))

		//Si solamente colocamos la variable nos va a traer la referencia de memoria y
		//no se va a ejecutar la funcion, pero si le colocamos los parentesis ahi si se ejecuta
	*/
}
