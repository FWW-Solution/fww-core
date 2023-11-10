package dto_airport

type ResponseAirport struct {
	City      string `json:"city"`
	CreatedAt string `json:"created_at"`
	Iata      string `json:"iata"`
	Icao      string `json:"icao"`
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Province  string `json:"province"`
	UpdatedAt string `json:"updated_at"`
}
