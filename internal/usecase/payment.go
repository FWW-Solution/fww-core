package usecase

import (
	"errors"
	"fww-core/internal/data/dto_payment"
	"fww-core/internal/entity"
	"time"
)

// RequestPayment implements UseCase.
func (u *useCase) RequestPayment(req *dto_payment.Request, paymentCodeID string) error {
	// TODO: validate payment expiration date

	resultBooking, err := u.repository.FindBookingByID(req.BookingID)
	if err != nil {
		return err
	}

	if resultBooking.ID == 0 {
		return errors.New("booking not found")
	}

	// Validate payment expired

	if resultBooking.PaymentExpiredAt.Before(time.Now()) {
		return errors.New("payment expired")
	}

	// TODO: validate payment method

	// TODO: Do payment process

	entityPayment := entity.Payment{}
	u.adapter.RequestPayment(entityPayment)

	// TODO: Create callback url 9optional)

	// TODO: Update database payment status

	_, err = u.repository.UpdatePayment(&entityPayment)
	if err != nil {
		return err
	}

	// TODO: Send  payment receipt to user (email) (async)

	return nil
}

// GetDetailPayment implements UseCase.
func (u *useCase) GetPaymentStatus(codePayment string) (dto_payment.StatusResponse, error) {
	panic("unimplemented")
}
