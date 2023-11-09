package usecase

import (
	"fww-core/internal/adapter"
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/repository"
)

type useCase struct {
	repository repository.Repository
	adapter    adapter.Adapter
}

type UseCase interface {
	RegisterPassanger(data *dto_passanger.RequestRegister) (dto_passanger.ResponseRegistered, error)
	DetailPassanger(id int64) (dto_passanger.ResponseDetail, error)
	UpdatePassanger(data *dto_passanger.RequestUpdate) (dto_passanger.ResponseUpdate, error)
}

func New(repository repository.Repository, adapter adapter.Adapter) UseCase {
	return &useCase{
		repository: repository,
		adapter:    adapter,
	}
}
