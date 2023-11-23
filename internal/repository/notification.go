package repository

import (
	"fmt"
	"fww-core/internal/data/dto_notification"
	"fww-core/internal/entity"
	"log"
)

// PaymentInvoiceReportByBookingID implements Repository.
func (r *repository) PaymentInvoiceReportByBookingCode(bookingCode string) (dto_notification.PaymentInvoiceAggregator, error) {
	query := fmt.Sprintf(`SELECT b.id, p.invoice_number, b.code_booking, p.total_payment, u.email FROM bookings as b INNER JOIN users as u ON b.user_id = u.id INNER JOIN payments as p ON p.booking_id = b.id WHERE b.code_booking = '%s'`, bookingCode)
	rows, err := r.db.Queryx(query)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return dto_notification.PaymentInvoiceAggregator{}, nil
	}

	if err != nil {
		return dto_notification.PaymentInvoiceAggregator{}, err
	}

	defer rows.Close()

	var entityPayment entity.Payment
	var entityBookingDetails []entity.BookingDetail
	var entityPassengers []entity.Passenger
	var entityPaymentMethods []entity.PaymentMethod
	var entityBooking entity.Booking
	var entityUser entity.User
	for rows.Next() {
		err = rows.Scan(&entityBooking.ID, &entityPayment.InvoiceNumber, &entityBooking.CodeBooking, &entityPayment.TotalPayment, &entityUser.Email)
		if err != nil {
			return dto_notification.PaymentInvoiceAggregator{}, err
		}
	}

	// Find Booking Detail
	queryBookingDetail := fmt.Sprintf(`SELECT p.full_name, bd.seat_number, bd.class, bd.baggage_capacity FROM booking_details as bd INNER JOIN passengers as p ON p.id = bd.passenger_id WHERE booking_id = %d`, entityBooking.ID)

	rowsBookingDetail, err := r.db.Queryx(queryBookingDetail)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return dto_notification.PaymentInvoiceAggregator{}, nil
	}

	if err != nil {
		return dto_notification.PaymentInvoiceAggregator{}, err
	}

	defer rowsBookingDetail.Close()

	for rowsBookingDetail.Next() {
		var entityBookingDetail entity.BookingDetail
		var entityPassenger entity.Passenger
		err = rowsBookingDetail.Scan(&entityPassenger.FullName, &entityBookingDetail.SeatNumber, &entityBookingDetail.Class, &entityBookingDetail.BaggageCapacity)
		if err != nil {
			return dto_notification.PaymentInvoiceAggregator{}, err
		}
		entityBookingDetails = append(entityBookingDetails, entityBookingDetail)
		entityPassengers = append(entityPassengers, entityPassenger)
	}

	// Find Payment Method
	queryPaymentMethod := `SELECT id, name, is_active FROM payment_methods`

	result, err := r.db.Queryx(queryPaymentMethod)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return dto_notification.PaymentInvoiceAggregator{}, nil
	}

	if err != nil {
		return dto_notification.PaymentInvoiceAggregator{}, err
	}

	for result.Next() {
		var row entity.PaymentMethod
		err := result.StructScan(&row)
		if err != nil {
			return dto_notification.PaymentInvoiceAggregator{}, err
		}
		entityPaymentMethods = append(entityPaymentMethods, row)
	}

	aggregator := dto_notification.PaymentInvoiceAggregator{
		Payment:        entityPayment,
		BookingDetails: entityBookingDetails,
		PaymentMethods: entityPaymentMethods,
		Passengers:     entityPassengers,
		Booking:        entityBooking,
		User:           entityUser,
	}

	return aggregator, nil
}

