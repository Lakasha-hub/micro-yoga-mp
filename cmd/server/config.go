package main

import "os"

type Config struct {
	ServerPort    string
	MySQLDSN      string
	MPAccessToken string
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
}

func LoadConfig() Config {
	return Config{
		ServerPort:    os.Getenv("SERVER_PORT"),
		MPAccessToken: os.Getenv("MP_ACCESS"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBName:        os.Getenv("DB_NAME"),
	}
}
