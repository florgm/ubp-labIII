package main

import (
	"fmt"

	"github.com/florgm/ubp-labIII/ejemplos/errores-estadistica/tendencia"
)

func main() {

	t := tendencia.Tendencia{
		Valores: []float32{8, 5, 7, 5, 4, 3, 4, 4, 6, 7, 9},
	}

	prom, err := t.Promedio()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("Promedio: ", prom)

	/*moda, err := t.Moda()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("Moda: ", moda)*/

	mediana, err := t.Mediana()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("Mediana: ", mediana)
}
