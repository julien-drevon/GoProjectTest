package main

import (
	"clean/core"
	"controller"
	"gateway"
	"net/http"
	"pokedex/domain"

	"github.com/gin-gonic/gin"
)

type configuration struct {
	Repo                  gateway.Repo
	ReferentialController controller.PokemonReferentialController[core.PaginationResult[domain.Pokemon]]
}

func main() {
	//repo, _ := gateway.NewRepoForWithPersistance("pokedex.json")
	controller := controller.NewReferentialController()
	//config := configuration{Repo: repo, ReferentialController: controller}
	router := gin.Default()
	router.GET("/referential", referential(controller))
	//router.POST("/albums")

	router.Run("localhost:8080")
}

func referential(controller controller.PokemonReferentialController[core.PaginationResult[domain.Pokemon]]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var result chan core.PaginationResult[domain.Pokemon] = make(chan core.PaginationResult[domain.Pokemon])
		var err chan error = make(chan error)

		go GetReferentialValue(controller, result, err)
		ReturnJSON(c, <-result, <-err)
	}
}

func GetReferentialValue(controller controller.PokemonReferentialController[core.PaginationResult[domain.Pokemon]], stringChan chan core.PaginationResult[domain.Pokemon], errChan chan error) {
	presenter := controller.GetReferential()
	result, err := presenter.Print()
	stringChan <- result
	errChan <- err
}

func ReturnJSON[T any](c *gin.Context, data T, err error) {
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		c.IndentedJSON(http.StatusOK, data)
	}
}
