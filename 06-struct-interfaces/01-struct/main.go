package main

import "fmt"

//Struct persona
type Persona struct {
	nombre string
	edad   int
}

//metodos
//metodo para imprimir
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

//metodo para editar el nombre
func (p *Persona) editNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

//herencia
type Empleado struct {
	Persona
	sueldo float32
}

func main() {
	p1 := Persona{"Juan", 28}

	p1.imprimir()
	//fmt.Println(p1)

	//p1.nombre = "Maria"
	p1.editNombre("Sofia")

	p1.imprimir()
	//fmt.Println(p1)

	p2 := Persona{
		nombre: "Ana",
		edad:   22,
	}

	p2.imprimir()
	p2.editNombre("Andrea")
	p2.imprimir()

	//fmt.Println(p2)

	e1 := Empleado{
		sueldo: 550,
	}
	e1.nombre = "Miguel"
	e1.edad = 52
	e1.imprimir()
	fmt.Println(e1)

}
