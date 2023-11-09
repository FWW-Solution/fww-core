package usecase

import (
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/repository"
)

type useCase struct {
	repository repository.Repository
}

// DetailPassanger implements UseCase.
func (*useCase) DetailPassanger(id int64) (dto_passanger.ResponseDetail, error) {
	panic("unimplemented")
}

// RegisterPassanger implements UseCase.
func (*useCase) RegisterPassanger(data dto_passanger.RequestRegister) (dto_passanger.ResponseRegistered, error) {
	panic("unimplemented")
}

// UpdatePassanger implements UseCase.
func (*useCase) UpdatePassanger(data dto_passanger.RequestUpdate) (dto_passanger.ResponseUpdate, error) {
	panic("unimplemented")
}

type UseCase interface {
	RegisterPassanger(data dto_passanger.RequestRegister) (dto_passanger.ResponseRegistered, error)
	DetailPassanger(id int64) (dto_passanger.ResponseDetail, error)
	UpdatePassanger(data dto_passanger.RequestUpdate) (dto_passanger.ResponseUpdate, error)
}

func New(repository repository.Repository) UseCase {
	return &useCase{
		repository: repository,
	}
}
