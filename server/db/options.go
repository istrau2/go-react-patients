package db

type Dialect string

const (
	DialectSQLLite Dialect = "sqlite"
)

type LogLevel int

const (
	LogLevelNone   LogLevel = -1
	LogLevelUnset  LogLevel = 0
	LogLevelErrors LogLevel = 1
)

type Options struct {
	MaxIdleConns int
	LogLevel     LogLevel
	Dialect      Dialect
	File         string
}

func (opts *Options) WithDefaults() *Options {
	if opts.LogLevel == LogLevelUnset {
		opts.LogLevel = LogLevelErrors
	}

	if opts.MaxIdleConns == 0 {
		opts.MaxIdleConns = 2
	}

	if opts.Dialect == DialectSQLLite && opts.File == "" {
		opts.File = "./sqllite.db"
	}

	return opts
}
