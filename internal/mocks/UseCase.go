// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	dto_airport "fww-core/internal/data/dto_airport"
	dto_booking "fww-core/internal/data/dto_booking"

	dto_flight "fww-core/internal/data/dto_flight"

	dto_passanger "fww-core/internal/data/dto_passanger"

	dto_payment "fww-core/internal/data/dto_payment"

	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// DetailPassanger provides a mock function with given fields: id
func (_m *UseCase) DetailPassanger(id int64) (dto_passanger.ResponseDetail, error) {
	ret := _m.Called(id)

	var r0 dto_passanger.ResponseDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (dto_passanger.ResponseDetail, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) dto_passanger.ResponseDetail); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(dto_passanger.ResponseDetail)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAirport provides a mock function with given fields: city, province, iata
func (_m *UseCase) GetAirport(city string, province string, iata string) ([]dto_airport.ResponseAirport, error) {
	ret := _m.Called(city, province, iata)

	var r0 []dto_airport.ResponseAirport
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) ([]dto_airport.ResponseAirport, error)); ok {
		return rf(city, province, iata)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) []dto_airport.ResponseAirport); ok {
		r0 = rf(city, province, iata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto_airport.ResponseAirport)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(city, province, iata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetailBooking provides a mock function with given fields: codeBooking
func (_m *UseCase) GetDetailBooking(codeBooking string) (dto_booking.BookResponse, error) {
	ret := _m.Called(codeBooking)

	var r0 dto_booking.BookResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (dto_booking.BookResponse, error)); ok {
		return rf(codeBooking)
	}
	if rf, ok := ret.Get(0).(func(string) dto_booking.BookResponse); ok {
		r0 = rf(codeBooking)
	} else {
		r0 = ret.Get(0).(dto_booking.BookResponse)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(codeBooking)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetailFlightByID provides a mock function with given fields: id
func (_m *UseCase) GetDetailFlightByID(id int64) (dto_flight.ResponseFlightDetail, error) {
	ret := _m.Called(id)

	var r0 dto_flight.ResponseFlightDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (dto_flight.ResponseFlightDetail, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) dto_flight.ResponseFlightDetail); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(dto_flight.ResponseFlightDetail)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFlights provides a mock function with given fields: departureTime, ArrivalTime, limit, offset
func (_m *UseCase) GetFlights(departureTime string, ArrivalTime string, limit int, offset int) ([]dto_flight.ResponseFlight, error) {
	ret := _m.Called(departureTime, ArrivalTime, limit, offset)

	var r0 []dto_flight.ResponseFlight
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) ([]dto_flight.ResponseFlight, error)); ok {
		return rf(departureTime, ArrivalTime, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) []dto_flight.ResponseFlight); ok {
		r0 = rf(departureTime, ArrivalTime, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto_flight.ResponseFlight)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(departureTime, ArrivalTime, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentStatus provides a mock function with given fields: codePayment
func (_m *UseCase) GetPaymentStatus(codePayment string) (dto_payment.StatusResponse, error) {
	ret := _m.Called(codePayment)

	var r0 dto_payment.StatusResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (dto_payment.StatusResponse, error)); ok {
		return rf(codePayment)
	}
	if rf, ok := ret.Get(0).(func(string) dto_payment.StatusResponse); ok {
		r0 = rf(codePayment)
	} else {
		r0 = ret.Get(0).(dto_payment.StatusResponse)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(codePayment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterPassanger provides a mock function with given fields: data
func (_m *UseCase) RegisterPassanger(data *dto_passanger.RequestRegister) (dto_passanger.ResponseRegistered, error) {
	ret := _m.Called(data)

	var r0 dto_passanger.ResponseRegistered
	var r1 error
	if rf, ok := ret.Get(0).(func(*dto_passanger.RequestRegister) (dto_passanger.ResponseRegistered, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*dto_passanger.RequestRegister) dto_passanger.ResponseRegistered); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(dto_passanger.ResponseRegistered)
	}

	if rf, ok := ret.Get(1).(func(*dto_passanger.RequestRegister) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestBooking provides a mock function with given fields: data, bookingIDCode
func (_m *UseCase) RequestBooking(data *dto_booking.Request, bookingIDCode string) error {
	ret := _m.Called(data, bookingIDCode)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_booking.Request, string) error); ok {
		r0 = rf(data, bookingIDCode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestPayment provides a mock function with given fields: req, paymentCodeID
func (_m *UseCase) RequestPayment(req *dto_payment.Request, paymentCodeID string) error {
	ret := _m.Called(req, paymentCodeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_payment.Request, string) error); ok {
		r0 = rf(req, paymentCodeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePassanger provides a mock function with given fields: data
func (_m *UseCase) UpdatePassanger(data *dto_passanger.RequestUpdate) (dto_passanger.ResponseUpdate, error) {
	ret := _m.Called(data)

	var r0 dto_passanger.ResponseUpdate
	var r1 error
	if rf, ok := ret.Get(0).(func(*dto_passanger.RequestUpdate) (dto_passanger.ResponseUpdate, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*dto_passanger.RequestUpdate) dto_passanger.ResponseUpdate); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(dto_passanger.ResponseUpdate)
	}

	if rf, ok := ret.Get(1).(func(*dto_passanger.RequestUpdate) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCase(t mockConstructorTestingTNewUseCase) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
