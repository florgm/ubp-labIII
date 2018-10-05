package main

import (
	"fmt"

	"github.com/florgm/ubp-labIII/ejemplos/ej_03/estructuras"
)

/*func main() {
	alumn := &estructuras.Alumnos{
		Nombre:   "Pedro",
		Legajo:   63107,
		Promedio: 7.83,
	}
	fmt.Println("Viejo ", alumn.Nombre)
	modificarAlumno(alumn)
	fmt.Println("Nuevo", alumn.Nombre)
}

func modificarAlumno(al *estructuras.Alumnos) {
	al.Nombre = "Jose"
	al.Promedio = 8.02
}*/

func main() {
	alumn := &estructuras.Alumnos{
		Nombre:   "Pedro",
		Legajo:   63107,
		Promedio: 7.83,
	}
	fmt.Println("Viejo ", alumn.Nombre)
	alumn.Modificar()
	fmt.Println("Nuevo", alumn.Nombre)
}
