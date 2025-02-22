package config

import "os"

type Config struct {
	ServerPort    string
	MySQLDSN      string
	MPAccessToken string
}

func LoadConfig() Config {
	return Config{
		ServerPort:    os.Getenv("SERVER_PORT"),
		MySQLDSN:      os.Getenv("MYSQL_DSN"),
		MPAccessToken: os.Getenv("MP_ACCESS"),
	}
}
