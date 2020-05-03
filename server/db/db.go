package db

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/db/models"
)

func New(opts *Options) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	switch opts.Dialect {
	case DialectSQLLite:
		db, err = gorm.Open("sqlite3", opts.File)
	default:
		err = errors.New("Unrecognized database dialect")
	}

	// db.DB().SetMaxIdleConns(opts.MaxIdleConns)

	db.LogMode(opts.LogLevel == LogLevelErrors)

	db.AutoMigrate(
		&models.Patient{},
	)

	return db, err
}
