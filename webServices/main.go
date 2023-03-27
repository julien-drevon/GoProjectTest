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
	repo, _ := gateway.NewRepoWithPersistance("pokedex.json")
	//config := configuration{Repo: repo, ReferentialController: controller}
	router := gin.Default()
	router.GET("/referential", (ReferentialWebService{ReferentialController: controller.NewReferentialController()}).GetReferential())
	router.POST("/post", (PokedexWebService{PokedexController: controller.NewControllerFileWeb(&repo)}).Post())

	router.Run("localhost:8080")
}
