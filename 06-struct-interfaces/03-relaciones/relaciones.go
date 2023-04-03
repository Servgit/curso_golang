package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

// relacion uno a mcuhos
type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {
	//relacion de uno a uno
	/*
		juan := User{
			nombre: "Juan",
			email:  "juan@gmail.com",
			activo: true,
		}

		miguel := User{
			nombre: "Miguel",
			email:  "miguel@gmail.com",
			activo: true,
		}

		juanStudent := Student{
			user:   juan,
			codigo: "001arsd",
		}
	*/
	//relacion de uno a muchos
	video1 := Video{titulo: "01-Introduccion"}
	video2 := Video{titulo: "02-Instalacion"}

	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{video1, video2},
	}

	video1.curso = curso
	video2.curso = curso

	fmt.Println(video1.curso.titulo)

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
