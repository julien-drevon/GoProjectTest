package main

import (
	"controller"
	"gateway"

	"github.com/gin-gonic/gin"
)

type configuration struct {
	Repo gateway.Repo
}

func main() {
	//repo, _ := gateway.NewRepoForWithPersistance("pokedex.json")
	//config := configuration{Repo: repo, ReferentialController: controller}
	router := gin.Default()
	router.GET("/referential", (ReferentialWebService{ReferentialController: controller.NewReferentialController()}).GetReferential())
	//router.POST("/albums")

	router.Run("localhost:8080")
}
