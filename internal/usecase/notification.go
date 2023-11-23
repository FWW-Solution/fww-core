package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/data/dto_notification"
	"fww-core/internal/data/dto_payment"
	"text/template"
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
					<p>Payment Amount: {{.PaymentAmount}}</p>
					<p>Passenger Details:</p>
					<ul>
						{{range .PassengerDetails}}
							<li>{{.PassangerName}}</li>
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
							<li>{{.PassangerName}}</li>
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

		// // TODO: Populate data to template

		templateSendInvoice, err := u.populateTemplateInvoice(&specModel, templateSendInvoice)
		if err != nil {
			return err
		}

		specNotification := dto_notification.SendEmailRequest{
			To:      result.User.Email,
			Subject: "[FWW] Invoice",
			Body:    templateSendInvoice,
		}

		jsonData, err := json.Marshal(specNotification)
		if err != nil {
			return err
		}

		fmt.Println(string(jsonData))
		return nil
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

		// TODO: Populate data to template
		// spec := dto_notification.ModelPaymentReceipt{}

		templateSendReceipt, err := u.populateTemplateReceipt(&specModel, templateSendReceipt)
		if err != nil {
			return err
		}

		specNotification := dto_notification.SendEmailRequest{
			To:      result.User.Email,
			Subject: "[FWW] Receipt",
			Body:    templateSendReceipt,
		}

		jsonData, err := json.Marshal(specNotification)
		if err != nil {
			return err
		}

		fmt.Println(string(jsonData))
		return nil

		// // err = u.adapter.SendNotification(&specNotification)
		// // if err != nil {
		// // 	return err
		// // }

	case "send_ticket":
		result, err := u.repository.TicketRedeemedReportByBookingCode(data.CodeBooking)
		if err != nil {
			return err
		}

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

		// // TODO: Populate data to template
		// // spec := dto_notification.ModelTicketRedeemed{}
		templateSendReceipt, err := u.populateTemplateTicket(&specModel, templateSendTicket)
		if err != nil {
			return err
		}

		specNotification := dto_notification.SendEmailRequest{
			To:      result.User.Email,
			Subject: "[FWW] Receipt",
			Body:    templateSendReceipt,
		}

		jsonData, err := json.Marshal(specNotification)
		if err != nil {
			return err
		}

		fmt.Println(string(jsonData))

		// err = u.adapter.SendNotification(&specNotification)
		// if err != nil {
		// 	return err
		// }

	default:
		return errors.New("route not found")

	}
	return nil
}

func (u *useCase) populateTemplateInvoice(data *dto_notification.ModelInvoice, templateHtml string) (string, error) {
	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(templateHtml))
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (u *useCase) populateTemplateReceipt(data *dto_notification.ModelPaymentReceipt, templateHtml string) (string, error) {
	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(templateHtml))
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (u *useCase) populateTemplateTicket(data *dto_notification.ModelTicketRedeemed, templateHtml string) (string, error) {
	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(templateHtml))
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
