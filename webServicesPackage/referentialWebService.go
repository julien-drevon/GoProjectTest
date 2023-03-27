package webServicesPackage

import (
	"clean/core"
	"controller"
	"pokedex/domain"
	"presenter"

	"github.com/gin-gonic/gin"
)

type ReferentialWebService struct {
	ReferentialController *controller.PokemonReferentialController[presenter.HttpResponse[core.PaginationResult[domain.Pokemon]]]
}

func (this *ReferentialWebService) GetReferential() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		var resultChan chan presenter.HttpResponse[core.PaginationResult[domain.Pokemon]] = make(chan presenter.HttpResponse[core.PaginationResult[domain.Pokemon]])
		var errChan chan error = make(chan error)

		go this.GetReferentialValue(resultChan, errChan)
		result := <-resultChan
		err := <-errChan

		this.IndentedJSON(ginContext, result, err)
	}
}

func (this *ReferentialWebService) IndentedJSON(ginContext *gin.Context, result presenter.HttpResponse[core.PaginationResult[domain.Pokemon]], err error) {
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

func (this *ReferentialWebService) GetReferentialValue(stringChan chan presenter.HttpResponse[core.PaginationResult[domain.Pokemon]], errChan chan error) {
	refCotroller := this.ReferentialController
	presenter := refCotroller.GetReferential()
	result, err := presenter.Print()
	stringChan <- result
	errChan <- err
}
