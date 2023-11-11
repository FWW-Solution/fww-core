package usecase_test

import (
	"database/sql"
	"fmt"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRequestBooking(t *testing.T) {
	setup()
	t.Run("success", func(t *testing.T) {
		ID := int64(1)
		bookingIDCode := "123qwe"
		flightID := int64(1)
		reminingSeat := int64(50)
		resultReminingSeatString := fmt.Sprintf("%d", reminingSeat)
		reminingIntSeat := int(reminingSeat)
		BookDetails := []dto_booking.BookDetail{
			{
				Baggage:     20,
				SeatNumber:  "A1",
				PassangerID: 1,
				Class:       "Economy",
			},
		}

		req := &dto_booking.Request{
			FlightID:    flightID,
			BookDetails: BookDetails,
			UserID:      1,
		}

		entityBookingNull := entity.Booking{}

		bookingDate := time.Now().Round(time.Minute)
		paymentExpiredAt := time.Now().Add(time.Hour * 24).Round(time.Minute)

		entityBooking := &entity.Booking{
			FlightID:         flightID,
			CodeBooking:      bookingIDCode,
			BookingDate:      bookingDate,
			PaymentExpiredAt: paymentExpiredAt,
			BookingStatus:    "pending",
			CaseID:           0,
			UserID:           1,
		}

		entityBookingDetail := &entity.BookingDetail{
			PassengerID:     1,
			SeatNumber:      "A1",
			BaggageCapacity: 20,
			Class:           "Economy",
			BookingID:       ID,
		}

		entityReservation := &entity.FlightReservation{
			FlightID:     flightID,
			Class:        "Economy",
			ReservedSeat: 172 - (reminingIntSeat + len(BookDetails)),
			TotalSeat:    172,
			UpdatedAt: sql.NullTime{
				Time:  time.Now().Round(time.Minute),
				Valid: true,
			},
		}

		repositoryMock.On("FindBookingByBookingIDCode", mock.Anything).Return(entityBookingNull, nil)
		redisMock.ExpectGet("flight-1-seat").SetVal(resultReminingSeatString)
		repositoryMock.On("FindReminingSeat", req.FlightID).Return(reminingSeat, nil)
		repositoryMock.On("UpdateFlightReservation", entityReservation).Return(ID, nil)
		repositoryMock.On("InsertBooking", entityBooking).Return(ID, nil)
		repositoryMock.On("InsertBookingDetail", entityBookingDetail).Return(ID, nil)

		err := uc.RequestBooking(req, bookingIDCode)

		assert.NoError(t, err)

	})
}

func TestGetDetailBooking(t *testing.T) {
	setup()
	t.Run("success", func(t *testing.T) {
		ID := int64(1)
		bookingIDCode := "123qwe"
		codeFlight := "123qwe"
		flightID := int64(1)

		bookingDate := time.Now().Round(time.Minute)
		paymentExpiredAt := time.Now().Add(time.Hour * 24).Round(time.Minute)

		entityBooking := entity.Booking{
			ID:               ID,
			FlightID:         flightID,
			CodeBooking:      bookingIDCode,
			BookingDate:      bookingDate,
			PaymentExpiredAt: paymentExpiredAt,
			BookingStatus:    "pending",
			CaseID:           0,
			UserID:           1,
		}

		entityPassenger := entity.Passenger{
			CovidVaccineStatus: "VACCINATED I",
			CreatedAt:          time.Now(),
			DateOfBirth:        dateOnly,
			FullName:           "John Doe",
			Gender:             "Male",
			ID:                 ID,
			IDNumber:           "1234567890",
			IDType:             "KTP",
			IsIDVerified:       true,
			UpdatedAt:          time.Now(),
		}

		entityBookingDetails := []entity.BookingDetail{
			{
				ID:              ID,
				PassengerID:     entityPassenger.ID,
				SeatNumber:      "A1",
				BaggageCapacity: 20,
				Class:           "Economy",
				BookingID:       ID,
			},
		}

		entityFlight := entity.Flight{
			ID:                   flightID,
			CodeFlight:           codeFlight,
			DepartureTime:        time.Now(),
			ArrivalTime:          time.Now(),
			DepartureAirportName: "I Gusti Ngurah Rai International Airport",
			ArrivalAirportName:   "Soekarno-Hatta International Airport",
			DepartureAirportID:   ID,
			ArrivalAirportID:     ID,
			Status:               "On Time",
			PlaneID:              ID,
		}

		entityFlightPrice := entity.FlightPrice{
			ID:       ID,
			Price:    172,
			Class:    "Economy",
			FlightID: flightID,
		}

		// booking expired at resultFlight i day before DepartureTime
		bookingExpiredAt := entityFlight.DepartureTime.AddDate(0, 0, -1)

		// payment expired at 6 hours after booking date
		paymentExpiredAtResult := entityBooking.BookingDate.Add(time.Hour * 6)

		expected := dto_booking.BookResponse{
			ArrivalAirport:   "Soekarno-Hatta International Airport",
			ArrivalTime:      timeNow,
			BookExpiredAt:    bookingExpiredAt.Format("2006-01-02 15:04:05"),
			CodeBooking:      bookingIDCode,
			CodeFlight:       codeFlight,
			DepartureAirport: "I Gusti Ngurah Rai International Airport",
			DepartureTime:    timeNow,
			Details: []dto_booking.BookResponseDetail{
				{
					Baggage:       20,
					Class:         "Economy",
					PassangerName: "John Doe",
					Price:         172,
					SeatNumber:    "A1",
				},
			},
			ID:               ID,
			PaymentExpiredAt: paymentExpiredAtResult.Format("2006-01-02 15:04:05"),
			TotalPrice:       172,
		}

		repositoryMock.On("FindBookingByBookingIDCode", bookingIDCode).Return(entityBooking, nil)
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetails, nil)
		repositoryMock.On("FindFlightByID", flightID).Return(entityFlight, nil)
		repositoryMock.On("FindFlightPriceByID", entityFlight.ID).Return(entityFlightPrice, nil)
		repositoryMock.On("FindDetailPassanger", ID).Return(entityPassenger, nil)

		result, err := uc.GetDetailBooking(bookingIDCode)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)

	})
}
