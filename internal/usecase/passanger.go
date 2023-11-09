package usecase

import "fww-core/internal/data/dto_passanger"

// DetailPassanger implements UseCase.
func (*useCase) DetailPassanger(id int64) (dto_passanger.ResponseDetail, error) {
	panic("unimplemented")
}

// RegisterPassanger implements UseCase.
func (*useCase) RegisterPassanger(data *dto_passanger.RequestRegister) (dto_passanger.ResponseRegistered, error) {
	panic("unimplemented")
}

// UpdatePassanger implements UseCase.
func (*useCase) UpdatePassanger(data *dto_passanger.RequestUpdate) (dto_passanger.ResponseUpdate, error) {
	panic("unimplemented")
}
