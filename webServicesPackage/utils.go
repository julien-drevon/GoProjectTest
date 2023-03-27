package webServicesPackage

import (
	"clean/core"
	"controller"
	"gateway"
	"net/http"
	"pokedex/domain"
	"presenter"

	"github.com/gin-gonic/gin"
	//"github.com/google/wire"
)

func IndentedJSON(c *gin.Context, indentedJsonWithoutErr func(*gin.Context), err error) {
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	} else {
		indentedJsonWithoutErr(c)
	}
}

func Run() *gin.Engine {
	config, err := initConfiguration("pokedex.json")
	if err != nil {
		panic(err)
	}

	router := initRoutes(config)

	return router
}

func initConfiguration(jsonFilePath string) (Configuration, error) {
	repo, err := gateway.NewRepoWithPersistance(jsonFilePath)
	return NewConfiguration(repo, *controller.NewReferentialController(), *controller.NewControllerFileWeb(&repo)), err
	// wire.Build(
	// 	gateway.NewRepoWithPersistance,
	// 	controller.NewReferentialController,
	// 	controller.NewControllerFileWeb,
	// 	NewConfiguration,
	// 	//wire.Struct(new(Configuration), "*"),
	// )
	// return Configuration{}, nil
	// wire.Build(gateway.NewRepoWithPersistance, controller.NewReferentialController, controller.NewControllerFileWeb, NewConfiguration)
	// return Configuration{}
}

func initRoutes(config Configuration) *gin.Engine {
	router := gin.Default()

	router.GET("/referential", config.ReferentialWebService.GetReferential())
	router.POST("/post", config.PokedexWebService.Post())

	return router
}

func NewConfiguration(repo gateway.Repo, referentialController controller.PokemonReferentialController[presenter.HttpResponse[core.PaginationResult[domain.Pokemon]]], pokedexController controller.PokedexController[presenter.HttpResponse[domain.PokemonsPlayer]]) Configuration {
	return Configuration{
		Repo:                  repo,
		ReferentialWebService: ReferentialWebService{ReferentialController: &referentialController},
		PokedexWebService:     PokedexWebService{PokedexController: &pokedexController},
	}
}

type Configuration struct {
	Repo gateway.Repo
	ReferentialWebService
	PokedexWebService
}
