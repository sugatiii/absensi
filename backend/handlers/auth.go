package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sugatiii/absensi/database"
	"github.com/sugatiii/absensi/middleware"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct { Nama, Email, Password string }
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Input tidak valid"})
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	_, err := database.DB.Exec("INSERT INTO users (nama, email, password) VALUES ($1,$2,$3)", input.Nama, input.Email, string(hash))
	if err != nil {
		c.JSON(409, gin.H{"error": "Email sudah terdaftar"})
		return
	}
	c.JSON(201, gin.H{"message": "Registrasi berhasil"})
}

func Login(c *gin.Context) {
	var input struct { Email, Password string }
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Input tidak valid"})
		return
	}
	var id int
	var nama, email, password, role string
	err := database.DB.QueryRow("SELECT id, nama, email, password, role FROM users WHERE email=$1", input.Email).Scan(&id, &nama, &email, &password, &role)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(password), []byte(input.Password)) != nil {
		c.JSON(401, gin.H{"error": "Email atau password salah"})
		return
	}
	token, _ := middleware.GenerateToken(id, email, role)
	c.JSON(200, gin.H{"token": token, "user": gin.H{"id": id, "nama": nama, "email": email, "role": role}})
}