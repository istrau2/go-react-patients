package env

import (
	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/db"
)

type Env struct {
	DB *db.Options
}

func New() (*Env, error) {
	dbOptions, err := NewDBOptions()

	if (err != nil) {
		return nil, err
	}

	return &Env{
		DB: dbOptions,
	}, nil
}