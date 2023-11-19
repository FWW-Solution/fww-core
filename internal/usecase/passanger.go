package usecase

import (
	"errors"
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/entity"
	"time"
)

// DetailPassanger implements UseCase.
func (u *useCase) DetailPassanger(id int64) (dto_passanger.ResponseDetail, error) {
	result, err := u.repository.FindDetailPassanger(id)
	if err != nil {
		return dto_passanger.ResponseDetail{}, err
	}

	if result.ID == 0 {
		return dto_passanger.ResponseDetail{}, errors.New("passanger not found")
	}

	response := dto_passanger.ResponseDetail{
		CovidVaccineStatus: result.CovidVaccineStatus,
		CreatedAt:          result.CreatedAt.Format("2006-01-02 15:04:05"),
		DateOfBirth:        result.DateOfBirth.Format("2006-01-02"),
		FullName:           result.FullName,
		Gender:             result.Gender,
		ID:                 result.ID,
		IDNumber:           result.IDNumber,
		IDType:             result.IDType,
		IsIDVerified:       result.IsIDVerified,
		UpdatedAt:          result.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

// RegisterPassanger implements UseCase.
func (u *useCase) RegisterPassanger(data *dto_passanger.RequestRegister) (dto_passanger.ResponseRegistered, error) {
	// convert string to time
	dateOfBirth, err := time.Parse("2006-01-02", data.DateOfBirth)
	if err != nil {
		return dto_passanger.ResponseRegistered{}, err
	}
	entity := entity.Passenger{
		ID:                 0,
		FullName:           data.FullName,
		Gender:             data.Gender,
		DateOfBirth:        dateOfBirth,
		IDNumber:           data.IDNumber,
		IDType:             data.IDType,
		CovidVaccineStatus: "",
		IsIDVerified:       false,
	}
	result, err := u.repository.RegisterPassanger(&entity)
	if err != nil {
		return dto_passanger.ResponseRegistered{}, err
	}
	dataPassanger := dto_passanger.RequestBPM{
		IDNumber: data.IDNumber,
	}
	// Check Passanger Information
	err = u.adapter.CheckPassangerInformations(&dataPassanger)

	if err != nil {
		return dto_passanger.ResponseRegistered{}, err
	}

	return dto_passanger.ResponseRegistered{
		ID: result,
	}, nil

}

// UpdatePassanger implements UseCase.
func (u *useCase) UpdatePassanger(data *dto_passanger.RequestUpdate) (dto_passanger.ResponseUpdate, error) {
	// select data from database
	result, err := u.repository.FindDetailPassanger(data.ID)
	if err != nil {
		return dto_passanger.ResponseUpdate{}, err
	}
	if result.ID == 0 {
		return dto_passanger.ResponseUpdate{}, errors.New("passanger not found")
	}

	// update result partially if data request is not empty
	if data.FullName != "" {
		result.FullName = data.FullName
	}
	if data.Gender != "" {
		result.Gender = data.Gender
	}
	if data.DateOfBirth != "" {
		dateOfBirth, err := time.Parse("2006-01-02", data.DateOfBirth)
		if err != nil {
			return dto_passanger.ResponseUpdate{}, err
		}
		result.DateOfBirth = dateOfBirth
	}
	if data.IDNumber != "" {
		result.IDNumber = data.IDNumber
	}
	if data.IDType != "" {
		result.IDType = data.IDType
	}

	resultUpdate, err := u.repository.UpdatePassanger(&result)
	if err != nil {
		return dto_passanger.ResponseUpdate{}, err
	}

	response := dto_passanger.ResponseUpdate{
		ID: resultUpdate,
	}

	return response, nil
}

// UpdatePassangerByIDNumber implements UseCase.
func (u *useCase) UpdatePassangerByIDNumber(data *dto_passanger.RequestUpdateBPM) error {
	// find passanger by id number

	result, err := u.repository.FindPassangerByIDNumber(data.IDNumber)
	if err != nil {
		return err
	}

	if result.ID == 0 {
		return errors.New("passanger not found")
	}

	// update result partially if data request is not empty
	if data.VaccineStatus != "" {
		result.CovidVaccineStatus = data.VaccineStatus
	}
	if data.IsVerifiedDukcapil {
		result.IsIDVerified = true
	}
	if data.CaseID != 0 {
		result.CaseID = data.CaseID
	}

	_, err = u.repository.UpdatePassanger(&result)
	if err != nil {
		return err
	}

	return nil
}
