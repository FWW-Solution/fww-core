package entity

import (
	"database/sql"
	"time"
)

// User represents a user entity in the database.
type User struct {
	ID        int64      `db:"id"`
	FullName  string     `db:"full_name"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	Password  string     `db:"password"` // encrypted
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

// Passenger represents a passenger entity.
type Passenger struct {
	ID                 int64      `db:"id"`
	FullName           string     `db:"full_name"`
	Gender             string     `db:"gender"` // enum (male, female)
	DateOfBirth        time.Time  `db:"date_of_birth"`
	IDNumber           string     `db:"id_number"`
	IDType             string     `db:"id_type"`              // enum (passport, ktp, driver_license)
	CovidVaccineStatus string     `db:"covid_vaccine_status"` // enum (Vaccine I, Vaccine II, Vaccine III, Not Vaccinated)
	IsIDVerified       bool       `db:"is_id_verified"`
	CaseID             int64      `db:"case_id"`
	CreatedAt          time.Time  `db:"created_at"`
	UpdatedAt          time.Time  `db:"updated_at"`
	DeletedAt          *time.Time `db:"deleted_at"`
}

// PlaneInformation represents information about a plane.
type PlaneInformation struct {
	ID                   int64      `db:"id"`
	CodePlane            string     `db:"code_plane"`
	TotalBaggageCapacity int        `db:"total_baggage_capacity"`
	Type                 string     `db:"type"` // enum (airbus, boeing)
	Variant              string     `db:"variant"`
	CreatedAt            time.Time  `db:"created_at"`
	UpdatedAt            time.Time  `db:"updated_at"`
	DeletedAt            *time.Time `db:"deleted_at"`
}

type PlaneInformationDetail struct {
	ID                int64      `json:"id" db:"id"`
	Class             string     `json:"class" db:"class"` // enum (economy, business, first)
	TotalSeatCapacity int        `json:"total_seat_capacity" db:"total_seat_capacity"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
	PlaneID           int        `json:"plane_id" db:"plane_id"`
	DeletedAt         *time.Time `json:"deleted_at" db:"deleted_at"`
}

// Flight represents a flight entity in the database.
type Flight struct {
	ID                   int64        `db:"id"`
	CodeFlight           string       `db:"code_flight"`
	DepartureTime        time.Time    `db:"departure_time"`
	ArrivalTime          time.Time    `db:"arrival_time"`
	DepartureAirportName string       `db:"departure_airport_name"`
	ArrivalAirportName   string       `db:"arrival_airport_name"`
	DepartureAirportID   int64        `db:"departure_airport_id"`
	ArrivalAirportID     int64        `db:"arrival_airport_id"`
	Status               string       `db:"status"` // enum (on_time, delayed, canceled)
	CreatedAt            time.Time    `db:"created_at"`
	UpdatedAt            sql.NullTime `db:"updated_at"`
	DeletedAt            sql.NullTime `db:"deleted_at"`
	PlaneID              int64        `db:"plane_id"`
}

// FlightPrice represents the price of a flight and its metadata.
type FlightPrice struct {
	ID        int64        `db:"id"`
	Price     float64      `db:"price"`
	Class     string       `db:"class"` // enum (economy, business, first)
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
	FlightID  int64        `db:"flight_id"`
}

// FlightReservation represents a flight reservation entity in the database.
type FlightReservation struct {
	ID           int64        `db:"id"`
	Class        string       `db:"class"` // enum (business, economy)
	ReservedSeat int          `db:"reserved_seat"`
	TotalSeat    int          `db:"total_seat"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
	FlightID     int64        `db:"flight_id"`
}

// Airport represents an airport entity in the database.
type Airport struct {
	ID        int64          `db:"id"`
	Name      string         `db:"name"`
	City      string         `db:"city"`
	Province  string         `db:"province"`
	IATA      sql.NullString `db:"iata"`
	ICAO      sql.NullString `db:"icao"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	DeletedAt sql.NullTime   `db:"deleted_at"`
}

// Booking represents a booking made by a user for a flight.
type Booking struct {
	ID               int64        `db:"id"`
	CodeBooking      string       `db:"code_booking"`
	BookingDate      time.Time    `db:"booking_date"`
	PaymentExpiredAt time.Time    `db:"payment_expired_at"`
	BookingStatus    string       `db:"booking_status"` // enum (pending, paid, canceled)
	CaseID           int64        `db:"case_id"`
	CreatedAt        time.Time    `db:"created_at"`
	UpdatedAt        sql.NullTime `db:"updated_at"`
	DeletedAt        sql.NullTime `db:"deleted_at"`
	UserID           int64        `db:"user_id"`
	FlightID         int64        `db:"flight_id"`
}

// BookingDetail represents the details of a booking made by a passenger.
type BookingDetail struct {
	ID              int64        `db:"id"`
	PassengerID     int64        `db:"passenger_id"`
	SeatNumber      string       `db:"seat_number"`
	BaggageCapacity int          `db:"baggage_capacity"`
	Class           string       `db:"class"` // enum (economy, business, first)
	CreatedAt       time.Time    `db:"created_at"`
	UpdatedAt       sql.NullTime `db:"updated_at"`
	DeletedAt       sql.NullTime `db:"deleted_at"`
	BookingID       int64        `db:"booking_id"`
}

// Payment represents a payment made by a user for a booking.
type Payment struct {
	ID            int64      `db:"id"`
	InvoiceNumber string     `db:"invoice_number"`
	TotalPayment  float64    `db:"total_payment"`
	PaymentMethod string     `db:"payment_method"` // enum (credit_card, debit_card, bank_transfer)
	PaymentDate   time.Time  `db:"payment_date"`
	PaymentStatus string     `db:"payment_status"` // enum (pending, paid, canceled)
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at"`
	DeletedAt     *time.Time `db:"deleted_at"`
	BookingID     int64      `db:"booking_id"`
}

type PaymentMethod struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	IsActive  bool         `db:"is_active"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

// Ticket represents a flight ticket entity
type Ticket struct {
	ID                 int64      `db:"id"`
	CodeTicket         string     `db:"code_ticket"`
	IsBoardingPass     bool       `db:"is_boarding_pass"`
	IsEligibleToFlight bool       `db:"is_eligible_to_flight"`
	CreatedAt          time.Time  `db:"created_at"`
	UpdatedAt          time.Time  `db:"updated_at"`
	DeletedAt          *time.Time `db:"deleted_at"`
	BookingID          int64      `db:"booking_id"`
}

type WorkflowDetail struct {
	ID        int64      `db:"id"`
	CaseID    int64      `db:"case_id"`
	TaskName  string     `db:"task_name"`
	TaskID    string     `db:"task_id"`
	Status    string     `db:"status"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
