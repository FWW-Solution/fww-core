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
		flightIDReminingSeat := fmt.Sprintf("flight-%d-seat", flightID)
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

		entityFlight := entity.Flight{
			ID:                   flightID,
			CodeFlight:           "123qwe",
			DepartureTime:        time.Now().Round(time.Minute),
			ArrivalTime:          time.Now().Round(time.Minute),
			DepartureAirportName: "Soekarno-Hatta International Airport",
			ArrivalAirportName:   "I Gusti Ngurah Rai International Airport",
			DepartureAirportID:   ID,
			ArrivalAirportID:     ID,
			Status:               "On Time",
			CreatedAt:            time.Now().Round(time.Minute),
			UpdatedAt:            sql.NullTime{},
			DeletedAt:            sql.NullTime{},
			PlaneID:              ID,
		}

		bookingDate := time.Now().Round(time.Minute)
		paymentExpiredAt := time.Now().Add(time.Hour * 6).Round(time.Minute)
		bookingExpiredAt := entityFlight.DepartureTime.AddDate(0, 0, -1).Round(time.Minute)

		entityBooking := &entity.Booking{
			FlightID:         flightID,
			CodeBooking:      bookingIDCode,
			BookingDate:      bookingDate,
			BookingExpiredAt: bookingExpiredAt,
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
			ReminingSeat: reminingIntSeat - 1,
			TotalSeat:    172,
			UpdatedAt: sql.NullTime{
				Time:  time.Now().Round(time.Minute),
				Valid: true,
			},
		}

		requestBpm := dto_booking.RequestBPM{
			CodeBooking:    "123qwe",
			PaymentExpired: paymentExpiredAt,
		}

		repositoryMock.On("FindBookingByBookingIDCode", mock.Anything).Return(entityBookingNull, nil)
		redisMock.ExpectGet(flightIDReminingSeat).SetVal(resultReminingSeatString)
		repositoryMock.On("FindReminingSeat", req.FlightID).Return(reminingSeat, nil)
		redisMock.ExpectSet(flightIDReminingSeat, reminingIntSeat-1, 0).SetVal("OK")
		repositoryMock.On("FindFlightByID", req.FlightID).Return(entityFlight, nil)
		repositoryMock.On("InsertBooking", entityBooking).Return(ID, nil)
		repositoryMock.On("UpdateFlightReservation", entityReservation).Return(ID, nil)
		repositoryMock.On("InsertBookingDetail", entityBookingDetail).Return(ID, nil).Once()
		adapterMock.On("RequestGenerateInvoice", &requestBpm).Return(nil)

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
		paymentExpiredAt := time.Now().Add(time.Hour * 6).Round(time.Minute)

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
			CreatedAt:          time.Now().Round(time.Minute),
			DateOfBirth:        dateOnly,
			FullName:           "John Doe",
			Gender:             "Male",
			ID:                 ID,
			IDNumber:           "1234567890",
			IDType:             "KTP",
			IsIDVerified:       true,
			UpdatedAt:          time.Now().Round(time.Minute),
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
			DepartureTime:        time.Now().Round(time.Minute),
			ArrivalTime:          time.Now().Round(time.Minute),
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

		expected := dto_booking.BookResponse{
			ArrivalAirport:   "Soekarno-Hatta International Airport",
			ArrivalTime:      time.Now().Round(time.Minute).Format("2006-01-02 15:04:05"),
			BookExpiredAt:    bookingExpiredAt.Round(time.Minute).Format("2006-01-02 15:04:05"),
			CodeBooking:      bookingIDCode,
			CodeFlight:       codeFlight,
			DepartureAirport: "I Gusti Ngurah Rai International Airport",
			DepartureTime:    time.Now().Round(time.Minute).Format("2006-01-02 15:04:05"),
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
			PaymentExpiredAt: entityBooking.PaymentExpiredAt.Format("2006-01-02 15:04:05"),
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

	t.Run("Error FindBookingByBookingIDCode", func(t *testing.T) {
		bookingIDCode := "123qwe"
		repositoryMock.On("FindBookingByBookingIDCode", bookingIDCode).Return(entity.Booking{}, sql.ErrConnDone)
		_, err := uc.GetDetailBooking(bookingIDCode)
		if err != nil {
			assert.Error(t, err)
		}
	})

	t.Run("Error FindBookingDetailByBookingID", func(t *testing.T) {
		ID := int64(0)
		bookingIDCode := "123qwe"

		entityBooking := entity.Booking{
			ID:            ID,
			FlightID:      ID,
			CodeBooking:   bookingIDCode,
			BookingStatus: "pending",
			CaseID:        0,
			UserID:        1,
		}

		repositoryMock.On("FindBookingByBookingIDCode", bookingIDCode).Return(entityBooking, nil)
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return([]entity.BookingDetail{}, sql.ErrNoRows)
		_, err := uc.GetDetailBooking(bookingIDCode)
		if err != nil {
			assert.Error(t, err)
		}
	})

	t.Run("Error FindFlightByID", func(t *testing.T) {
		ID := int64(0)
		bookingIDCode := "123qwe"

		entityBooking := entity.Booking{
			ID:            ID,
			FlightID:      ID,
			CodeBooking:   bookingIDCode,
			BookingStatus: "pending",
			CaseID:        0,
			UserID:        1,
		}

		entityBookingDetail := []entity.BookingDetail{
			{
				ID:              ID,
				PassengerID:     ID,
				SeatNumber:      "A1",
				BaggageCapacity: 20,
				Class:           "Economy",
				BookingID:       ID,
			},
		}

		repositoryMock.On("FindBookingByBookingIDCode", bookingIDCode).Return(entityBooking, nil)
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetail, nil)
		repositoryMock.On("FindFlightByID", entityBooking.FlightID).Return(entity.Flight{}, sql.ErrNoRows)
		_, err := uc.GetDetailBooking(bookingIDCode)
		if err != nil {
			assert.Error(t, err)
		}
	})

	t.Run("Error FindFlightPriceByID", func(t *testing.T) {
		ID := int64(0)
		bookingIDCode := "123qwe"

		entityBooking := entity.Booking{
			ID:            ID,
			FlightID:      ID,
			CodeBooking:   bookingIDCode,
			BookingStatus: "pending",
			CaseID:        0,
			UserID:        1,
		}

		entityBookingDetail := []entity.BookingDetail{
			{
				ID:              ID,
				PassengerID:     ID,
				SeatNumber:      "A1",
				BaggageCapacity: 20,
				Class:           "Economy",
				BookingID:       ID,
			},
		}

		entityFlight := entity.Flight{
			ID:                   ID,
			CodeFlight:           "123qwe",
			DepartureTime:        time.Now().Round(time.Minute),
			ArrivalTime:          time.Now().Round(time.Minute),
			DepartureAirportName: "I Gusti Ngurah Rai International Airport",
			ArrivalAirportName:   "Soekarno-Hatta International Airport",
			DepartureAirportID:   ID,
			ArrivalAirportID:     ID,
			Status:               "On Time",
			PlaneID:              ID,
		}

		repositoryMock.On("FindBookingByBookingIDCode", bookingIDCode).Return(entityBooking, nil)
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetail, nil)
		repositoryMock.On("FindFlightByID", entityBooking.FlightID).Return(entityFlight, nil)
		repositoryMock.On("FindFlightPriceByID", entityFlight.ID).Return(entity.FlightPrice{}, sql.ErrNoRows)
		_, err := uc.GetDetailBooking(bookingIDCode)
		if err != nil {
			assert.Error(t, err)
		}
	})

	t.Run("Error FindDetailPassanger", func(t *testing.T) {
		ID := int64(1)
		bookingIDCode := "123qwe"

		entityBooking := entity.Booking{
			ID:            ID,
			FlightID:      ID,
			CodeBooking:   bookingIDCode,
			BookingStatus: "pending",
			CaseID:        0,
			UserID:        1,
		}

		entityBookingDetail := []entity.BookingDetail{
			{
				ID:              ID,
				PassengerID:     ID,
				SeatNumber:      "A1",
				BaggageCapacity: 20,
				Class:           "Economy",
				BookingID:       ID,
			},
		}

		entityFlight := entity.Flight{
			ID:                   ID,
			CodeFlight:           "123qwe",
			DepartureTime:        time.Now().Round(time.Minute),
			ArrivalTime:          time.Now().Round(time.Minute),
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
			FlightID: ID,
		}

		repositoryMock.On("FindBookingByBookingIDCode", bookingIDCode).Return(entityBooking, nil)
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetail, nil)
		repositoryMock.On("FindFlightByID", entityBooking.FlightID).Return(entityFlight, nil)
		repositoryMock.On("FindFlightPriceByID", entityFlight.ID).Return(entityFlightPrice, nil)
		repositoryMock.On("FindDetailPassanger", ID).Return(entity.Passenger{}, sql.ErrNoRows).Once()
		_, err := uc.GetDetailBooking(bookingIDCode)
		if err != nil {
			assert.Error(t, err)
		}
	})
}

