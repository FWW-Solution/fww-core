package usecase_test

import (
	"database/sql"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/data/dto_ticket"
	"fww-core/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRedeemTicket(t *testing.T) {
	setup()
	t.Run("success", func(t *testing.T) {
		codeBookingUUID := "123e4567-e89b-12d3-a456-426614174000"
		// codeTicketUUID := "123e4567-e89b-12d3-a456-2341235r31324"

		ticketID64 := int64(1)

		timePayment := time.Now().Add(time.Hour * 6)

		entityBooking := entity.Booking{
			ID:               1,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now(),
			PaymentExpiredAt: timePayment,
			BookingExpiredAt: timePayment.Add(time.Hour * 6),
			BookingStatus:    "redeemed",
			CaseID:           0,
			CreatedAt:        time.Now(),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		entityBookingDetails := []entity.BookingDetail{
			{
				ID:          1,
				BookingID:   1,
				PassengerID: 1,
				SeatNumber:  "1A",
				Class:       "economy",
			},
		}

		entityPassenger := entity.Passenger{
			ID:       1,
			IDNumber: "1234567890",
			IDType:   "KTP",
			FullName: "John Doe",
		}

		boardingTime := entityBooking.BookingExpiredAt.Add((24 * time.Hour) - (time.Minute * 30))

		expect := dto_ticket.Response{
			BordingTime: boardingTime.Format("2006-01-02 15:04:05"),
		}

		expectBordingTicket := expect.BordingTime

		repositoryMock.On("FindBookingByBookingIDCode", codeBookingUUID).Return(entityBooking, nil)
		// Validate Booking date
		repositoryMock.On("UpdateBooking", &entityBooking).Return(entityBooking.ID, nil)
		// Check Peduli Lindungi
		repositoryMock.On("UpsertTicket", mock.Anything).Return(ticketID64, nil)

		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetails, nil)

		repositoryMock.On("FindDetailPassanger", entityBookingDetails[0].ID).Return(entityPassenger, nil).Once()

		adapterMock.On("RedeemTicket", mock.Anything).Return(nil)

		result, err := uc.RedeemTicket(codeBookingUUID)

		assert.NoError(t, err)
		assert.Equal(t, expectBordingTicket, result.BordingTime)
	})

	t.Run("Error FindBookingByBookingIDCode", func(t *testing.T) {
		codeBookingUUID := "123e4567-e89b-12d3-a456-426614174000"
		expected := dto_ticket.Response{}

		repositoryMock.On("FindBookingByBookingIDCode", codeBookingUUID).Return(entity.Booking{}, sql.ErrNoRows)
		repositoryMock.On("FindDetailPassanger", mock.Anything).Return(entity.Passenger{}, sql.ErrNoRows).Once()

		result, err := uc.RedeemTicket(codeBookingUUID)
		assert.NotNil(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Error UpdateBooking", func(t *testing.T) {
		codeBookingUUID := "123e4567-e89b-12d3-a456-426614174000"
		expected := dto_ticket.Response{}

		entityBooking := entity.Booking{
			ID:               1,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now(),
			PaymentExpiredAt: time.Now(),
			BookingExpiredAt: time.Now(),
			BookingStatus:    "redeemed",
			CaseID:           0,
			CreatedAt:        time.Now(),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		repositoryMock.On("FindBookingByBookingIDCode", codeBookingUUID).Return(entityBooking, nil)
		repositoryMock.On("UpdateBooking", &entityBooking).Return(int64(0), sql.ErrNoRows)
		repositoryMock.On("FindDetailPassanger", mock.Anything).Return(entity.Passenger{}, sql.ErrNoRows).Once()

		result, err := uc.RedeemTicket(codeBookingUUID)
		assert.NotNil(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Error UpsertTicket", func(t *testing.T) {
		codeBookingUUID := "123e4567-e89b-12d3-a456-426614174000"
		expected := dto_ticket.Response{}

		entityBooking := entity.Booking{
			ID:               1,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now(),
			PaymentExpiredAt: time.Now(),
			BookingExpiredAt: time.Now(),
			BookingStatus:    "redeemed",
			CaseID:           0,
			CreatedAt:        time.Now(),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		repositoryMock.On("FindBookingByBookingIDCode", codeBookingUUID).Return(entityBooking, nil)
		repositoryMock.On("UpdateBooking", &entityBooking).Return(entityBooking.ID, nil)
		repositoryMock.On("UpsertTicket", mock.Anything).Return(int64(0), sql.ErrNoRows)
		repositoryMock.On("FindDetailPassanger", mock.Anything).Return(entity.Passenger{}, sql.ErrNoRows).Once()

		result, err := uc.RedeemTicket(codeBookingUUID)
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, expected, result)
	})

	t.Run("Error FindBookingDetailByBookingID", func(t *testing.T) {
		codeBookingUUID := "123e4567-e89b-12d3-a456-426614174000"
		expected := dto_ticket.Response{}

		entityBooking := entity.Booking{
			ID:               1,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now(),
			PaymentExpiredAt: time.Now(),
			BookingExpiredAt: time.Now(),
			BookingStatus:    "redeemed",
			CaseID:           0,
			CreatedAt:        time.Now(),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		repositoryMock.On("FindBookingByBookingIDCode", codeBookingUUID).Return(entityBooking, nil)
		repositoryMock.On("UpdateBooking", &entityBooking).Return(entityBooking.ID, nil)
		repositoryMock.On("UpsertTicket", mock.Anything).Return(int64(1), nil)
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return([]entity.BookingDetail{}, sql.ErrNoRows)
		repositoryMock.On("FindDetailPassanger", mock.Anything).Return(entity.Passenger{}, sql.ErrNoRows).Once()

		result, err := uc.RedeemTicket(codeBookingUUID)
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, expected, result)
	})

	t.Run("Error FindDetailPassanger", func(t *testing.T) {
		codeBookingUUID := "123e4567-e89b-12d3-a456-426614174000"

		entityBooking := entity.Booking{
			ID:               1,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now(),
			PaymentExpiredAt: time.Now(),
			BookingExpiredAt: time.Now(),
			BookingStatus:    "redeemed",
			CaseID:           0,
			CreatedAt:        time.Now(),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		entityBookingDetails := []entity.BookingDetail{
			{
				ID:          1,
				BookingID:   1,
				PassengerID: 1,
				SeatNumber:  "1A",
				Class:       "economy",
			},
		}

		repositoryMock.On("FindBookingByBookingIDCode", codeBookingUUID).Return(entityBooking, nil)
		repositoryMock.On("UpdateBooking", &entityBooking).Return(entityBooking.ID, nil)
		repositoryMock.On("UpsertTicket", mock.Anything).Return(int64(1), nil)
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetails, nil)
		repositoryMock.On("FindDetailPassanger", entityBookingDetails[0].ID).Return(entity.Passenger{}, sql.ErrNoRows).Once()

		_, err := uc.RedeemTicket(codeBookingUUID)
		if err != nil {
			assert.Error(t, err)
		}
	})

	t.Run("Error RedeemTicket", func(t *testing.T) {
		codeBookingUUID := "123e4567-e89b-12d3-a456-426614174000"

		entityBooking := entity.Booking{
			ID:               1,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now(),
			PaymentExpiredAt: time.Now(),
			BookingExpiredAt: time.Now(),
			BookingStatus:    "redeemed",
			CaseID:           0,
			CreatedAt:        time.Now(),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		entityBookingDetails := []entity.BookingDetail{
			{
				ID:          1,
				BookingID:   1,
				PassengerID: 1,
				SeatNumber:  "1A",
				Class:       "economy",
			},
		}

		entityPassenger := entity.Passenger{
			ID:       1,
			IDNumber: "1234567890",
			IDType:   "KTP",
			FullName: "John Doe",
		}

		repositoryMock.On("FindBookingByBookingIDCode", codeBookingUUID).Return(entityBooking, nil)
		repositoryMock.On("UpdateBooking", &entityBooking).Return(entityBooking.ID, nil)
		repositoryMock.On("UpsertTicket", mock.Anything).Return(int64(1), nil)
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetails, nil)
		repositoryMock.On("FindDetailPassanger", entityBookingDetails[0].ID).Return(entityPassenger, nil).Once()

		adapterMock.On("RedeemTicket", mock.Anything).Return(sql.ErrNoRows)

		_, err := uc.RedeemTicket(codeBookingUUID)
		if err != nil {
			assert.Error(t, err)
		}
	})
}

func TestUpdateTicket(t *testing.T) {
	setup()
	t.Run("success", func(t *testing.T) {
		// codeBookingUUID := "123e4567-e89b-12d3-a456-426614174000"
		codeTicketUUID := "123e4567-e89b-12d3-a456-2341235r31324"
		req := dto_ticket.RequestUpdateTicket{
			CodeTicket:         codeTicketUUID,
			BookingDetailID:    1,
			IsEligibleToFlight: false,
		}

		entityTicket := entity.Ticket{
			ID:                 1,
			CodeTicket:         codeTicketUUID,
			IsBoardingPass:     false,
			IsEligibleToFlight: false,
			BookingID:          1,
		}

		specUpdate := dto_booking.BookDetailRequest{
			BookingDetailID:    req.BookingDetailID,
			IsEligibleToFlight: req.IsEligibleToFlight,
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
			ID:                 specUpdate.BookingDetailID,
			PassengerID:        1,
			SeatNumber:         "A1",
			BaggageCapacity:    20,
			Class:              "economy",
			IsEligibleToFlight: specUpdate.IsEligibleToFlight,
			CreatedAt:          time.Now().Round(time.Minute),
			BookingID:          1,
		}

		repositoryMock.On("FindTicketByCodeTicket", codeTicketUUID).Return(entityTicket, nil)
		repositoryMock.On("FindBookingDetailByID", req.BookingDetailID).Return(entityBookingDetail, nil)
		repositoryMock.On("UpdateBookingDetail", entityRequest).Return(int64(1), nil)

		err := uc.UpdateTicket(&req)
		assert.NoError(t, err)
	})

	t.Run("Error FindTicketByCodeTicket", func(t *testing.T) {
		codeTicketUUID := ""
		req := dto_ticket.RequestUpdateTicket{
			CodeTicket:         codeTicketUUID,
			BookingDetailID:    1,
			IsEligibleToFlight: false,
		}

		repositoryMock.On("FindTicketByCodeTicket", codeTicketUUID).Return(entity.Ticket{}, sql.ErrNoRows)

		err := uc.UpdateTicket(&req)
		assert.NotNil(t, err)
	})

	t.Run("Error UpdateDetailBooking", func(t *testing.T) {
		codeTicketUUID := ""
		req := dto_ticket.RequestUpdateTicket{
			CodeTicket:         codeTicketUUID,
			BookingDetailID:    1,
			IsEligibleToFlight: false,
		}

		entityTicket := entity.Ticket{
			ID:                 1,
			CodeTicket:         codeTicketUUID,
			IsBoardingPass:     false,
			IsEligibleToFlight: false,
			BookingID:          1,
		}

		specUpdate := dto_booking.BookDetailRequest{
			BookingDetailID:    req.BookingDetailID,
			IsEligibleToFlight: req.IsEligibleToFlight,
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
			ID:                 specUpdate.BookingDetailID,
			PassengerID:        1,
			SeatNumber:         "A1",
			BaggageCapacity:    20,
			Class:              "economy",
			IsEligibleToFlight: specUpdate.IsEligibleToFlight,
			CreatedAt:          time.Now().Round(time.Minute),
			BookingID:          1,
		}

		repositoryMock.On("FindTicketByCodeTicket", codeTicketUUID).Return(entityTicket, nil)
		repositoryMock.On("FindBookingDetailByID", req.BookingDetailID).Return(entityBookingDetail, nil)
		repositoryMock.On("UpdateBookingDetail", entityRequest).Return(int64(0), sql.ErrNoRows)

		err := uc.UpdateTicket(&req)
		assert.NotNil(t, err)
	})

	t.Run("Error FindTicketByCodeTicket Return ID 0", func(t *testing.T) {
		codeTicketUUID := ""
		req := dto_ticket.RequestUpdateTicket{
			CodeTicket:         codeTicketUUID,
			BookingDetailID:    1,
			IsEligibleToFlight: false,
		}

		entityTicket := entity.Ticket{
			ID:                 0,
			CodeTicket:         codeTicketUUID,
			IsBoardingPass:     false,
			IsEligibleToFlight: false,
			BookingID:          1,
		}

		repositoryMock.On("FindTicketByCodeTicket", codeTicketUUID).Return(entityTicket, nil)

		err := uc.UpdateTicket(&req)
		assert.NotNil(t, err)
	})
}
