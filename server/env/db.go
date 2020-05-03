package env

import (
	"errors"
	"os"

	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/db"
)

func NewDBOptions() (*db.Options, error) {
	var (
		opts *db.Options
		err  error
	)

	dbEnv := os.Getenv("DB_DIALECT")

	if dbEnv == "" {
		dbEnv = string(db.DialectSQLLite)
	}

	switch dbEnv {
	case string(db.DialectSQLLite):
		opts = (&db.Options{
			Dialect: db.DialectSQLLite,
		}).WithDefaults()
	default:
		err = errors.New("Unrecognized DB_DIALECT environment variable")
	}

	return opts, err
}
