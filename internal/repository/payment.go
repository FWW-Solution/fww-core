package repository

import "fww-core/internal/entity"

// FindPaymentDetailByInvoice implements Repository.
func (r *repository) FindPaymentDetailByInvoice(invoiceNumber string) (entity.Payment, error) {
	panic("unimplemented")
}

// UpdatePayment implements Repository.
func (r *repository) UpdatePayment(data *entity.Payment) (int64, error) {
	panic("unimplemented")
}

// FindPaymentMethodStatus implements Repository.
func (*repository) FindPaymentMethodStatus() ([]entity.PaymentMethod, error) {
	panic("unimplemented")
}