func TestUpdateDetailBooking(t *testing.T) {
	setup()
	t.Run("success", func(t *testing.T) {
		req := &dto_booking.BookDetailRequest{
			BookingDetailID:    1,
			IsEligibleToFlight: true,
		}

		entityBookingDetail := entity.BookingDetail{
			ID:                 1,
			PassengerID:        1,
			SeatNumber:         "A1",
			BaggageCapacity:    20,
			Class:              "economy",
			IsEligibleToFlight: false,
			CreatedAt:          time.Now().Round(time.Minute),
			BookingID:          1,
		}

		entityRequest := &entity.BookingDetail{
			ID:                 req.BookingDetailID,
			PassengerID:        1,
			SeatNumber:         "A1",
			BaggageCapacity:    20,
			Class:              "economy",
			IsEligibleToFlight: req.IsEligibleToFlight,
			CreatedAt:          time.Now().Round(time.Minute),
			BookingID:          1,
		}

		repositoryMock.On("FindBookingDetailByID", req.BookingDetailID).Return(entityBookingDetail, nil)
		repositoryMock.On("UpdateBookingDetail", entityRequest).Return(int64(1), nil)

		err := uc.UpdateDetailBooking(req)
		assert.NoError(t, err)
	})

	t.Run("Error FindBookingDetailByID", func(t *testing.T) {
		req := &dto_booking.BookDetailRequest{
			BookingDetailID:    1,
			IsEligibleToFlight: true,
		}

		repositoryMock.On("FindBookingDetailByID", req.BookingDetailID).Return(entity.BookingDetail{}, sql.ErrConnDone)
		err := uc.UpdateDetailBooking(req)
		if err != nil {
			assert.Error(t, err)
		}
	})

	t.Run("Error UpdateBookingDetail", func(t *testing.T) {
		req := &dto_booking.BookDetailRequest{
			BookingDetailID:    1,
			IsEligibleToFlight: true,
		}

		entityBookingDetail := entity.BookingDetail{
			ID:                 1,
			PassengerID:        1,
			SeatNumber:         "A1",
			BaggageCapacity:    20,
			Class:              "economy",
			IsEligibleToFlight: false,
			CreatedAt:          time.Now().Round(time.Minute),
			BookingID:          1,
		}

		entityRequest := &entity.BookingDetail{
			ID:                 req.BookingDetailID,
			PassengerID:        1,
			SeatNumber:         "A1",
			BaggageCapacity:    20,
			Class:              "economy",
			IsEligibleToFlight: req.IsEligibleToFlight,
			CreatedAt:          time.Now().Round(time.Minute),
			BookingID:          1,
		}

		repositoryMock.On("FindBookingDetailByID", req.BookingDetailID).Return(entityBookingDetail, nil)
		repositoryMock.On("UpdateBookingDetail", entityRequest).Return(int64(0), sql.ErrConnDone)
		err := uc.UpdateDetailBooking(req)
		if err != nil {
			assert.Error(t, err)
		}
	})
}

