// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	dto_airport "fww-core/internal/data/dto_airport"
	dto_booking "fww-core/internal/data/dto_booking"

	dto_flight "fww-core/internal/data/dto_flight"

	dto_notification "fww-core/internal/data/dto_notification"

	dto_passanger "fww-core/internal/data/dto_passanger"

	dto_payment "fww-core/internal/data/dto_payment"

	dto_ticket "fww-core/internal/data/dto_ticket"

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

// DoPayment provides a mock function with given fields: codePayment
func (_m *UseCase) DoPayment(codePayment string) error {
	ret := _m.Called(codePayment)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(codePayment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateInvoice provides a mock function with given fields: caseID, codeBooking
func (_m *UseCase) GenerateInvoice(caseID int64, codeBooking string) error {
	ret := _m.Called(caseID, codeBooking)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, string) error); ok {
		r0 = rf(caseID, codeBooking)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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

// GetPaymentMethod provides a mock function with given fields:
func (_m *UseCase) GetPaymentMethod() ([]dto_payment.MethodResponse, error) {
	ret := _m.Called()

	var r0 []dto_payment.MethodResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]dto_payment.MethodResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []dto_payment.MethodResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto_payment.MethodResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
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

// InquiryNotification provides a mock function with given fields: data
func (_m *UseCase) InquiryNotification(data *dto_notification.Request) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_notification.Request) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RedeemTicket provides a mock function with given fields: codeBooking
func (_m *UseCase) RedeemTicket(codeBooking string) (dto_ticket.Response, error) {
	ret := _m.Called(codeBooking)

	var r0 dto_ticket.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (dto_ticket.Response, error)); ok {
		return rf(codeBooking)
	}
	if rf, ok := ret.Get(0).(func(string) dto_ticket.Response); ok {
		r0 = rf(codeBooking)
	} else {
		r0 = ret.Get(0).(dto_ticket.Response)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(codeBooking)
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

// RequestPayment provides a mock function with given fields: req
func (_m *UseCase) RequestPayment(req *dto_payment.Request) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_payment.Request) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDetailBooking provides a mock function with given fields: data
func (_m *UseCase) UpdateDetailBooking(data *dto_booking.BookDetailRequest) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_booking.BookDetailRequest) error); ok {
		r0 = rf(data)
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

// UpdatePassangerByIDNumber provides a mock function with given fields: data
func (_m *UseCase) UpdatePassangerByIDNumber(data *dto_passanger.RequestUpdateBPM) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_passanger.RequestUpdateBPM) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePayment provides a mock function with given fields: req
func (_m *UseCase) UpdatePayment(req *dto_payment.RequestUpdatePayment) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_payment.RequestUpdatePayment) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTicket provides a mock function with given fields: req
func (_m *UseCase) UpdateTicket(req *dto_ticket.RequestUpdateTicket) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*dto_ticket.RequestUpdateTicket) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
