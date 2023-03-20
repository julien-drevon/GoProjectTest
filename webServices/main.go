package main

import (
	"controller"
	"gateway"
	"net/http"

	"github.com/gin-gonic/gin"
)

var repo = gateway.NewRepo()

func main() {
	//repo := gateway.NewRepo()

	router := gin.Default()
	router.GET("/referentiel", referentiel)
	//router.POST("/albums")

	router.Run("localhost:8080")
}
func referentiel(c *gin.Context) {
	controller := controller.NewReferentielController()
	presenter := controller.GetReferentiel()
	result, err := presenter.Print()

	ReturnJSON(c, result, err)
}

func ReturnJSON[T any](c *gin.Context, data T, err error) {
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusOK, data)
	}
}
