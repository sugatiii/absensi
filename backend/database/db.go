package database

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(dsn string) {
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Gagal konek database:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Gagal ping database:", err)
	}
	log.Println("Database connected!")
}

func Migrate() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, nama VARCHAR(100), email VARCHAR(100) UNIQUE, password VARCHAR(255), role VARCHAR(20) DEFAULT 'karyawan', created_at TIMESTAMP DEFAULT NOW())`,
		`CREATE TABLE IF NOT EXISTS absensi (id SERIAL PRIMARY KEY, user_id INT REFERENCES users(id), type VARCHAR(10), timestamp TIMESTAMP DEFAULT NOW())`,
	}
	for _, q := range queries {
		if _, err := DB.Exec(q); err != nil {
			log.Println("Migrasi error:", err)
		}
	}
	log.Println("Migrasi selesai!")
}