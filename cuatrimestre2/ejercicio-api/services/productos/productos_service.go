package productos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	productosDomain "github.com/florgm/ubp-labIII/cuatrimestre2/ejercicio-api/domain/productos"
)

var (
	productsList = make([]productosDomain.Producto, 0)
)

func LoadProducts() {
	bytes, err := ioutil.ReadFile("./productos.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(bytes, &productsList); err != nil {
		panic(err)
	}

	fmt.Println(productsList)
}

func GetProductsList() *[]productosDomain.Producto {
	return &productsList
}

func AddProduct(body []byte) {
	var p productosDomain.Producto

	json.Unmarshal(body, &p)

	productsList = append(productsList, p)

	bytes, _ := json.Marshal(&productsList)

	ioutil.WriteFile("./productos.json", bytes, 0644)
}
