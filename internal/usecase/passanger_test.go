package usecase_test

import (
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDetailPassanger(t *testing.T) {
	setup()
	t.Run("Success", func(t *testing.T) {
		id := int64(1)
		expected := dto_passanger.ResponseDetail{
			CovidVaccineStatus: "VACCINATED I",
			CreatedAt:          time.Now().Round(time.Minute).Format("2006-01-02 15:04:05"),
			DateOfBirth:        dateTime,
			FullName:           "John Doe",
			Gender:             "Male",
			ID:                 id,
			IDNumber:           "1234567890",
			IDType:             "KTP",
			IsIDVerified:       true,
			UpdatedAt:          time.Now().Round(time.Minute).Format("2006-01-02 15:04:05"),
		}

		entity := entity.Passenger{
			CovidVaccineStatus: "VACCINATED I",
			CreatedAt:          time.Now().Round(time.Minute),
			DateOfBirth:        time.Now().Round(time.Minute),
			FullName:           "John Doe",
			Gender:             "Male",
			ID:                 id,
			IDNumber:           "1234567890",
			IDType:             "KTP",
			IsIDVerified:       true,
			UpdatedAt:          time.Now().Round(time.Minute),
		}
		repositoryMock.On("FindDetailPassanger", id).Return(entity, nil)

		res, err := uc.DetailPassanger(id)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res != expected {
			t.Errorf("expected %v, got %v", expected, res)
		}
	})
}

func TestRegisterPassanger(t *testing.T) {
	setup()
	t.Run("Success", func(t *testing.T) {
		req := &dto_passanger.RequestRegister{
			DateOfBirth: dateTime,
			FullName:    "John Doe",
			Gender:      "Male",
			IDNumber:    "1234567890",
			IDType:      "KTP",
		}

		entity := &entity.Passenger{
			ID:                 0,
			FullName:           "John Doe",
			Gender:             "Male",
			IDNumber:           "1234567890",
			IDType:             "KTP",
			CovidVaccineStatus: "",
			DateOfBirth:        dateOnly,
			IsIDVerified:       false,
		}
		dataPassanger := &dto_passanger.RequestBPM{
			IDNumber: entity.IDNumber,
		}

		expected := dto_passanger.ResponseRegistered{
			ID: 1,
		}
		id := int64(1)
		repositoryMock.On("RegisterPassanger", entity).Return(id, nil)
		adapterMock.On("CheckPassangerInformations", dataPassanger).Return(nil)

		res, err := uc.RegisterPassanger(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res != expected {
			t.Errorf("expected %v, got %v", expected, res)
		}
	})
}

func TestUpdatePassanger(t *testing.T) {
	setup()
	t.Run("Sucess", func(t *testing.T) {
		id := int64(1)
		req := &dto_passanger.RequestUpdate{
			DateOfBirth: dateTime,
			FullName:    "John Doe",
			Gender:      "Male",
			ID:          id,
			IDNumber:    "1234567890",
			IDType:      "KTP",
		}

		entity := entity.Passenger{
			CovidVaccineStatus: "VACCINATED I",
			CreatedAt:          time.Now(),
			DateOfBirth:        dateOnly,
			FullName:           "John Doe",
			Gender:             "Male",
			ID:                 id,
			IDNumber:           "1234567890",
			IDType:             "KTP",
			IsIDVerified:       true,
			UpdatedAt:          time.Now(),
		}
		expected := dto_passanger.ResponseUpdate{
			ID: id,
		}

		repositoryMock.On("FindDetailPassanger", id).Return(entity, nil)
		repositoryMock.On("UpdatePassanger", &entity).Return(id, nil)

		result, err := uc.UpdatePassanger(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}

func TestUpdatePassangerByIDNumber(t *testing.T) {
	setup()
	t.Run("Success", func(t *testing.T) {
		idNumber := "1234567890"
		req := &dto_passanger.RequestUpdateBPM{
			IDNumber:           idNumber,
			VaccineStatus:      "VACCINATED I",
			IsVerifiedDukcapil: true,
			CaseID:             123,
		}

		entityPassenger := entity.Passenger{
			ID:                 1,
			FullName:           "Koko",
			Gender:             "Male",
			DateOfBirth:        time.Time{},
			IDNumber:           idNumber,
			IDType:             "KTP",
			CovidVaccineStatus: "",
			IsIDVerified:       false,
			CaseID:             0,
		}

		reqUpdateEntity := entity.Passenger{
			ID:                 1,
			FullName:           "Koko",
			Gender:             "Male",
			DateOfBirth:        time.Time{},
			IDNumber:           idNumber,
			IDType:             "KTP",
			CovidVaccineStatus: req.VaccineStatus,
			IsIDVerified:       req.IsVerifiedDukcapil,
			CaseID:             req.CaseID,
		}

		repositoryMock.On("FindPassangerByIDNumber", idNumber).Return(entityPassenger, nil)
		repositoryMock.On("UpdatePassanger", &reqUpdateEntity).Return(int64(1), nil)

		err := uc.UpdatePassangerByIDNumber(req)
		assert.Nil(t, err)
	})
}
