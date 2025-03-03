package usecase

import (
	"ProjectBot1/entity"
)

type UseCase struct {
	weatherClient weatherClient
}

type weatherClient interface {
	Get(loc string) entity.Weather
}

func NewUseCase(client weatherClient) *UseCase {
	return &UseCase{weatherClient: client}
}

func (uc *UseCase) FetchWeather(location string) entity.Weather {
	return uc.weatherClient.Get(location)
}

func (uc *UseCase) FetchTraffic(location string) entity.Weather {
	return uc.weatherClient.Get(location)
}
