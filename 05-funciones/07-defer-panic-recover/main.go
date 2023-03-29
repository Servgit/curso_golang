package main

import (
	"fmt"
	"os"
)

func main() {

	//funcion
	defer func() { //utilizamos defer para que esta funcion se ejecute al terminar nuestra applicacion
		if error := recover(); error != nil { //recover sirve para controlar todos los panicos que tal vez sucedan en nuestra applicaicon y para que no se detenga de forma brusca
			fmt.Println("Ups!, al parecer el programa no finalizo de forma correcta")
		}
	}()

	if file, error := os.Open("hola.txt"); error != nil {
		panic("No fue posible obtener el archivo!!")
	} else {

		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}

}
