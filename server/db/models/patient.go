package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Patient struct {
	gorm.Model
	MRN         int
	FirstName   string
	LastName    string
	DateOfBirth *time.Time

}
