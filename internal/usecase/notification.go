package usecase

import (
	"errors"
	"fww-core/internal/data/dto_notification"
)

var (
	templateSendInvoice = `
		<html>
			<head>
				<title>Invoice</title>
				</head>
				<body>
					<h1>Invoice</h1>
					<p>Invoice Number: {{.InvoiceNumber}}</p>
					<p>Booking Code: {{.BookingCode}}</p>
					<p>Payment Method: {{.PaymentMethod}}</p>
					<p>Payment Amount: {{.PaymentAmount}}</p>
					<p>Payment Date: {{.PaymentDate}}</p>
					<p>Passenger Details:</p>
					<ul>
						{{range .PassengerDetails}}
							<li>{{.Name}}</li>
							<l1>{{.SeatNumber}}</li>
							{{end}}
							</ul>
							</body>
							</html>
							`
	templateSendReceipt = `
		<html>
			<head>
				<title>Receipt</title>
				</head>
				<body>
					<h1>Receipt</h1>
					<p>Invoice Number: {{.InvoiceNumber}}</p>
					<p>Booking Code: {{.BookingCode}}</p>
					<p>Payment Method: {{.PaymentMethod}}</p>
					<p>Payment Amount: {{.PaymentAmount}}</p>
					<p>Payment Date: {{.PaymentDate}}</p>
					</body>
					</html>
					`
	templateSendTicket = `
		<html>
			<head>
				<title>Ticket</title>
				</head>
				<body>
					<h1>Ticket</h1>
					<p>Ticket Code: {{.TicketCode}}</p>
					<p>Flight Number: {{.FlightNumber}}</p>
					<p>Flight Departure Time: {{.FlightDepartureTime}}</p>
					<p>Flight Arrival Time: {{.FlightArrivalTime}}</p>
					<p>Flight Departure Airport: {{.FlightDepartureAirport}}</p>
					<p>Flight Arrival Airport: {{.FlightArrivalAirport}}</p>
					<p>Passenger Details:</p>
					<ul>
						{{range .PassengerDetails}}
							<li>{{.Name}}</li>
							<l1>{{.SeatNumber}}</li>
							{{end}}
							</ul>
							<p>Boarding Time: {{.BoardingTime}}</p>
							</body>
							</html>
							`
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
