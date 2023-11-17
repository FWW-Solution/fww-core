package dto_ticket

type Request struct {
	CodeBooking string `json:"code_booking"`
}

type Response struct {
	BordingTime string `json:"bording_time"`
	CodeTicket  string `json:"code_ticket"`
}
