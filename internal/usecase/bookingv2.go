package usecase

// // RequestBooking implements UseCase.
// func (u *useCase) RequestBooking(data *dto_booking.Request, bookingIDCode string) error {
// 	ctx := context.Background()
// 	// Check Booking ID Code
// 	resultBooking, err := u.repository.GetBookingByBookingIDCode(bookingIDCode)
// 	if err != nil {
// 		return err
// 	}

// 	if resultBooking.ID != 0 {
// 		return errors.New("booking id code already exist")
// 	}

// 	// Check Remining Seat Each Class
// 	for _, v := range data.BookDetails {
// 		flightIDReminingSeat := "flight-" + string(data.FlightID) + "-" + v.Class
// 		result := u.redis.Get(ctx, flightIDReminingSeat)
// 		if result.Err() != nil {
// 			resultSeat, err := u.repository.FindReminingSeat(data.FlightID, v.Class)
// 			if err != nil {
// 				return err
// 			}
// 			if resultSeat <= 0 {
// 				return errors.New("no remaning seat for class " + v.Class + "")
// 			}
// 			u.redis.Set(ctx, flightIDReminingSeat, resultSeat, 0)
// 		}

// 		resultSeat, err := result.Int()
// 		if err != nil {
// 			return err
// 		}

// 		if resultSeat <= 0 {
// 			return errors.New("no remaning seat for class " + v.Class + "")
// 		}
// 	}

// 	flightIDKey := "flight-" + string(data.FlightID)

// 	// Lock Transaction Redis
// 	rc := redis.InitMutex(flightIDKey)
// 	redis.LockMutex(rc)
// 	defer redis.UnlockMutex(rc)

// 	// Check Remining Seat

// 	// Insert Booking
// 	bookingEntity := entity.Booking{
// 		CodeBooking:      bookingIDCode,
// 		BookingDate:      time.Now(),
// 		PaymentExpiredAt: time.Now().Add(time.Hour * 24),
// 		BookingStatus:    "pending",
// 		CaseID:           0,
// 		UserID:           data.UserID,
// 		FlightID:         data.FlightID,
// 	}

// 	bookingID, err := u.repository.InsertBooking(&bookingEntity)
// 	if err != nil {
// 		return err
// 	}

// 	// Insert Booking Detail
// 	for _, v := range data.BookDetails {
// 		flightIDReminingSeat := "flight-" + string(data.FlightID) + "-" + v.Class

// 		entityBookingDetail := entity.BookingDetail{
// 			BookingID:      bookingID,
// 			PassengerID:    v.PassangerID,
// 			SeatNumber:     v.SeatNumber,
// 			BagageCapacity: v.Baggage,
// 			Class:          v.Class,
// 		}

// 		_, err := u.repository.InsertBookingDetail(&entityBookingDetail)
// 		if err != nil {
// 			return err
// 		}

// 		result := u.redis.Get(ctx, flightIDReminingSeat)
// 		resultSeat, err := result.Int()
// 		if err != nil {
// 			return err
// 		}
// 		// Update Flight Reservation
// 		entityReservation := entity.FlightReservation{
// 			ID:    bookingID,
// 			Class: v.Class,
// 			// each class have total seat different
// 			ReservedSeat: 172 - resultSeat,
// 			// each class have total seat different
// 			TotalSeat: 172,
// 			UpdatedAt: sql.NullTime{
// 				Time:  time.Now(),
// 				Valid: true,
// 			},
// 			DeletedAt: sql.NullTime{
// 				Time:  time.Time{},
// 				Valid: false,
// 			},
// 			FlightID: data.FlightID,
// 		}
// 		_, err = u.repository.UpdateFlightReservation(&entityReservation)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return err

// }
