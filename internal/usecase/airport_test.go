package usecase_test

import (
	"fww-core/internal/data/dto_airport"
	"fww-core/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAirport(t *testing.T) {
	setup()
	t.Run("Success", func(t *testing.T) {
		city := "Jakarta"
		province := "DKI Jakarta"
		iata := ""
		entityResult := []entity.Airport{
			{
				ID:        1,
				Name:      "Soekarno-Hatta International Airport",
				City:      "Tangerang",
				Province:  "Banten",
				IATA:      "CGK",
				ICAO:      "WIII",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        2,
				Name:      "Halim Perdanakusuma International Airport",
				City:      "Jakarta",
				Province:  "DKI Jakarta",
				IATA:      "HLP",
				ICAO:      "WIHH",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		expected := []dto_airport.ResponseAirport{
			{
				ID:        1,
				Name:      "Soekarno-Hatta International Airport",
				City:      "Tangerang",
				Province:  "Banten",
				Iata:      "CGK",
				Icao:      "WIII",
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
			},
			{
				ID:        2,
				Name:      "Halim Perdanakusuma International Airport",
				City:      "Jakarta",
				Province:  "DKI Jakarta",
				Iata:      "HLP",
				Icao:      "WIHH",
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
			},
		}

		repositoryMock.On("FindAirport", city, province, iata).Return(entityResult, nil)

		res, err := uc.GetAirport(city, province, iata)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expected, res)
	})
}