// PaymentReceiptReportByBookingID implements Repository.
func (r *repository) PaymentReceiptReportByBookingCode(bookingCode string) (dto_notification.PaymentReceiptAggregator, error) {
	query := fmt.Sprintf(`SELECT p.invoice_number, b.code_booking, p.total_payment, p.payment_method, p.payment_date, u.email FROM bookings as b INNER JOIN users as u ON u.id = b.user_id INNER JOIN payments as p ON p.booking_id = b.id WHERE b.code_booking = '%s'`, bookingCode)
	rows, err := r.db.Queryx(query)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return dto_notification.PaymentReceiptAggregator{}, nil
	}

	if err != nil {
		return dto_notification.PaymentReceiptAggregator{}, err
	}

	defer rows.Close()

	var entityPayment entity.Payment
	var entityBooking entity.Booking
	var entityUser entity.User
	for rows.Next() {
		err = rows.Scan(&entityPayment.InvoiceNumber, &entityBooking.CodeBooking, &entityPayment.TotalPayment, &entityPayment.PaymentMethod, &entityPayment.PaymentDate, &entityUser.Email)
		if err != nil {
			return dto_notification.PaymentReceiptAggregator{}, err
		}
	}

	aggregator := dto_notification.PaymentReceiptAggregator{
		Payment: entityPayment,
		Booking: entityBooking,
		User:    entityUser,
	}

	return aggregator, nil
}

// TicketRedeemedReportByBookingID implements Repository.
func (r *repository) TicketRedeemedReportByBookingCode(bookingCode string) (dto_notification.TicketRedeemAgregator, error) {
	query := fmt.Sprintf(`SELECT b.id, f.code_flight, f.departure_time, f.arrival_time, f.departure_airport_name, f.arrival_airport_name,  t.code_ticket, t.boarding_time, u.email FROM bookings as b INNER JOIN users as u ON b.user_id = u.id INNER JOIN flights as f ON f.id = b.flight_id INNER JOIN tickets as t ON t.booking_id = b.id WHERE b.code_booking = '%s'`, bookingCode)

	rows, err := r.db.Queryx(query)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return dto_notification.TicketRedeemAgregator{}, nil
	}

	if err != nil {
		return dto_notification.TicketRedeemAgregator{}, err
	}

	defer rows.Close()

	var entityTicket entity.Ticket
	var entityBooking entity.Booking
	var entityFlight entity.Flight
	var entityUser entity.User
	var entityBookingDetails []entity.BookingDetail
	var entityPassengers []entity.Passenger
	for rows.Next() {
		err = rows.Scan(&entityBooking.ID, &entityFlight.CodeFlight, &entityFlight.DepartureTime, &entityFlight.ArrivalTime, &entityFlight.DepartureAirportName, &entityFlight.ArrivalAirportName, &entityTicket.CodeTicket, &entityTicket.BoardingTime, &entityUser.Email)
		if err != nil {
			log.Println(err)
			return dto_notification.TicketRedeemAgregator{}, err
		}
	}

	// Find Booking Detail
	queryBookingDetail := fmt.Sprintf(`SELECT p.full_name, bd.seat_number, bd.class, bd.baggage_capacity FROM booking_details as bd INNER JOIN passengers as p ON p.id = bd.passenger_id WHERE booking_id = %d`, entityBooking.ID)

	rowsBookingDetail, err := r.db.Queryx(queryBookingDetail)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return dto_notification.TicketRedeemAgregator{}, nil
	}

	if err != nil {
		return dto_notification.TicketRedeemAgregator{}, err
	}

	defer rowsBookingDetail.Close()

	for rowsBookingDetail.Next() {
		var entityBookingDetail entity.BookingDetail
		var entityPassenger entity.Passenger
		err = rowsBookingDetail.Scan(&entityPassenger.FullName, &entityBookingDetail.SeatNumber, &entityBookingDetail.Class, &entityBookingDetail.BaggageCapacity)
		if err != nil {
			return dto_notification.TicketRedeemAgregator{}, err
		}
		entityBookingDetails = append(entityBookingDetails, entityBookingDetail)
		entityPassengers = append(entityPassengers, entityPassenger)
	}

	aggregator := dto_notification.TicketRedeemAgregator{
		Ticket:         entityTicket,
		Booking:        entityBooking,
		BookingDetails: entityBookingDetails,
		Flight:         entityFlight,
		Passengers:     entityPassengers,
		User:           entityUser,
	}

	return aggregator, nil

}
