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

	// Total payment from booking details and flight price
	totalPayment := float64(0)

	bookingDetails, err := u.repository.FindBookingDetailByBookingID(resultBooking.ID)
	if err != nil {
		return err
	}

	bookingPrice, err := u.repository.FindFlightPriceByID(resultBooking.FlightID)
	if err != nil {
		return err
	}

	totalPayment += (bookingPrice.Price * float64(len(bookingDetails)))

	// Validate payment expired

	if resultBooking.PaymentExpiredAt.Before(time.Now()) {
		return errors.New("payment expired")
	}

	// TODO: validate payment method

	paymentMethods, err := u.repository.FindPaymentMethodStatus()
	if err != nil {
		return err
	}

	isValid := false
	for _, v := range paymentMethods {
		if v.Name == req.PaymentMethod && v.IsActive {
			isValid = true
		}
	}

	if !isValid {
		return errors.New("payment method not found / not active")
	}

	// TODO: Do payment process

	entityPayment := entity.Payment{
		InvoiceNumber: paymentCodeID,
		TotalPayment:  totalPayment,
		PaymentMethod: req.PaymentMethod,
		PaymentDate:   time.Now().Round(time.Second),
		PaymentStatus: "pending",
		BookingID:     resultBooking.ID,
	}
	u.adapter.RequestPayment(entityPayment)

	// TODO: Create callback url (optional)

	// TODO: Update database payment status

	_, err = u.repository.UpsertPayment(&entityPayment)
	if err != nil {
		return err
	}

	// TODO: Send  payment receipt to user (email) (async)

	u.adapter.SendNotification(entityPayment)

	return nil
}

// GetDetailPayment implements UseCase.
func (u *useCase) GetPaymentStatus(codePayment string) (dto_payment.StatusResponse, error) {
	result, err := u.repository.FindPaymentDetailByInvoice(codePayment)
	if err != nil {
		return dto_payment.StatusResponse{}, err
	}

	if result.ID == 0 {
		return dto_payment.StatusResponse{}, errors.New("payment not found")
	}

	return dto_payment.StatusResponse{
		Status: result.PaymentStatus,
	}, nil

}

// GetPaymentMethod implements UseCase.
func (u *useCase) GetPaymentMethod() ([]dto_payment.MethodResponse, error) {
	result, err := u.repository.FindPaymentMethodStatus()
	if err != nil {
		return []dto_payment.MethodResponse{}, err
	}

	var response []dto_payment.MethodResponse

	for _, v := range result {
		response = append(response, dto_payment.MethodResponse{
			ID:       v.ID,
			IsActive: v.IsActive,
			Name:     v.Name,
		})
	}

	return response, nil
}
