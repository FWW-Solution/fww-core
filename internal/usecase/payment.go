package usecase

import "fww-core/internal/data/dto_payment"

// RequestPayment implements UseCase.
func (u *useCase) RequestPayment(req *dto_payment.Request, paymentCodeID string) error {
	panic("unimplemented")

	// TODO: validate payment expiration date

	// TODO: validate payment method

	// TODO: Do payment process

	// TODO: Create callback url 9optional)

	// TODO: Update database payment status

	// TODO: Send  payment receipt to user (email) (async)

}

// GetDetailPayment implements UseCase.
func (u *useCase) GetDetailPayment(codePayment string) (dto_payment.StatusResponse, error) {
	panic("unimplemented")
}
