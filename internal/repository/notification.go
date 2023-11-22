package repository

import "fww-core/internal/data/dto_notification"

// PaymentInvoiceReportByBookingID implements Repository.
func (r *repository) PaymentInvoiceReportByBookingCode(bookingCode string) (dto_notification.PaymentInvoiceAggregator, error) {
	panic("unimplemented")
}

// PaymentReceiptReportByBookingID implements Repository.
func (*repository) PaymentReceiptReportByBookingCode(bookingCode string) (dto_notification.PaymentReceiptAggregator, error) {
	panic("unimplemented")
}

// TicketRedeemedReportByBookingID implements Repository.
func (*repository) TicketRedeemedReportByBookingCode(bookingCode string) (dto_notification.TicketRedeemAgregator, error) {
	panic("unimplemented")
}
