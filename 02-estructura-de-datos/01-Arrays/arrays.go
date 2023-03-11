package main

import "fmt"

func main() {
	//
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)
	fmt.Println(numeros[1])

	//arrays con valores

	nombres := [3]string{"Nezuko", "Takemichi", "Noelle"}

	fmt.Println(nombres)

	colores := [...]string{
		"Rojo",
		"Azul",
		"Amarillo",
		"Naranja",
	}

	fmt.Println(colores, len(colores))

	//indices definidos
	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euros",
	}

	fmt.Println(monedas, len(monedas))

	//sub array
	subMoneda := monedas[0:3]
	fmt.Println(subMoneda)
}
