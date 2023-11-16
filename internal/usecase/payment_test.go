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
			PaymentExpiredAt: timeNow,
			Status:           "success",
		}

		entityPayment := entity.Payment{
			ID:            1,
			InvoiceNumber: paymentCodeUUID,
			TotalPayment:  100000,
			PaymentMethod: "bank_transfer",
			PaymentDate:   time.Now(),
			PaymentStatus: "success",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Time{},
			DeletedAt:     &time.Time{},
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
		bookingID := int64(1)

		request := dto_payment.Request{
			BookingID:     bookingID,
			PaymentMethod: "bank_transfer",
		}

		// entityPayment := entity.Payment{
		// 	ID:            1,
		// 	InvoiceNumber: paymentCodeUUID,
		// 	TotalPayment:  100000,
		// 	PaymentMethod: "bank_transfer",
		// 	PaymentDate:   time.Now(),
		// 	PaymentStatus: "success",
		// 	CreatedAt:     time.Now(),
		// 	UpdatedAt:     time.Time{},
		// 	DeletedAt:     &time.Time{},
		// 	BookingID:     1,
		// }

		entityBooking := entity.Booking{
			ID:               bookingID,
			CodeBooking:      codeBookingUUID,
			BookingDate:      time.Now(),
			PaymentExpiredAt: time.Now(),
			BookingStatus:    "pending",
			CaseID:           0,
			CreatedAt:        time.Now(),
			UpdatedAt:        sql.NullTime{},
			DeletedAt:        sql.NullTime{},
			UserID:           1,
			FlightID:         1,
		}

		repositoryMock.On("FindBookingDetailByID", request.BookingID).Return(entityBooking, nil).Once()
		repositoryMock.On("UpdatePayment", paymentCodeUUID).Return(nil, nil).Once()

		err := uc.RequestPayment(&request, paymentCodeUUID)
		assert.Nil(t, err)

	})
}
