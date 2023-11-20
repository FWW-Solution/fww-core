package usecase_test

import (
	"database/sql"
	"fww-core/internal/data/dto_payment"
	"fww-core/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPaymentStatus(t *testing.T) {
	setup()
	t.Run("Sucess", func(t *testing.T) {
		paymentCodeUUID := "c5e8a9b0-5c0e-4a0d-9b3a-8b6d1a9b0c3c"

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
		codeBookingUUID := "c5e8a9b0-5c0e-4a0d-9b3a-9b0c3cc5e8a9"
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

		entitiesBookingDetail := []entity.BookingDetail{
			{
				ID:              bookingID,
				PassengerID:     1,
				SeatNumber:      "A1",
				BaggageCapacity: 20,
				Class:           "Economy",
				CreatedAt:       time.Now().Round(time.Minute),
				UpdatedAt:       sql.NullTime{},
				DeletedAt:       sql.NullTime{},
				BookingID:       bookingID,
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
				ID:        1,
				Name:      "bri",
				IsActive:  false,
				CreatedAt: time.Now().Round(time.Minute),
				UpdatedAt: sql.NullTime{},
				DeletedAt: sql.NullTime{},
			},
		}

		entityPrice := entity.FlightPrice{
			ID:        1,
			Price:     100,
			Class:     "Economy",
			CreatedAt: time.Now().Round(time.Minute),
			UpdatedAt: sql.NullTime{},
			DeletedAt: sql.NullTime{},
			FlightID:  1,
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

		// entityTicket := entity.Ticket{
		// 	ID:                 1,
		// 	CodeTicket:         codeTicketUUID,
		// 	IsBoardingPass:     false,
		// 	IsEligibleToFlight: false,
		// 	CreatedAt:          time.Now(),
		// 	UpdatedAt:          sql.NullTime{},
		// 	DeletedAt:          sql.NullTime{},
		// 	BookingID:          1,
		// }

		repositoryMock.On("FindBookingByID", request.BookingID).Return(entityBooking, nil).Once()
		repositoryMock.On("FindBookingDetailByBookingID", request.BookingID).Return(entitiesBookingDetail, nil).Once()
		repositoryMock.On("FindFlightPriceByID", entityBooking.FlightID).Return(entityPrice, nil).Once()
		repositoryMock.On("FindPaymentMethodStatus").Return(entitiesPaymentMethod, nil).Once()
		adapterMock.On("RequestPayment", entityPayment).Return(nil).Once()
		adapterMock.On("SendNotification", entityPayment).Return(nil).Once()
		repositoryMock.On("UpsertPayment", &entityPayment).Return(entityPayment.ID, nil).Once()
		// repositoryMock.On("InsertTicket", entityTicket).Return(entityBooking.ID, nil).Once()

		err := uc.RequestPayment(&request)
		assert.Nil(t, err)

	})
}
