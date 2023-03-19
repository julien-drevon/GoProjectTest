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
	controller := controller.NewControllerJSonAndMemory(repo)
	presenter := controller.GetReferentiel()
	result, err := presenter.Print()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}
