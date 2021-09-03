package store

type Config struct {
	DataBaseURL string `toml:"db_url"`
}

func NewConfig() *Config {
	return &Config{
		DataBaseURL: "mongodb://localhost",
	}
}
