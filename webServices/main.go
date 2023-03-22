package main

import (
	"clean/core"
	"controller"
	"gateway"
	"net/http"
	"pokedex/domain"
	"presenter"

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

func referential(controller controller.PokemonReferentialController[presenter.HttpResponse[core.PaginationResult[domain.Pokemon]]]) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		var resultChan chan presenter.HttpResponse[core.PaginationResult[domain.Pokemon]] = make(chan presenter.HttpResponse[core.PaginationResult[domain.Pokemon]])
		var errChan chan error = make(chan error)

		go GetReferentialValue(controller, resultChan, errChan)
		result := <-resultChan
		err := <-errChan

		IndentedJSON(
			ginContext,
			func(c *gin.Context) {
				if result.Error != "" {
					c.IndentedJSON(result.Status, result.Error)
				} else {
					c.IndentedJSON(result.Status, result.Data)
				}
			},
			err)
	}
}

func GetReferentialValue(controller controller.PokemonReferentialController[presenter.HttpResponse[core.PaginationResult[domain.Pokemon]]], stringChan chan presenter.HttpResponse[core.PaginationResult[domain.Pokemon]], errChan chan error) {
	presenter := controller.GetReferential()
	result, err := presenter.Print()
	stringChan <- result
	errChan <- err
}

func IndentedJSON(c *gin.Context, indentedJsonWithoutErr func(*gin.Context), err error) {
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		indentedJsonWithoutErr(c)
	}
}
