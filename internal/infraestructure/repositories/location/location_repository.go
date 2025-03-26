package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lleoserrano/estados_cidades_api/internal/domain/entities"
	"github.com/lleoserrano/estados_cidades_api/internal/infraestructure/repositories/location/dto"
	"net/http"
	"time"
)

type LocationRepository struct{}

func NewLocationRepository() *LocationRepository {
	return &LocationRepository{}
}

func (l *LocationRepository) GetStates() ([]entities.StateEntity, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := http.NewRequest(http.MethodGet, "https://brasilapi.com.br/api/ibge/uf/v1", nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error on request state: ", err)
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error on close response body: %v", err)
		}
	}()

	var statesResp []dto.BrazilApiStateResponse
	err = json.NewDecoder(resp.Body).Decode(&statesResp)
	if err != nil {
		fmt.Println("Error on decode response body: ", err)
		return nil, err
	}

	var states []entities.StateEntity

	for _, state := range statesResp {
		states = append(states, entities.StateEntity{
			Acronym: state.Tag,
			Name:    state.Name,
		})
	}

	return states, nil
}
