package usecase

import "fww-core/internal/data/dto_airport"

// GetAirport implements UseCase.
func (u *useCase) GetAirport(city string, province string, iata string) ([]dto_airport.ResponseAirport, error) {
	result, err := u.repository.FindAirport(city, province, iata)
	if err != nil {
		return nil, err
	}

	var response []dto_airport.ResponseAirport

	for _, v := range result {
		response = append(response, dto_airport.ResponseAirport{
			ID:        v.ID,
			City:      v.City,
			Province:  v.Province,
			Name:      v.Name,
			Iata:      v.IATA,
			Icao:      v.ICAO,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return response, nil
}
