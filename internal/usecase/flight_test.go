package usecase_test

import (
	"database/sql"
	"fww-core/internal/data/dto_flight"
	"fww-core/internal/entity"
	"testing"
	"time"
)

func TestGetDetailFlightByID(t *testing.T) {
	setup()
	t.Run("Success", func(t *testing.T) {
		id := int64(1)
		expected := dto_flight.ResponseFlightDetail{
			ArrivalAirportName:  "Soekarno-Hatta International Airport",
			ArrivalTime:         time.Now().Round(time.Hour).Format("2006-01-02 15:04:05"),
			CodeFlight:          "GA-001",
			DepartureTime:       time.Now().Round(time.Hour).Format("2006-01-02 15:04:05"),
			DepatureAirportName: "New International Airport - Yogyakarta",
			FlightPrice:         1000000,
			ReminingSeat:        172,
			Status:              "On Time",
		}

		entityFlight := entity.Flight{
			ID:                   id,
			CodeFlight:           "GA-001",
			DepartureTime:        time.Now().Round(time.Hour),
			ArrivalTime:          time.Now().Round(time.Hour),
			DepartureAirportName: "New International Airport - Yogyakarta",
			ArrivalAirportName:   "Soekarno-Hatta International Airport",
			DepartureAirportID:   id,
			ArrivalAirportID:     id,
			Status:               "On Time",
			CreatedAt:            time.Now().Round(time.Hour),
			UpdatedAt: sql.NullTime{
				Time:  time.Now().Round(time.Hour),
				Valid: true,
			},
			DeletedAt: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			PlaneID: id,
		}

		entityFlightPrice := entity.FlightPrice{
			ID:        1,
			Price:     1000000,
			Class:     "Economy",
			CreatedAt: time.Now().Round(time.Hour),
			UpdatedAt: sql.NullTime{
				Time:  time.Now().Round(time.Hour),
				Valid: true,
			},
			DeletedAt: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			FlightID: id,
		}

		entityFlightReservation := entity.FlightReservation{
			ID:           1,
			Class:        "Economy",
			ReminingSeat: 172,
			TotalSeat:    172,
			CreatedAt:    time.Now().Round(time.Hour),
			UpdatedAt: sql.NullTime{
				Time:  time.Now().Round(time.Hour),
				Valid: true,
			},
			DeletedAt: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			FlightID: id,
		}

		repositoryMock.On("FindFlightByID", id).Return(entityFlight, nil).Once()
		repositoryMock.On("FindFlightPriceByID", id).Return(entityFlightPrice, nil).Once()
		repositoryMock.On("FindFlightReservationByID", id).Return(entityFlightReservation, nil).Once()

		result, err := uc.GetDetailFlightByID(id)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}

	})
}

func TestGetFlights(t *testing.T) {
	setup()
	t.Run("Success", func(t *testing.T) {
		departureTime := dateTime
		arrivalTime := dateTime
		limit := 5
		offset := 1
		expected := []dto_flight.ResponseFlight{
			{
				ArrivalAirportName:  "Soekarno-Hatta International Airport",
				ArrivalTime:         time.Now().Round(time.Hour).Format("2006-01-02 15:04:05"),
				CodeFlight:          "GA-001",
				DepartureTime:       time.Now().Round(time.Hour).Format("2006-01-02 15:04:05"),
				DepatureAirportName: "New International Airport - Yogyakarta",
				FlightPrice:         1000000,
				ReminingSeat:        172,
				Status:              "On Time",
			},
			{
				ArrivalAirportName:  "Soekarno-Hatta International Airport",
				ArrivalTime:         timeNow,
				CodeFlight:          "GA-001",
				DepartureTime:       timeNow,
				DepatureAirportName: "New International Airport - Yogyakarta",
				FlightPrice:         1000000,
				ReminingSeat:        172,
				Status:              "On Time",
			},
		}

		entityFlights := []entity.Flight{
			{
				ID:                   1,
				CodeFlight:           "GA-001",
				DepartureTime:        time.Now().Round(time.Hour),
				ArrivalTime:          time.Now().Round(time.Hour),
				DepartureAirportName: "New International Airport - Yogyakarta",
				ArrivalAirportName:   "Soekarno-Hatta International Airport",
				DepartureAirportID:   1,
				ArrivalAirportID:     1,
				Status:               "On Time",
				CreatedAt:            time.Now().Round(time.Hour),
				UpdatedAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				DeletedAt: sql.NullTime{
					Time:  time.Time{},
					Valid: false,
				},
				PlaneID: 1,
			},
			{
				ID:                   2,
				CodeFlight:           "GA-001",
				DepartureTime:        time.Now().Round(time.Minute),
				ArrivalTime:          time.Now().Round(time.Minute),
				DepartureAirportName: "New International Airport - Yogyakarta",
				ArrivalAirportName:   "Soekarno-Hatta International Airport",
				DepartureAirportID:   1,
				ArrivalAirportID:     1,
				Status:               "On Time",
				CreatedAt:            time.Now().Round(time.Minute),
				UpdatedAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				DeletedAt: sql.NullTime{
					Time:  time.Time{},
					Valid: false,
				},
				PlaneID: 1,
			},
		}

		repositoryMock.On("FindFlights", departureTime, arrivalTime, limit, offset).Return(entityFlights, nil).Once()

		result, err := uc.GetFlights(departureTime, arrivalTime, limit, offset)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != len(expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

}
