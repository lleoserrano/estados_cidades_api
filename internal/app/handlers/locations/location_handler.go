package locations

import (
	"github.com/gin-gonic/gin"
	"github.com/lleoserrano/estados_cidades_api/internal/app/handlers/locations/dto"
	repositories "github.com/lleoserrano/estados_cidades_api/internal/infraestructure/repositories/location"
)

type LocationHandler struct {
	locationRepository *repositories.LocationRepository
}

func NewLocationHandler(locationRepository *repositories.LocationRepository) *LocationHandler {
	return &LocationHandler{locationRepository: locationRepository}
}

func (l *LocationHandler) GetAllStates(c *gin.Context) {
	states, err := l.locationRepository.GetStates()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	var statesResponse []dto.StateResponse
	for _, s := range states {
		statesResponse = append(statesResponse, dto.StateResponse{
			Acronym: s.Acronym,
			Name:    s.Name,
		})
	}
	c.JSON(200, statesResponse)
}
