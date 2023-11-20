package usecase

import (
	"fww-core/internal/adapter"
	"fww-core/internal/data/dto_airport"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/data/dto_flight"
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/data/dto_payment"
	"fww-core/internal/data/dto_ticket"
	"fww-core/internal/repository"

	"github.com/redis/go-redis/v9"
)

type useCase struct {
	repository repository.Repository
	adapter    adapter.Adapter
	redis      *redis.Client
}

type UseCase interface {
	RegisterPassanger(data *dto_passanger.RequestRegister) (dto_passanger.ResponseRegistered, error)
	DetailPassanger(id int64) (dto_passanger.ResponseDetail, error)
	UpdatePassanger(data *dto_passanger.RequestUpdate) (dto_passanger.ResponseUpdate, error)
	UpdatePassangerByIDNumber(data *dto_passanger.RequestUpdateBPM) error
	// Airport
	GetAirport(city string, province string, iata string) ([]dto_airport.ResponseAirport, error)
	// Flight
	GetFlights(departureTime string, ArrivalTime string, limit int, offset int) ([]dto_flight.ResponseFlight, error)
	GetDetailFlightByID(id int64) (dto_flight.ResponseFlightDetail, error)
	// Booking
	RequestBooking(data *dto_booking.Request, bookingIDCode string) error
	GetDetailBooking(codeBooking string) (dto_booking.BookResponse, error)
	// Payment
	RequestPayment(req *dto_payment.Request) error
	GetPaymentStatus(codePayment string) (dto_payment.StatusResponse, error)
	GetPaymentMethod() ([]dto_payment.MethodResponse, error)
	DoPayment(codePayment string) error
	GenerateInvoice(caseID int64, codeBooking string) error
	UpdatePayment(req *dto_payment.RequestUpdatePayment) error
	// Ticket
	RedeemTicket(codeBooking string) (dto_ticket.Response, error)
	UpdateTicket(req *dto_ticket.RequestUpdateTicket) error
}

func New(repository repository.Repository, adapter adapter.Adapter, redis *redis.Client) UseCase {
	return &useCase{
		repository: repository,
		adapter:    adapter,
		redis:      redis,
	}
}
