package dto_passanger

type RequestRegister struct {
	FullName    string `json:"full_name" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	IDNumber    string `json:"id_number" validate:"required,min=16,max=16"`
	IDType      string `json:"id_type" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
}

type RequestUpdate struct {
	DateOfBirth string `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
	FullName    string `json:"full_name" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	ID          int64  `json:"id" validate:"required,numeric"`
	IDNumber    string `json:"id_number" validate:"required,min=16,max=16"`
	IDType      string `json:"id_type" validate:"required"`
}

type RequestBPM struct {
	IDNumber string `json:"id_number"`
}

type RequestUpdateBPM struct {
	IDNumber           string `json:"id_number"`
	VaccineStatus      string `json:"vaccine_status"`
	IsVerifiedDukcapil bool   `json:"is_verified_dukcapil"`
	CaseID             int64  `json:"case_id"`
}

type ResponseRegistered struct {
	ID int64 `json:"id"`
}

type ResponseDetail struct {
	CovidVaccineStatus string `json:"covid_vaccine_status"`
	CreatedAt          string `json:"created_at"`
	DateOfBirth        string `json:"date_of_birth"`
	FullName           string `json:"full_name"`
	Gender             string `json:"gender"`
	ID                 int64  `json:"id"`
	IDNumber           string `json:"id_number"`
	IDType             string `json:"id_type"`
	IsIDVerified       bool   `json:"is_id_verified"`
	UpdatedAt          string `json:"updated_at"`
}

type ResponseUpdate struct {
	ID int64 `json:"id"`
}
