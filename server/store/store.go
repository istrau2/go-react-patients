package store

import (
	"github.com/jinzhu/gorm"
)

type Store struct {
	Patient *PatientStore
}

func New(db *gorm.DB) *Store {
	return &Store{
		Patient: NewPatientStore(db),
	}
}
