package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/data/dto_notification"
	"fww-core/internal/data/dto_payment"
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

		// Transform data to model
		var paymentMethodResponse []dto_payment.MethodResponse
		for _, paymentMethod := range result.PaymentMethods {
			spec := dto_payment.MethodResponse{
				ID:       paymentMethod.ID,
				Name:     paymentMethod.Name,
				IsActive: paymentMethod.IsActive,
			}
			paymentMethodResponse = append(paymentMethodResponse, spec)
		}

		var passengerDetails []dto_booking.BookResponseDetail
		for i, bookingDetail := range result.BookingDetails {
			spec := dto_booking.BookResponseDetail{
				SeatNumber:    bookingDetail.SeatNumber,
				Class:         bookingDetail.Class,
				Baggage:       bookingDetail.BaggageCapacity,
				PassangerName: result.Passengers[i].FullName,
			}
			passengerDetails = append(passengerDetails, spec)
		}
		specModel := dto_notification.ModelInvoice{
			InvoiceNumber:     result.Payment.InvoiceNumber,
			BookingCode:       result.Booking.CodeBooking,
			PaymentAmount:     result.Payment.TotalPayment,
			PassengerDetails:  passengerDetails,
			PaymentMethodList: paymentMethodResponse,
		}

		jsonData, err := json.Marshal(specModel)
		if err != nil {
			return err
		}

		fmt.Println(string(jsonData))

		return nil

		// // TODO: Populate data to template

		// specNotification := dto_notification.SendEmailRequest{
		// 	To:      result.User.Email,
		// 	Subject: "[FWW] Invoice",
		// 	Body:    templateSendInvoice,
		// }

		// err = u.adapter.SendNotification(&specNotification)
		// if err != nil {
		// 	return err
		// }

	case "send_receipt":
		result, err := u.repository.PaymentReceiptReportByBookingCode(data.CodeBooking)
		if err != nil {
			return err
		}

		// Transform data to model
		specModel := dto_notification.ModelPaymentReceipt{
			InvoiceNumber: result.Payment.InvoiceNumber,
			BookingCode:   result.Booking.CodeBooking,
			PaymentAmount: result.Payment.TotalPayment,
			PaymentDate:   result.Payment.PaymentDate,
			PaymentMethod: result.Payment.PaymentMethod,
		}

		jsonData, err := json.Marshal(specModel)
		if err != nil {
			return err
		}

		fmt.Println(string(jsonData))

		return nil

		// TODO: Populate data to template
		// spec := dto_notification.ModelPaymentReceipt{}

		// specNotification := dto_notification.SendEmailRequest{
		// 	To:      result.User.Email,
		// 	Subject: "[FWW] Receipt",
		// 	Body:    templateSendReceipt,
		// }

		// // err = u.adapter.SendNotification(&specNotification)
		// // if err != nil {
		// // 	return err
		// // }

	case "send_ticket":
		result, err := u.repository.TicketRedeemedReportByBookingCode(data.CodeBooking)
		if err != nil {
			return err
		}

		fmt.Println(result)

		// Transform data to model
		var passengerDetails []dto_booking.BookResponseDetail
		for i, bookingDetail := range result.BookingDetails {
			spec := dto_booking.BookResponseDetail{
				SeatNumber:    bookingDetail.SeatNumber,
				Class:         bookingDetail.Class,
				Baggage:       bookingDetail.BaggageCapacity,
				PassangerName: result.Passengers[i].FullName,
			}
			passengerDetails = append(passengerDetails, spec)
		}

		specModel := dto_notification.ModelTicketRedeemed{
			TicketCode:             result.Ticket.CodeTicket,
			FlightNumber:           result.Flight.CodeFlight,
			FlightDepartureTime:    result.Flight.DepartureTime.Format("2006-01-02 15:04:05"),
			FlightArrivalTime:      result.Flight.ArrivalTime.Format("2006-01-02 15:04:05"),
			FlightDepartureAirport: result.Flight.DepartureAirportName,
			FlightArrivalAirport:   result.Flight.ArrivalAirportName,
			PassengerDetails:       passengerDetails,
			BoardingTime:           result.Ticket.BoardingTime.Time.Format("2006-01-02 15:04:05"),
		}

		jsonData, err := json.Marshal(specModel)
		if err != nil {
			return err
		}

		fmt.Println(string(jsonData))

		return nil

		// // TODO: Populate data to template
		// // spec := dto_notification.ModelTicketRedeemed{}

		// specNotification := dto_notification.SendEmailRequest{
		// 	To:      result.User.Email,
		// 	Subject: "[FWW] Ticket",
		// 	Body:    templateSendTicket,
		// }

		// err = u.adapter.SendNotification(&specNotification)
		// if err != nil {
		// 	return err
		// }

	default:
		return errors.New("route not found")

	}
	return nil
}
