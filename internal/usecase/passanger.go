package usecase

import "fww-core/internal/data/dto_passanger"

// DetailPassanger implements UseCase.
func (u *useCase) DetailPassanger(id int64) (dto_passanger.ResponseDetail, error) {
	result, err := u.repository.FindDetailPassanger(id)
	if err != nil {
		return dto_passanger.ResponseDetail{}, err
	}

	return result, nil
}

// RegisterPassanger implements UseCase.
func (u *useCase) RegisterPassanger(data *dto_passanger.RequestRegister) (dto_passanger.ResponseRegistered, error) {
	result, err := u.repository.RegisterPassanger(data)
	if err != nil {
		return dto_passanger.ResponseRegistered{}, err
	}

	// Check Passanger Information
	u.adapter.CheckPassangerInformations(nil)

	return dto_passanger.ResponseRegistered{
		ID: result,
	}, nil

}

// UpdatePassanger implements UseCase.
func (u *useCase) UpdatePassanger(data *dto_passanger.RequestUpdate) (dto_passanger.ResponseUpdate, error) {
	result, err := u.repository.UpdatePassanger(data)
	if err != nil {
		return dto_passanger.ResponseUpdate{}, err
	}

	return result, nil
}
