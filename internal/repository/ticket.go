package repository

import (
	"fmt"
	"fww-core/internal/entity"
)

// UpsertTicket implements Repository.
func (r *repository) UpsertTicket(data *entity.Ticket) (int64, error) {
	var query string
	if data.ID == 0 {
		query = fmt.Sprintf(`INSERT INTO tickets (booking_id, code_ticket, is_boarding_pass, is_eligible_to_flight) VALUES (%d, '%s', %v, %v) ON CONFLICT (id) DO UPDATE SET is_boarding_pass = '%v', is_eligible_to_flight = '%v', updated_at = NOW() WHERE tickets.id = %d RETURNING id`, data.BookingID, data.CodeTicket, data.IsBoardingPass, data.IsEligibleToFlight, data.IsBoardingPass, data.IsEligibleToFlight, data.ID)
	} else {
		query = fmt.Sprintf(`INSERT INTO tickets (id, booking_id, code_ticket, is_boarding_pass, is_eligible_to_flight) VALUES (%d, %d, '%s', %v, %v) ON CONFLICT (id) DO UPDATE SET is_boarding_pass = '%v', is_eligible_to_flight = '%v', updated_at = NOW() WHERE tickets.id = %d RETURNING id`, data.ID, data.BookingID, data.CodeTicket, data.IsBoardingPass, data.IsEligibleToFlight, data.IsBoardingPass, data.IsEligibleToFlight, data.ID)
	}

	fmt.Println(query)

	// do sqlx transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	var id int64
	err = tx.QueryRowx(query).Scan(&id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return id, nil
}
