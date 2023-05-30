package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ThailanTec/go-hexagonal/adapter/dto"
	"github.com/ThailanTec/go-hexagonal/application/core/ports"
	"github.com/gin-gonic/gin"
)

/*
func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service ports.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProducts(service)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/product", n.With(
		negroni.Wrap(createProducts(service)),
	)).Methods("POST", "OPTIONS")
} */

func GetProducts(c *gin.Context) {
	var service ports.ProductServiceInterface
	id := c.Param("id")
	product, err := service.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot get id")
	}
	err = c.ShouldBindJSON(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot get product")
	}
}

func CreateProducts(c *gin.Context) {
	var service ports.ProductServiceInterface

	var productDto dto.Product

	err := json.NewDecoder(c.Request.Body).Decode(&productDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot get data from body")
		return
	}
	product, err := service.Create(productDto.Name, productDto.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot create product")
		return
	}
	err = c.ShouldBindJSON(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, "cannot get product")
	}

}
