package db

type DB struct {
	Config

	db *string
}

func New(cfg ...Config) *DB {
	c := defaultConfig()

	if len(cfg) > 0 {
		c = cfg[0]
	}

	return &DB{
		Config: c,
	}
}
