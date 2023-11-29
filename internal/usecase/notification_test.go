package usecase_test

import (
	"database/sql"
	"fww-core/internal/data/dto_notification"
	"fww-core/internal/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInquiryNotification(t *testing.T) {
	setup()
	t.Run("success send_invoice", func(t *testing.T) {
		req := dto_notification.Request{
			CodeBooking: "123e4567-e89b-12d3-a456-426614174000",
			Route:       "send_invoice",
		}

		entityInvoiceAggregator := dto_notification.PaymentInvoiceAggregator{
			Payment: entity.Payment{
				ID:            0,
				InvoiceNumber: "",
				TotalPayment:  0,
				PaymentMethod: "",
				PaymentDate:   time.Time{},
				PaymentStatus: "",
				CreatedAt:     time.Time{},
				UpdatedAt:     sql.NullTime{},
				DeletedAt:     sql.NullTime{},
				BookingID:     0,
			},
			BookingDetails: []entity.BookingDetail{},
			PaymentMethods: []entity.PaymentMethod{},
			Passengers:     []entity.Passenger{},
			Booking:        entity.Booking{},
			User:           entity.User{},
		}

		templateSendInvoice := "\n\t\t<html>\n\t\t\t<head>\n\t\t\t\t<title>Invoice</title>\n\t\t\t\t</head>\n\t\t\t\t<body>\n\t\t\t\t\t<h1>Invoice</h1>\n\t\t\t\t\t<p>Invoice Number: </p>\n\t\t\t\t\t<p>Booking Code: </p>\n\t\t\t\t\t<p>Payment Amount: 0</p>\n\t\t\t\t\t<p>Passenger Details:</p>\n\t\t\t\t\t<ul>\n\t\t\t\t\t\t\n\t\t\t\t\t\t\t</ul>\n\t\t\t\t\t\t\t</body>\n\t\t\t\t\t\t\t</html>\n\t\t\t\t\t\t\t"

		specNotification := dto_notification.SendEmailRequest{
			To:      entityInvoiceAggregator.User.Email,
			Subject: "[FWW] Invoice",
			Body:    templateSendInvoice,
		}

		repositoryMock.On("PaymentInvoiceReportByBookingCode", req.CodeBooking).Return(entityInvoiceAggregator, nil).Once()
		adapterMock.On("SendNotification", &specNotification).Return(nil).Once()

		err := uc.InquiryNotification(&req)
		assert.Nil(t, err)

	})

	t.Run("success send_receipt", func(t *testing.T) {
		req := dto_notification.Request{
			CodeBooking: "123e4567-e89b-12d3-a456-426614174000",
			Route:       "send_receipt",
		}

		entityPaymentReceiptAggregator := dto_notification.PaymentReceiptAggregator{
			Payment:        entity.Payment{},
			BookingDetails: []entity.BookingDetail{},
			Booking:        entity.Booking{},
			User:           entity.User{},
		}

		templateSendReceipt := "\n\t\t<html>\n\t\t\t<head>\n\t\t\t\t<title>Receipt</title>\n\t\t\t\t</head>\n\t\t\t\t<body>\n\t\t\t\t\t<h1>Receipt</h1>\n\t\t\t\t\t<p>Invoice Number: </p>\n\t\t\t\t\t<p>Booking Code: </p>\n\t\t\t\t\t<p>Payment Method: </p>\n\t\t\t\t\t<p>Payment Amount: 0</p>\n\t\t\t\t\t<p>Payment Date: 0001-01-01 00:00:00 +0000 UTC</p>\n\t\t\t\t\t</body>\n\t\t\t\t\t</html>\n\t\t\t\t\t"

		specNotification := dto_notification.SendEmailRequest{
			To:      entityPaymentReceiptAggregator.User.Email,
			Subject: "[FWW] Receipt",
			Body:    templateSendReceipt,
		}

		repositoryMock.On("PaymentReceiptReportByBookingCode", req.CodeBooking).Return(entityPaymentReceiptAggregator, nil).Once()
		adapterMock.On("SendNotification", &specNotification).Return(nil).Once()

		err := uc.InquiryNotification(&req)
		assert.Nil(t, err)
	})

	t.Run("success send_ticket", func(t *testing.T) {
		req := dto_notification.Request{
			CodeBooking: "123e4567-e89b-12d3-a456-426614174000",
			Route:       "send_ticket",
		}

		entityTicketRedeemAggregator := dto_notification.TicketRedeemAgregator{
			Ticket:         entity.Ticket{},
			Booking:        entity.Booking{},
			BookingDetails: []entity.BookingDetail{},
			Passengers:     []entity.Passenger{},
			Flight:         entity.Flight{},
			User:           entity.User{},
		}

		templateSendTicket := "\n\t\t<html>\n\t\t\t<head>\n\t\t\t\t<title>Ticket</title>\n\t\t\t\t</head>\n\t\t\t\t<body>\n\t\t\t\t\t<h1>Ticket</h1>\n\t\t\t\t\t<p>Ticket Code: </p>\n\t\t\t\t\t<p>Flight Number: </p>\n\t\t\t\t\t<p>Flight Departure Time: 0001-01-01 00:00:00</p>\n\t\t\t\t\t<p>Flight Arrival Time: 0001-01-01 00:00:00</p>\n\t\t\t\t\t<p>Flight Departure Airport: </p>\n\t\t\t\t\t<p>Flight Arrival Airport: </p>\n\t\t\t\t\t<p>Passenger Details:</p>\n\t\t\t\t\t<ul>\n\t\t\t\t\t\t\n\t\t\t\t\t\t\t</ul>\n\t\t\t\t\t\t\t<p>Boarding Time: 0001-01-01 00:00:00</p>\n\t\t\t\t\t\t\t</body>\n\t\t\t\t\t\t\t</html>\n\t\t\t\t\t\t\t"

		specNotification := dto_notification.SendEmailRequest{
			To:      entityTicketRedeemAggregator.User.Email,
			Subject: "[FWW] Ticket",
			Body:    templateSendTicket,
		}

		repositoryMock.On("TicketRedeemedReportByBookingCode", req.CodeBooking).Return(entityTicketRedeemAggregator, nil).Once()
		adapterMock.On("SendNotification", &specNotification).Return(nil).Once()

		err := uc.InquiryNotification(&req)
		assert.Nil(t, err)
	})

}
