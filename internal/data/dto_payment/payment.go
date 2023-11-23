package dto_payment

type Request struct {
	BookingID     int64  `json:"booking_id"`
	PaymentMethod string `json:"payment_method"`
}

type RequestInvoice struct {
	CaseID      int64  `json:"case_id"`
	CodeBooking string `json:"code_booking"`
}

type RequestUpdatePayment struct {
	InvoiceNumber string `json:"invoice_number"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
}

type DoPayment struct {
	CaseID        int64   `json:"case_id"`
	InvoiceNumber string  `json:"invoice_number"`
	PaymentMethod string  `json:"payment_method"`
	PaymentAmount float64 `json:"payment_ammount"`
}

type AsyncPaymentResponse struct {
	PaymentIDCode string `json:"payment_id_code"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type MethodResponse struct {
	ID       int64  `json:"id"`
	IsActive bool   `json:"is_active"`
	Name     string `json:"name"`
}
