package main

import (
	"controller"
	"gateway"
	"net/http"

	"github.com/gin-gonic/gin"
)

var repo = gateway.NewRepo()

func main() {

	router := gin.Default()
	router.GET("/referential", referential)
	//router.POST("/albums")

	router.Run("localhost:8080")
}
func referential(c *gin.Context) {
	controller := controller.NewReferentialController()
	presenter := controller.GetReferential()
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
