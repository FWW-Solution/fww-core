package usecase

import (
	"errors"
	"fww-core/internal/data/dto_notification"
)

// InquiryNotification implements UseCase.
func (u *useCase) InquiryNotification(data *dto_notification.Request) error {
	switch data.Route {
	case "send_invoice":
		result, err := u.repository.PaymentInvoiceReportByBookingCode(data.CodeBooking)
		if err != nil {
			return err
		}

		// TODO: Populate data to template
		// spec := dto_notification.ModelInvoice{}

		u.adapter.SendNotification(result)

	case "send_receipt":
		result, err := u.repository.PaymentReceiptReportByBookingCode(data.CodeBooking)
		if err != nil {
			return err
		}

		// TODO: Populate data to template
		// spec := dto_notification.ModelPaymentReceipt{}

		u.adapter.SendNotification(result)

	case "send_ticket":
		result, err := u.repository.TicketRedeemedReportByBookingCode(data.CodeBooking)
		if err != nil {
			return err
		}

		// TODO: Populate data to template
		// spec := dto_notification.ModelTicketRedeemed{}

		u.adapter.SendNotification(result)
	default:
		return errors.New("route not found")

	}
	return nil
}
