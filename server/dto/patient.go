package dto

import (
	"time"

	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/db/models"
)

type PatientDTO struct {
	MRN         int
	FirstName   string
	LastName    string
	DateOfBirth *time.Time
}

func NewPatientDTO_FM(model *models.Patient) *PatientDTO {
	return &PatientDTO{
		MRN:         model.MRN,
		FirstName:   model.FirstName,
		LastName:    model.LastName,
		DateOfBirth: model.DateOfBirth,
	}
}

type PatientListDTO struct {
	Count int
	List  []PatientDTO
}

func NewPatientListDTO_FM(models []models.Patient, count int) *PatientListDTO {
	list := make([]PatientDTO, 0)

	for _, p := range models {
		list = append(list, *NewPatientDTO_FM(&p))
	}

	return &PatientListDTO{
		Count: count,
		List:  list,
	}
}
