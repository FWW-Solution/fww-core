package usecase_test

import (
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/entity"
	"fww-core/internal/mocks"
	"fww-core/internal/usecase"
	"testing"
	"time"
)

var (
	uc             usecase.UseCase
	repositoryMock *mocks.Repository
	adapterMock    *mocks.Adapter
	timeNow        = time.Now().Format("2006-01-02 15:04:05")
	dateTime       = time.Now().Format("2006-01-02")
	t              = time.Now()
	dateOnly       = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.UTC().Location())
)

func setup() {
	repositoryMock = &mocks.Repository{}
	adapterMock = &mocks.Adapter{}
	uc = usecase.New(repositoryMock, adapterMock)
}

func TestDetailPassanger(t *testing.T) {
	setup()
	t.Run("Success", func(t *testing.T) {
		id := int64(1)
		expected := dto_passanger.ResponseDetail{
			CovidVaccineStatus: "VACCINATED I",
			CreatedAt:          timeNow,
			DateOfBirth:        dateTime,
			FullName:           "John Doe",
			Gender:             "Male",
			ID:                 id,
			IDNumber:           "1234567890",
			IDType:             "KTP",
			IsIDVerified:       true,
			UpdatedAt:          timeNow,
		}

		entity := entity.Passenger{
			CovidVaccineStatus: "VACCINATED I",
			CreatedAt:          time.Now(),
			DateOfBirth:        time.Now(),
			FullName:           "John Doe",
			Gender:             "Male",
			ID:                 id,
			IDNumber:           "1234567890",
			IDType:             "KTP",
			IsIDVerified:       true,
			UpdatedAt:          time.Now(),
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
		expected := dto_passanger.ResponseRegistered{
			ID: 1,
		}
		id := int64(1)
		repositoryMock.On("RegisterPassanger", entity).Return(id, nil)
		adapterMock.On("CheckPassangerInformations", nil)

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
