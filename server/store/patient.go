package store

import (
	"strings"

	"github.com/jinzhu/gorm"

	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/db/models"
)

type PatientStore struct {
	db *gorm.DB
}

func NewPatientStore(db *gorm.DB) *PatientStore {
	return &PatientStore{
		db: db,
	}
}

func (ps *PatientStore) ByParialName(match string) ([]models.Patient, int, error) {
	var (
		patients   []models.Patient
		count      int
		dbInstance *gorm.DB
	)

	if match != "" {
		dbInstance = ps.db.Where("first_name LIKE ?", match).Or("last_name LIKE ?", "%"+strings.ToLower(match)+"%")
	}

	if dbInstance == nil {
		dbInstance = ps.db
	}

	dbInstance = dbInstance.Find(&patients).Count(&count)
	if dbInstance.Error != nil {
		return nil, 0, dbInstance.Error
	}

	return patients, count, nil
}
