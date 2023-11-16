package repository

import (
	"fmt"
	"fww-core/internal/entity"
)

// FindPaymentDetailByInvoice implements Repository.
func (r *repository) FindPaymentDetailByInvoice(invoiceNumber string) (entity.Payment, error) {
	query := fmt.Sprintf(`SELECT id, invoice_number, total_payment, payment_method, payment_date, payment_status, created_at, updated_at, deleted_at, booking_id FROM payments WHERE invoice_number = '%s'`, invoiceNumber)

	// hanldle entity
	var row entity.Payment

	result, err := r.db.Queryx(query)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return entity.Payment{}, nil
	}

	if err != nil {
		return entity.Payment{}, err
	}

	for result.Next() {
		err := result.StructScan(&row)
		if err != nil {
			return entity.Payment{}, err
		}
	}

	return row, nil

}

// UpdatePayment implements Repository.
func (r *repository) UpsertPayment(data *entity.Payment) (int64, error) {
	query := fmt.Sprintf(`INSERT INTO payments (invoice_number, total_payment, payment_method, payment_date, payment_status, booking_id) VALUES ('%s', %f, '%s', '%s', '%s', %d) ON CONFLICT (conflict_target) DO UPDATE payments SET payment_status = '%s' WHERE id = %d`, data.InvoiceNumber, data.TotalPayment, data.PaymentMethod, data.PaymentDate, data.PaymentStatus, data.BookingID, data.PaymentStatus, data.ID)

	result, err := r.db.Exec(query)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// FindPaymentMethodStatus implements Repository.
func (r *repository) FindPaymentMethodStatus() ([]entity.PaymentMethod, error) {
	query := fmt.Sprintf(`SELECT id, name, is_active FROM payment_methods`)
	var rows []entity.PaymentMethod

	result, err := r.db.Queryx(query)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return []entity.PaymentMethod{}, nil
	}

	if err != nil {
		return []entity.PaymentMethod{}, err
	}

	for result.Next() {
		var row entity.PaymentMethod
		err := result.StructScan(&row)
		if err != nil {
			return []entity.PaymentMethod{}, err
		}
		rows = append(rows, row)
	}

	return rows, nil
}
