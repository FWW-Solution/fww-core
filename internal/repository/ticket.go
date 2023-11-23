package repository

import (
	"fmt"
	"fww-core/internal/entity"
)

// UpsertTicket implements Repository.
func (r *repository) UpsertTicket(data *entity.Ticket) (int64, error) {
	var query string
	if data.ID == 0 {
		query = fmt.Sprintf(`INSERT INTO tickets (booking_id, code_ticket, is_boarding_pass, is_eligible_to_flight, boarding_time) VALUES (%d, '%s', %v, %v, '%s') ON CONFLICT (id) DO UPDATE SET is_boarding_pass = '%v', is_eligible_to_flight = '%v', boarding_time = '%s', updated_at = NOW() WHERE tickets.id = %d RETURNING id`, data.BookingID, data.CodeTicket, data.IsBoardingPass, data.IsEligibleToFlight, data.BoardingTime.Time.Format("2006-01-02 15:04:05"), data.IsBoardingPass, data.IsEligibleToFlight, data.BoardingTime.Time.Format("2006-01-02 15:04:05"), data.ID)
	} else {
		query = fmt.Sprintf(`INSERT INTO tickets (id, booking_id, code_ticket, is_boarding_pass, is_eligible_to_flight) VALUES (%d, %d, '%s', %v, %v) ON CONFLICT (id) DO UPDATE SET is_boarding_pass = '%v', is_eligible_to_flight = '%v', boarding_time = '%s', updated_at = NOW() WHERE tickets.id = %d RETURNING id`, data.ID, data.BookingID, data.CodeTicket, data.IsBoardingPass, data.IsEligibleToFlight, data.IsBoardingPass, data.IsEligibleToFlight, data.BoardingTime.Time.Format("2006-01-02 15:04:05"), data.ID)
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

// FindTicketByCodeTicket implements Repository.
func (r *repository) FindTicketByCodeTicket(codeTicket string) (entity.Ticket, error) {
	query := fmt.Sprintf(`SELECT id, booking_id, code_ticket, is_boarding_pass, is_eligible_to_flight, created_at, updated_at FROM tickets WHERE code_ticket = '%s'`, codeTicket)
	var ticket entity.Ticket
	err := r.db.QueryRowx(query).StructScan(&ticket)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return entity.Ticket{}, nil
	}

	if err != nil {
		return entity.Ticket{}, err
	}

	return ticket, nil
}
