package main

import (
	"fmt"
	"strings"
)

func esPalindromo(palabra string) {
	//Oso
	fmt.Println(palabra)
	palabra = strings.ToLower(palabra)
	palabra = strings.ReplaceAll(palabra, " ", "")
	fmt.Println(palabra)
}

func main() {
	esPalindromo("Luz azul")

}
