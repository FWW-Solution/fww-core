package usecase

import (
	"errors"
	"fww-core/internal/data/dto_flight"
	"fww-core/internal/tools"
)

// GetDetailFlightByID implements UseCase.
func (u *useCase) GetDetailFlightByID(id int64) (dto_flight.ResponseFlightDetail, error) {
	resultFlight, err := u.repository.FindFlightByID(id)
	if err != nil {
		return dto_flight.ResponseFlightDetail{}, tools.ErrorBuilder(err)
	}

	if resultFlight.ID == 0 {
		return dto_flight.ResponseFlightDetail{}, errors.New("flight not found")
	}

	resultFlightPrice, err := u.repository.FindFlightPriceByID(id)
	if err != nil {
		return dto_flight.ResponseFlightDetail{}, tools.ErrorBuilder(err)
	}

	if resultFlightPrice.ID == 0 {
		return dto_flight.ResponseFlightDetail{}, errors.New("flight price not found")
	}

	resultFlightReservation, err := u.repository.FindFlightReservationByID(id)
	if err != nil {
		return dto_flight.ResponseFlightDetail{}, tools.ErrorBuilder(err)
	}

	if resultFlightReservation.ID == 0 {
		return dto_flight.ResponseFlightDetail{}, errors.New("flight reservation not found")
	}

	response := dto_flight.ResponseFlightDetail{
		ArrivalAirportName:  resultFlight.ArrivalAirportName,
		ArrivalTime:         resultFlight.ArrivalTime.Format("2006-01-02 15:04:05"),
		CodeFlight:          resultFlight.CodeFlight,
		DepartureTime:       resultFlight.DepartureTime.Format("2006-01-02 15:04:05"),
		DepatureAirportName: resultFlight.DepartureAirportName,
		FlightPrice:         resultFlightPrice.Price,
		ReminingSeat:        resultFlightReservation.ReminingSeat,
		Status:              resultFlight.Status,
	}

	return response, nil
}

// GetFlights implements UseCase.
func (u *useCase) GetFlights(departureTime string, arrivalTime string, limit int, offset int) ([]dto_flight.ResponseFlight, error) {
	result, err := u.repository.FindFlights(departureTime, arrivalTime, limit, offset)
	if err != nil {
		return nil, tools.ErrorBuilder(err)
	}

	var response []dto_flight.ResponseFlight

	for _, v := range result {
		response = append(response, dto_flight.ResponseFlight{
			ArrivalAirportName:  v.ArrivalAirportName,
			ArrivalTime:         v.ArrivalTime.Format("2006-01-02 15:04:05"),
			CodeFlight:          v.CodeFlight,
			DepartureTime:       v.DepartureTime.Format("2006-01-02 15:04:05"),
			DepatureAirportName: v.DepartureAirportName,
			Status:              v.Status,
		})
	}

	return response, nil
}
