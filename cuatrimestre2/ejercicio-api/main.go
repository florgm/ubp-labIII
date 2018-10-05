package main

import (
	productosController "github.com/florgm/ubp-labIII/cuatrimestre2/ejercicio-api/controllers/productos"
	productosService "github.com/florgm/ubp-labIII/cuatrimestre2/ejercicio-api/services/productos"
	"github.com/gin-gonic/gin"
)

func main() {
	productosService.LoadProducts()
	router := gin.Default()

	router.GET("/productos", productosController.GetProductos)
	router.POST("/productos", productosController.PostProductos)

	router.Run(":8080")
}
