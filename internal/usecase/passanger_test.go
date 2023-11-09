package usecase_test

import (
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/mocks"
	"fww-core/internal/usecase"
	"testing"
)

var (
	uc             usecase.UseCase
	repositoryMock *mocks.Repository
	adapterMock    *mocks.Adapter
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
			CreatedAt:          "2021-10-01T00:00:00Z",
			DateOfBirth:        "1990-10-01T00:00:00Z",
			FullName:           "John Doe",
			Gender:             "Male",
			ID:                 id,
			IDNumber:           "1234567890",
			IDType:             "KTP",
			IsIDVerified:       "VERIFIED",
			UpdatedAt:          "2021-10-01T00:00:00Z",
		}
		repositoryMock.On("FindDetailPassanger", id).Return(expected, nil)

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
			DateOfBirth: "1990-10-01T00:00:00Z",
			FullName:    "John Doe",
			Gender:      "Male",
			IDNumber:    "1234567890",
			IDType:      "KTP",
		}
		expected := dto_passanger.ResponseRegistered{
			ID: 1,
		}
		idInt64 := int64(1)
		repositoryMock.On("RegisterPassanger", req).Return(idInt64, nil)
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
			DateOfBirth: "1990-10-01T00:00:00Z",
			FullName:    "John Doe",
			Gender:      "Male",
			ID:          1,
			IDNumber:    "1234567890",
			IDType:      "KTP",
		}
		expected := dto_passanger.ResponseUpdate{
			CovidVaccineStatus: "VACCINATED I",
			CreatedAt:          "2021-10-01T00:00:00Z",
			DateOfBirth:        "1990-10-01T00:00:00Z",
			FullName:           "John Doe",
			Gender:             "Male",
			ID:                 id,
			IDNumber:           "1234567890",
			IDType:             "KTP",
			IsIDVerified:       "VERIFIED",
			UpdatedAt:          "2021-10-01T00:00:00Z",
		}
		repositoryMock.On("UpdatePassanger", req).Return(expected, nil)

		result, err := uc.UpdatePassanger(req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}
