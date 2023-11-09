package dto_passanger

type RequestRegister struct {
	DateOfBirth string `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
	FullName    string `json:"full_name" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	IDNumber    string `json:"id_number" validate:"required,numeric,min=16,max=16"`
	IDType      string `json:"id_type" validate:"required"`
}

type RequestUpdate struct {
	DateOfBirth string `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
	FullName    string `json:"full_name" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
	ID          int64  `json:"id" validate:"required,numeric"`
	IDNumber    string `json:"id_number" validate:"required,numeric,min=16,max=16"`
	IDType      string `json:"id_type" validate:"required"`
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
	IsIDVerified       string `json:"is_id_verified"`
	UpdatedAt          string `json:"updated_at"`
}

type ResponseUpdate struct {
	CovidVaccineStatus string `json:"covid_vaccine_status"`
	CreatedAt          string `json:"created_at"`
	DateOfBirth        string `json:"date_of_birth"`
	FullName           string `json:"full_name"`
	Gender             string `json:"gender"`
	ID                 int64  `json:"id"`
	IDNumber           string `json:"id_number"`
	IDType             string `json:"id_type"`
	IsIDVerified       string `json:"is_id_verified"`
	UpdatedAt          string `json:"updated_at"`
}
