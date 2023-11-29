package usecase_test

import (
	"database/sql"
	"fww-core/internal/data/dto_payment"
	"fww-core/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPaymentStatus(t *testing.T) {
	setup()
	t.Run("Sucess", func(t *testing.T) {
		paymentCodeUUID := "c5e8a9b0-5c0e-4a0d-9b3a-8b6da1a9b0c3c"

		expect := dto_payment.StatusResponse{
			Status: "success",
		}

		entityPayment := entity.Payment{
			ID:            1,
			InvoiceNumber: paymentCodeUUID,
			TotalPayment:  100000,
			PaymentMethod: "bank_transfer",
			PaymentDate:   time.Now(),
			PaymentStatus: "success",
			CreatedAt:     time.Now(),
			UpdatedAt:     sql.NullTime{},
			DeletedAt:     sql.NullTime{},
			BookingID:     1,
		}

		repositoryMock.On("FindPaymentDetailByInvoice", paymentCodeUUID).Return(entityPayment, nil).Once()

		result, err := uc.GetPaymentStatus(paymentCodeUUID)
		assert.Nil(t, err)
		assert.Equal(t, expect, result)

	})
}

func TestRequestPayment(t *testing.T) {
	setup()
	t.Run("Sucess", func(t *testing.T) {
		paymentCodeUUID := "c5e8a9b0-5c0e-4a0d-9b3a-8b6d1a9b0c3c"
		codeBookingUUID := "c5e8a9b0-5c0e-4a0d-9b3a-9b0c3cc5e8e9"
		// codeTicketUUID := "123e4567-e89b-12d3-a456-2341235r31324"
		bookingID := int64(1)

		request := dto_payment.Request{
			BookingID:     bookingID,
			PaymentMethod: "bca",
		}

		paymentExpiredAt := time.Now().Add(time.Hour * 6).Round(time.Minute)

		entityBooking := entity.Booking{
			ID:               bookingID,
			CodeBooking:      codeBookingUUID,
			BookingDate:      paymentExpiredAt,
			PaymentExpiredAt: paymentExpiredAt,
			BookingStatus:    "pending",
			CaseID:           0,
			CreatedAt:        time.Now().Round(time.Minute),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		entitiesPaymentMethod := []entity.PaymentMethod{
			{
				ID:        1,
				Name:      "bca",
				IsActive:  true,
				CreatedAt: time.Now().Round(time.Minute),
				UpdatedAt: sql.NullTime{},
				DeletedAt: sql.NullTime{},
			},
			{
				ID:        2,
				Name:      "bni",
				IsActive:  true,
				CreatedAt: time.Now().Round(time.Minute),
				UpdatedAt: sql.NullTime{},
				DeletedAt: sql.NullTime{},
			},
			{
				ID:        1,
				Name:      "bri",
				IsActive:  false,
				CreatedAt: time.Now().Round(time.Minute),
				UpdatedAt: sql.NullTime{},
				DeletedAt: sql.NullTime{},
			},
		}

		totalPayment := float64(100)

		entityPayment := entity.Payment{
			InvoiceNumber: paymentCodeUUID,
			TotalPayment:  totalPayment,
			PaymentMethod: request.PaymentMethod,
			PaymentDate:   time.Now().Round(time.Second),
			PaymentStatus: "pending",
			BookingID:     request.BookingID,
		}
		specDoPayment := &dto_payment.DoPayment{
			CaseID:        entityBooking.CaseID,
			InvoiceNumber: entityPayment.InvoiceNumber,
			PaymentMethod: entityPayment.PaymentMethod,
			PaymentAmount: entityPayment.TotalPayment,
		}

		repositoryMock.On("FindBookingByID", request.BookingID).Return(entityBooking, nil).Once()
		repositoryMock.On("FindPaymentMethodStatus").Return(entitiesPaymentMethod, nil).Once()
		repositoryMock.On("FindPaymentByBookingID", request.BookingID).Return(entityPayment, nil).Once()
		adapterMock.On("DoPayment", specDoPayment).Return(nil).Once()
		adapterMock.On("SendNotification", entityPayment).Return(nil).Once()

		err := uc.RequestPayment(&request)
		assert.Nil(t, err)

	})
}

func TestGenerateInvoice(t *testing.T) {
	setup()
	t.Run("Sucess", func(t *testing.T) {
		codeBookingUUID := "c5e8a9b0-5c0e-4a0d-9b3a-9b0c3ccre8a9"
		bookingID := int64(1)
		caseID := int64(1234)

		request := dto_payment.RequestInvoice{
			CaseID:      bookingID,
			CodeBooking: codeBookingUUID,
		}

		entityBooking := entity.Booking{
			ID:               bookingID,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now().Round(time.Minute),
			PaymentExpiredAt: time.Now().Add(time.Hour * 6).Round(time.Minute),
			BookingStatus:    "pending",
			CaseID:           0,
			CreatedAt:        time.Now().Round(time.Minute),
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

		flightPrice := entity.FlightPrice{
			ID:       1,
			FlightID: 1,
			Price:    100000,
			Class:    "economy",
		}

		updateBooking := entity.Booking{
			ID:               bookingID,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now().Round(time.Minute),
			PaymentExpiredAt: time.Now().Add(time.Hour * 6).Round(time.Minute),
			BookingStatus:    "pending",
			CaseID:           caseID,
			CreatedAt:        time.Now().Round(time.Minute),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		entityPayment := entity.Payment{
			InvoiceNumber: "INV-" + time.Now().Round(time.Second).Format("2006019123"),
			TotalPayment:  100000,
			PaymentMethod: "bank_transfer",
			PaymentDate:   time.Now().Round(time.Second),
			PaymentStatus: "pending",
			BookingID:     bookingID,
		}

		repositoryMock.On("FindBookingByBookingIDCode", request.CodeBooking).Return(entityBooking, nil).Once()
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetails, nil).Once()
		repositoryMock.On("FindFlightPriceByID", entityBooking.FlightID).Return(flightPrice, nil).Once()
		repositoryMock.On("UpdateBooking", &updateBooking).Return(entityPayment.ID, nil).Once()
		repositoryMock.On("UpsertPayment", mock.Anything).Return(entityPayment.ID, nil).Once()

		err := uc.GenerateInvoice(caseID, codeBookingUUID)
		assert.Nil(t, err)
	})
	t.Run("Error FindBookingByBookingIDCode", func(t *testing.T) {
		codeBookingUUID := "c5e8a9b0-5c0e-4a0d-9b3a-9b0c3yc5e8a9"
		bookingID := int64(1)
		caseID := int64(1234)

		request := dto_payment.RequestInvoice{
			CaseID:      bookingID,
			CodeBooking: codeBookingUUID,
		}

		repositoryMock.On("FindBookingByBookingIDCode", request.CodeBooking).Return(entity.Booking{}, sql.ErrNoRows).Once()

		err := uc.GenerateInvoice(caseID, codeBookingUUID)
		assert.NotNil(t, err)
	})
	t.Run("Error FindBookingDetailByBookingID", func(t *testing.T) {
		codeBookingUUID := "c5e8a9b0-5c0e-4a0d-9b3a-9b0cccc5e8a9"
		bookingID := int64(1)
		caseID := int64(1234)

		request := dto_payment.RequestInvoice{
			CaseID:      bookingID,
			CodeBooking: codeBookingUUID,
		}

		entityBooking := entity.Booking{
			ID:               bookingID,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now().Round(time.Minute),
			PaymentExpiredAt: time.Now().Add(time.Hour * 6).Round(time.Minute),
			BookingStatus:    "pending",
			CaseID:           0,
			CreatedAt:        time.Now().Round(time.Minute),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		repositoryMock.On("FindBookingByBookingIDCode", request.CodeBooking).Return(entityBooking, nil).Once()
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return([]entity.BookingDetail{}, sql.ErrNoRows).Once()

		err := uc.GenerateInvoice(caseID, codeBookingUUID)
		assert.NotNil(t, err)
	})
	t.Run("Error FindFlightPriceByID", func(t *testing.T) {
		codeBookingUUID := "c5e8a9b0-5c0e-4a0d-9b3a-9b0c3c45e8a9"
		bookingID := int64(1)
		caseID := int64(1234)

		request := dto_payment.RequestInvoice{
			CaseID:      bookingID,
			CodeBooking: codeBookingUUID,
		}

		entityBooking := entity.Booking{
			ID:               bookingID,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now().Round(time.Minute),
			PaymentExpiredAt: time.Now().Add(time.Hour * 6).Round(time.Minute),
			BookingStatus:    "pending",
			CaseID:           0,
			CreatedAt:        time.Now().Round(time.Minute),
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

		repositoryMock.On("FindBookingByBookingIDCode", request.CodeBooking).Return(entityBooking, nil).Once()
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetails, nil).Once()
		repositoryMock.On("FindFlightPriceByID", entityBooking.FlightID).Return(entity.FlightPrice{}, sql.ErrNoRows).Once()

		err := uc.GenerateInvoice(caseID, codeBookingUUID)
		assert.NotNil(t, err)
	})
	t.Run("Error UpdateBooking", func(t *testing.T) {
		codeBookingUUID := "c5e8a9b0-5c0e-4a0d-9b3a-9b0c3dc5e8a9"
		bookingID := int64(1)
		caseID := int64(1234)

		request := dto_payment.RequestInvoice{
			CaseID:      bookingID,
			CodeBooking: codeBookingUUID,
		}

		entityBooking := entity.Booking{
			ID:               bookingID,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now().Round(time.Minute),
			PaymentExpiredAt: time.Now().Add(time.Hour * 6).Round(time.Minute),
			BookingStatus:    "pending",
			CaseID:           0,
			CreatedAt:        time.Now().Round(time.Minute),
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

		flightPrice := entity.FlightPrice{
			ID:       1,
			FlightID: 1,
			Price:    100000,
			Class:    "economy",
		}

		updateBooking := entity.Booking{
			ID:               bookingID,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now().Round(time.Minute),
			PaymentExpiredAt: time.Now().Add(time.Hour * 6).Round(time.Minute),
			BookingStatus:    "pending",
			CaseID:           caseID,
			CreatedAt:        time.Now().Round(time.Minute),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		repositoryMock.On("FindBookingByBookingIDCode", request.CodeBooking).Return(entityBooking, nil).Once()
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetails, nil).Once()
		repositoryMock.On("FindFlightPriceByID", entityBooking.FlightID).Return(flightPrice, nil).Once()
		repositoryMock.On("UpdateBooking", &updateBooking).Return(int64(0), sql.ErrNoRows).Once()

		err := uc.GenerateInvoice(caseID, codeBookingUUID)
		assert.NotNil(t, err)
	})
	t.Run("Error UpsertPayment", func(t *testing.T) {
		codeBookingUUID := "c5e8a9b0-5c0e-4a0d-9b3a-9b0c3cwc5e8a9"
		bookingID := int64(1)
		caseID := int64(1234)

		request := dto_payment.RequestInvoice{
			CaseID:      bookingID,
			CodeBooking: codeBookingUUID,
		}

		entityBooking := entity.Booking{
			ID:               bookingID,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now().Round(time.Minute),
			PaymentExpiredAt: time.Now().Add(time.Hour * 6).Round(time.Minute),
			BookingStatus:    "pending",
			CaseID:           0,
			CreatedAt:        time.Now().Round(time.Minute),
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

		flightPrice := entity.FlightPrice{
			ID:       1,
			FlightID: 1,
			Price:    100000,
			Class:    "economy",
		}

		updateBooking := entity.Booking{
			ID:               bookingID,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now().Round(time.Minute),
			PaymentExpiredAt: time.Now().Add(time.Hour * 6).Round(time.Minute),
			BookingStatus:    "pending",
			CaseID:           caseID,
			CreatedAt:        time.Now().Round(time.Minute),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		entityPayment := entity.Payment{
			InvoiceNumber: "INV-" + time.Now().Round(time.Second).Format("2006019123"),
			TotalPayment:  100000,
			PaymentMethod: "bank_transfer",
			PaymentDate:   time.Now().Round(time.Second),
			PaymentStatus: "pending",
			BookingID:     bookingID,
		}

		repositoryMock.On("FindBookingByBookingIDCode", request.CodeBooking).Return(entityBooking, nil).Once()
		repositoryMock.On("FindBookingDetailByBookingID", entityBooking.ID).Return(entityBookingDetails, nil).Once()
		repositoryMock.On("FindFlightPriceByID", entityBooking.FlightID).Return(flightPrice, nil).Once()
		repositoryMock.On("UpdateBooking", &updateBooking).Return(entityPayment.ID, nil).Once()
		repositoryMock.On("UpsertPayment", mock.Anything).Return(int64(0), sql.ErrNoRows).Once()

		err := uc.GenerateInvoice(caseID, codeBookingUUID)
		assert.NotNil(t, err)
	})
}

func TestGetPaymentMethod(t *testing.T) {
	setup()
	t.Run("Sucess", func(t *testing.T) {
		expect := []dto_payment.MethodResponse{
			{
				ID:       1,
				Name:     "bca",
				IsActive: true,
			},
			{
				ID:       2,
				Name:     "bni",
				IsActive: true,
			},
			{
				ID:       3,
				Name:     "bri",
				IsActive: false,
			},
		}

		entitiesPaymentMethod := []entity.PaymentMethod{
			{
				ID:        1,
				Name:      "bca",
				IsActive:  true,
				CreatedAt: time.Now().Round(time.Minute),
				UpdatedAt: sql.NullTime{},
				DeletedAt: sql.NullTime{},
			},
			{
				ID:        2,
				Name:      "bni",
				IsActive:  true,
				CreatedAt: time.Now().Round(time.Minute),
				UpdatedAt: sql.NullTime{},
				DeletedAt: sql.NullTime{},
			},
			{
				ID:        3,
				Name:      "bri",
				IsActive:  false,
				CreatedAt: time.Now().Round(time.Minute),
				UpdatedAt: sql.NullTime{},
				DeletedAt: sql.NullTime{},
			},
		}

		repositoryMock.On("FindPaymentMethodStatus").Return(entitiesPaymentMethod, nil).Once()

		result, err := uc.GetPaymentMethod()
		assert.Nil(t, err)
		assert.Equal(t, expect, result)
	})
	t.Run("Error", func(t *testing.T) {
		entitiesPaymentMethod := []entity.PaymentMethod{}

		repositoryMock.On("FindPaymentMethodStatus").Return(entitiesPaymentMethod, sql.ErrNoRows).Once()

		_, err := uc.GetPaymentMethod()
		assert.NotNil(t, err)
	})
}

func TestUpdatePayment(t *testing.T) {
	setup()
	t.Run("Sucess", func(t *testing.T) {
		paymentCodeUUID := "c5e8a9b0-5c0e-4a0d-9b3a-8b6d1a9b0c3c"

		request := dto_payment.RequestUpdatePayment{
			InvoiceNumber: paymentCodeUUID,
			Status:        "success",
			PaymentMethod: "bank_transfer",
		}

		entityPayment := entity.Payment{
			ID:            1,
			InvoiceNumber: paymentCodeUUID,
			TotalPayment:  100000,
			PaymentMethod: "bank_transfer",
			PaymentDate:   timeTimeNow,
			PaymentStatus: "pending",
			BookingID:     1,
		}

		entityRequest := entity.Payment{
			ID:            1,
			InvoiceNumber: paymentCodeUUID,
			TotalPayment:  100000,
			PaymentMethod: request.PaymentMethod,
			PaymentDate:   time.Now().Round(time.Second),
			PaymentStatus: request.Status,
			BookingID:     1,
		}

		repositoryMock.On("FindPaymentDetailByInvoice", paymentCodeUUID).Return(entityPayment, nil).Once()
		repositoryMock.On("UpsertPayment", &entityRequest).Return(entityPayment.ID, nil).Once()

		err := uc.UpdatePayment(&request)
		assert.Nil(t, err)
	})
}
