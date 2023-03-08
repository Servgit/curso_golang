package main

import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	fmt.Print(hola)

	fmt.Print(mundo)

	nombre := "juan"
	edad := 25

	fmt.Printf(" hola, %s y su edad es %d \n", nombre, edad)
	//en el caso de que no sepamos que tipo de dato se va a imprimir utilizamos la v
	fmt.Printf(" hola, %v y su edad es %v \n", nombre, edad)

	mensaje := fmt.Sprintf(" hola, %v y su edad es %v", nombre, edad)

	fmt.Println(mensaje)

	fmt.Printf("nombre %T \n ", edad)

	fmt.Print("Ingrese otro nombre")
	fmt.Scanln(&nombre)

	fmt.Println("otro nombre: ", nombre)

}
