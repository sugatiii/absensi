package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/sugatiii/absensi/database"
)

func ListUsers(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, nama, email, role, created_at FROM users ORDER BY id")
	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil data"})
		return
	}
	defer rows.Close()
	var users []gin.H
	for rows.Next() {
		var id int
		var nama, email, role, created string
		rows.Scan(&id, &nama, &email, &role, &created)
		users = append(users, gin.H{"id": id, "nama": nama, "email": email, "role": role, "created_at": created})
	}
	c.JSON(200, gin.H{"users": users})
}

func CreateUser(c *gin.Context) {
	var input struct { Nama, Email, Password, Role string }
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Input tidak valid"})
		return
	}
	_, err := database.DB.Exec("INSERT INTO users (nama, email, password, role) VALUES ($1,$2,$3,$4)", input.Nama, input.Email, input.Password, input.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal buat user"})
		return
	}
	c.JSON(201, gin.H{"message": "User berhasil dibuat"})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input struct { Nama, Email, Role string }
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Input tidak valid"})
		return
	}
	_, err := database.DB.Exec("UPDATE users SET nama=$1, email=$2, role=$3 WHERE id=$4", input.Nama, input.Email, input.Role, id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal update user"})
		return
	}
	c.JSON(200, gin.H{"message": "User berhasil diupdate"})
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := database.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal hapus user"})
		return
	}
	c.JSON(200, gin.H{"message": "User berhasil dihapus"})
}