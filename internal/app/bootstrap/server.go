package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lleoserrano/estados_cidades_api/internal/app/handlers/locations"
	repositories "github.com/lleoserrano/estados_cidades_api/internal/infraestructure/repositories/location"
)

func StartServer() {
	e := gin.Default()
	configureRoutes(e)
	err := e.Run(":8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server started on port 8080")
}

func configureRoutes(e *gin.Engine) {
	locationRep := repositories.NewLocationRepository()
	locationHandler := locations.NewLocationHandler(locationRep)
	g := e.Group("/api/v1")
	{
		g.GET("/states", locationHandler.GetAllStates)
	}

}
