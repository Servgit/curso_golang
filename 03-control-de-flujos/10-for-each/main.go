package main

import "fmt"

func main() {
	nombres := [...]string{"Juan", "Miguel", "Jose"}
	/*
		for i := 0; i < len(nombres); i++ {
			fmt.Println(nombres[i])
		}
	*/

	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}
}
