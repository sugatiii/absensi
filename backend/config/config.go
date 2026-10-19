package config

import "os"

type Config struct {
	DatabaseURL string
	JWTSecret   string
	ServerPort  string
}

func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/absensi?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "secret-key-absensi-2024"),
		ServerPort:  getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}