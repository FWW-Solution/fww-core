package dto_notification

import (
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/data/dto_payment"
	"fww-core/internal/entity"
	"time"
)

type Request struct {
	CodeBooking string `json:"code_booking"`
	Route       string `json:"route"`
}

type SendEmailRequest struct {
	EmailAddress string   `json:"email_address" validate:"email"`
	To           string   `json:"to" validate:"required,email"`
	Cc           string   `json:"cc" validate:"email"`
	Bcc          string   `json:"bcc" validate:"email"`
	Subject      string   `json:"subject" validate:"required"`
	Body         string   `json:"body" validate:"required"`
	Attachments  []string `json:"attachments"`
}

type PaymentInvoiceAggregator struct {
	Payment        entity.Payment         `json:"payment"`
	BookingDetails []entity.BookingDetail `json:"booking_detail"`
	PaymentMethods []entity.PaymentMethod `json:"payment_methods"`
	Passengers     []entity.Passenger     `json:"passengers"`
	Booking        entity.Booking         `json:"booking"`
	User           entity.User            `json:"user"`
}

type PaymentReceiptAggregator struct {
	Payment        entity.Payment         `json:"payment"`
	BookingDetails []entity.BookingDetail `json:"booking_detail"`
	Booking        entity.Booking         `json:"booking"`
	User           entity.User            `json:"user"`
}

type TicketRedeemAgregator struct {
	Ticket         entity.Ticket          `json:"ticket"`
	Booking        entity.Booking         `json:"booking"`
	BookingDetails []entity.BookingDetail `json:"booking_detail"`
	Flight         entity.Flight          `json:"flight"`
	User           entity.User            `json:"user"`
}

type ModelTicketRedeemed struct {
	TicketCode             string                           `json:"ticket_code"`
	FlightNumber           string                           `json:"flight_number"`
	FlightDepartureTime    string                           `json:"flight_departure_time"`
	FlightArrivalTime      string                           `json:"flight_arrival_time"`
	FlightDepartureAirport string                           `json:"flight_departure_airport"`
	FlightArrivalAirport   string                           `json:"flight_arrival_airport"`
	PassengerDetails       []dto_booking.BookResponseDetail `json:"passenger_details"`
	BoardingTime           string                           `json:"boarding_time"`
}

type ModelInvoice struct {
	InvoiceNumber     string                           `json:"invoice_number"`
	BookingCode       string                           `json:"booking_code"`
	PaymentMethodList []dto_payment.MethodResponse     `json:"payment_methods"`
	PaymentAmount     float64                          `json:"payment_ammount"`
	PassengerDetails  []dto_booking.BookResponseDetail `json:"passenger_details"`
}

type ModelPaymentReceipt struct {
	InvoiceNumber string    `json:"invoice_number"`
	PaymentMethod string    `json:"payment_method"`
	PaymentAmount float64   `json:"payment_ammount"`
	PaymentDate   time.Time `json:"payment_date"`
	BookingCode   string    `json:"booking_code"`
}
