package usecase

import (
	"fww-core/internal/adapter"
	"fww-core/internal/data/dto_airport"
	"fww-core/internal/data/dto_booking"
	"fww-core/internal/data/dto_flight"
	"fww-core/internal/data/dto_passanger"
	"fww-core/internal/data/dto_payment"
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
	// Airport
	GetAirport(city string, province string, iata string) ([]dto_airport.ResponseAirport, error)
	// Flight
	GetFlights(departureTime string, ArrivalTime string, limit int, offset int) ([]dto_flight.ResponseFlight, error)
	GetDetailFlightByID(id int64) (dto_flight.ResponseFlightDetail, error)
	// Booking
	RequestBooking(data *dto_booking.Request, bookingIDCode string) error
	GetDetailBooking(codeBooking string) (dto_booking.BookResponse, error)
	// Payment
	RequestPayment(req *dto_payment.Request, paymentCodeID string) error
	GetPaymentStatus(codePayment string) (dto_payment.StatusResponse, error)
	GetPaymentMethod() ([]dto_payment.MethodResponse, error)
}

func New(repository repository.Repository, adapter adapter.Adapter, redis *redis.Client) UseCase {
	return &useCase{
		repository: repository,
		adapter:    adapter,
		redis:      redis,
	}
}
