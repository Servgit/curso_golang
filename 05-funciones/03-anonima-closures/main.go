package main

import "fmt"

//closures
func main() {

	/*
		// funcion anonima
		func() {
			fmt.Println("Hola desde la funcion anonima")
		}() //se le colocan los () para que se autoejecute la funcion
	*/
	//Tambien se puede asignar una funcion anomima a una variable
	myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola, %s desde la funcion anonima", nombre)
	}
	fmt.Println(myfunc("Juan"))

	/*Si solamente colocamos la variable nos va a traer la referencia de memoria y
	no se va a ejecutar la funcion, pero si le colocamos los parentesis ahi si se ejecuta
	*/
}
