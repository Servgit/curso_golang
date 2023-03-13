package main

import (
	"fmt"
	"strings"
)

func main() {

	dias := make(map[int]string)
	dia := "dia"

	fmt.Println(dias)

	//Agregar datos
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miercoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sabado"

	fmt.Println(dias)

	dias[7] = "SABADO"

	fmt.Println(dias)

	//convertimos los strings a mayusculas con el metodo strings.ToUpper()

	fmt.Println(strings.ToUpper(" habia un dia un texto en minuscula"))
	fmt.Println(strings.ToUpper(dia))

	//borramos el primer dia
	delete(dias, 1)
	fmt.Println(dias, len(dias))

	//Nuevo Mapa
	estudiantes := make(map[string][]int)

	estudiantes["Alex"] = []int{13, 15, 16}
	estudiantes["Juan"] = []int{14, 13, 17}

	fmt.Println(estudiantes)
	fmt.Println(estudiantes["Alex"])
	fmt.Println(estudiantes["Alex"][1])
}
