package usecase

import (
	"database/sql"
	"errors"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/data/dto_ticket"
	"fww-core/internal/entity"
	"time"

	"github.com/google/uuid"
)

// RedeemTicket implements UseCase.
func (u *useCase) RedeemTicket(codeBooking string) (dto_ticket.Response, error) {

	// Find Booking by code
	booking, err := u.repository.FindBookingByBookingIDCode(codeBooking)
	if err != nil {
		return dto_ticket.Response{}, err
	}

	if booking.ID == 0 {
		return dto_ticket.Response{}, errors.New("booking not found")
	}

	// Validate Booking Date
	if booking.BookingExpiredAt.Before(time.Now()) {
		return dto_ticket.Response{}, errors.New("booking expired")
	}

	// Update Booking Status
	booking.BookingStatus = "redeemed"
	_, err = u.repository.UpdateBooking(&booking)
	if err != nil {
		return dto_ticket.Response{}, err
	}

	// Upsert Ticket
	generatedUUID := uuid.New()

	boardingTime := booking.BookingExpiredAt.Add((24 * time.Hour) - (time.Minute * 30))

	entityTicket := entity.Ticket{
		ID:                 0,
		CodeTicket:         generatedUUID.String(),
		IsBoardingPass:     true,
		IsEligibleToFlight: true,
		BookingID:          booking.ID,
		BoardingTime: sql.NullTime{
			Time:  boardingTime,
			Valid: true,
		},
	}

	_, err = u.repository.UpsertTicket(&entityTicket)
	if err != nil {
		return dto_ticket.Response{}, err
	}

	bookingDetails, err := u.repository.FindBookingDetailByBookingID(booking.ID)
	if err != nil {
		return dto_ticket.Response{}, err
	}

	passengersInfo := make([]dto_ticket.PassengerInfoData, 0)
	for _, bookingDetail := range bookingDetails {
		passenger, err := u.repository.FindDetailPassanger(bookingDetail.PassengerID)
		if err != nil {
			return dto_ticket.Response{}, err
		}
		passengersInfo = append(passengersInfo, dto_ticket.PassengerInfoData{
			BookingDetailID: bookingDetail.ID,
			IDNumber:        passenger.IDNumber,
			VaccineStatus:   passenger.CovidVaccineStatus,
		})
	}

	specRedeem := dto_ticket.RequestRedeemTicketToBPM{
		CaseID:         booking.CaseID,
		CodeTicket:     entityTicket.CodeTicket,
		PassengersInfo: passengersInfo,
	}

	err = u.adapter.RedeemTicket(&specRedeem)
	if err != nil {
		return dto_ticket.Response{}, err
	}

	response := dto_ticket.Response{
		CodeTicket:  entityTicket.CodeTicket,
		BordingTime: boardingTime.Format("2006-01-02 15:04:05"),
	}

	return response, nil

}

// UpdateTicket implements UseCase.
func (u *useCase) UpdateTicket(req *dto_ticket.RequestUpdateTicket) error {
	// Find Ticket by code
	ticket, err := u.repository.FindTicketByCodeTicket(req.CodeTicket)
	if err != nil {
		return err
	}

	if ticket.ID == 0 {
		return errors.New("ticket not found")
	}

	// Update Detail Passenger
	specUpdate := dto_booking.BookDetailRequest{
		BookingDetailID:    req.BookingDetailID,
		IsEligibleToFlight: req.IsEligibleToFlight,
	}
	err = u.UpdateDetailBooking(&specUpdate)
	if err != nil {
		return err
	}

	return nil
}
