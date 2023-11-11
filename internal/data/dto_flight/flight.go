package dto_flight

type ResponseFlightDetail struct {
	ArrivalAirportName  string  `json:"arrival_airport_name"`
	ArrivalTime         string  `json:"arrival_time"`
	CodeFlight          string  `json:"code_flight"`
	DepartureTime       string  `json:"departure_time"`
	DepatureAirportName string  `json:"depature_airport_name"`
	FlightPrice         float64 `json:"flight_price"`
	ReminingSeat        int     `json:"remining_seat"`
	Status              string  `json:"status"`
}

type ResponseFlight struct {
	ArrivalAirportName  string  `json:"arrival_airport_name"`
	ArrivalTime         string  `json:"arrival_time"`
	CodeFlight          string  `json:"code_flight"`
	DepartureTime       string  `json:"departure_time"`
	DepatureAirportName string  `json:"depature_airport_name"`
	FlightPrice         float64 `json:"flight_price"`
	ReminingSeat        int     `json:"remining_seat"`
	Status              string  `json:"status"`
}
