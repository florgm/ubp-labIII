package productos

import (
	"net/http"

	productosService "github.com/florgm/ubp-labIII/cuatrimestre2/ejercicio-api/services/productos"
	"github.com/florgm/ubp-labIII/cuatrimestre2/ejercicio-api/utils"
	"github.com/gin-gonic/gin"
)

func GetProductos(c *gin.Context) {

	productsList := productosService.GetProductsList()

	c.JSON(http.StatusOK, productsList)
}

func PostProductos(c *gin.Context) {

	body := utils.GetJsonBody(c.Request)

	productosService.AddProduct(body)
	c.String(http.StatusCreated, "ok")
}