func TestUpdateBooking(t *testing.T) {
	setup()
	t.Run("success", func(t *testing.T) {
		request := &dto_booking.RequestUpdateBooking{
			CodeBooking: "123qwe",
			Status:      "success",
		}

		entityBooking := entity.Booking{
			ID:               1,
			CodeBooking:      "123qwe",
			BookingDate:      timeTimeNow,
			BookingExpiredAt: timeTimeNow,
			PaymentExpiredAt: timeTimeNow,
			BookingStatus:    "pending",
			CaseID:           123,
			UserID:           1,
			FlightID:         1,
		}

		requestBookingQuery := &entity.Booking{
			ID:               1,
			CodeBooking:      "123qwe",
			BookingDate:      timeTimeNow,
			BookingExpiredAt: timeTimeNow,
			PaymentExpiredAt: timeTimeNow,
			BookingStatus:    request.Status,
			CaseID:           123,
			UserID:           1,
			FlightID:         1,
		}

		repositoryMock.On("FindBookingByBookingIDCode", request.CodeBooking).Return(entityBooking, nil)
		repositoryMock.On("UpdateBooking", requestBookingQuery).Return(int64(1), nil)

		err := uc.UpdateBooking(request)
		assert.NoError(t, err)
	})

	t.Run("Error FindBookingByBookingIDCode", func(t *testing.T) {
		request := &dto_booking.RequestUpdateBooking{
			CodeBooking: "123qwe",
			Status:      "success",
		}
		repositoryMock.On("FindBookingByBookingIDCode", mock.Anything).Return(entity.Booking{}, sql.ErrNoRows)
		err := uc.UpdateBooking(request)
		if err != nil {
			assert.Error(t, err)
		}
	})

	t.Run("Error UpdateBooking", func(t *testing.T) {
		request := &dto_booking.RequestUpdateBooking{
			CodeBooking: "123qwe",
			Status:      "success",
		}
		entityBooking := entity.Booking{
			ID:               1,
			CodeBooking:      "123qwe",
			BookingDate:      timeTimeNow,
			BookingExpiredAt: timeTimeNow,
			PaymentExpiredAt: timeTimeNow,
			BookingStatus:    "pending",
			CaseID:           123,
			UserID:           1,
			FlightID:         1,
		}
		requestBookingQuery := &entity.Booking{
			ID:               1,
			CodeBooking:      "123qwe",
			BookingDate:      timeTimeNow,
			BookingExpiredAt: timeTimeNow,
			PaymentExpiredAt: timeTimeNow,
			BookingStatus:    request.Status,
			CaseID:           123,
			UserID:           1,
			FlightID:         1,
		}
		repositoryMock.On("FindBookingByBookingIDCode", request.CodeBooking).Return(entityBooking, nil)
		repositoryMock.On("UpdateBooking", requestBookingQuery).Return(int64(0), sql.ErrConnDone)
		err := uc.UpdateBooking(request)
		if err != nil {
			assert.Error(t, err)
		}
	})

}
