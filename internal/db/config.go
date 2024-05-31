package db

const (
	DefaultDBFileName = string("db")
)

type Config struct {
	FileName string
}

func defaultConfig() Config {
	return Config{
		FileName: DefaultDBFileName,
	}
}
