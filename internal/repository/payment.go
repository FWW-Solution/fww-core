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

	err := r.db.QueryRowx(query).StructScan(&row)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return entity.Payment{}, nil
	}

	if err != nil {
		return entity.Payment{}, err
	}

	return row, nil

}

// UpdatePayment implements Repository.
func (r *repository) UpsertPayment(data *entity.Payment) (int64, error) {
	var query string
	if data.ID == 0 {
		query = fmt.Sprintf(`INSERT INTO payments (invoice_number, total_payment, payment_method, payment_date, payment_status, booking_id) VALUES ('%s', %f, '%s', '%s', '%s', %d) ON CONFLICT (id) DO UPDATE SET payment_status = '%s', updated_at = NOW() WHERE payments.id = %d RETURNING id`, data.InvoiceNumber, data.TotalPayment, data.PaymentMethod, data.PaymentDate.Format("2006-01-02 15:04:05"), data.PaymentStatus, data.BookingID, data.PaymentStatus, data.ID)
	} else {
		query = fmt.Sprintf(`INSERT INTO payments (id, invoice_number, total_payment, payment_method, payment_date, payment_status, booking_id) VALUES (%d,'%s', %f, '%s', '%s', '%s', %d) ON CONFLICT (id) DO UPDATE SET payment_status = '%s', updated_at = NOW() WHERE payments.id = %d RETURNING id`, data.ID, data.InvoiceNumber, data.TotalPayment, data.PaymentMethod, data.PaymentDate.Format("2006-01-02 15:04:05"), data.PaymentStatus, data.BookingID, data.PaymentStatus, data.ID)
	}

	// do sqlx transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	var id int64
	err = tx.QueryRowx(query).Scan(&id)

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

// FindPaymentMethodStatus implements Repository.
func (r *repository) FindPaymentMethodStatus() ([]entity.PaymentMethod, error) {
	query := `SELECT id, name, is_active FROM payment_methods`
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
