package main

import "fmt"

//variable global
var mensaje string

func funcion1() {
	mensaje = "Hola desde la funcion uno"
	fmt.Println(mensaje)
}

func funcion2() {
	mensaje = "Hola desde la funcion dos"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "Hola desde la funcion principal"
	fmt.Println(mensaje)

	defer funcion1()
	funcion2()

}
