package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sugatiii/absensi/database"
)

func CheckIn(c *gin.Context) {
	userID := c.GetInt("user_id")
	_, err := database.DB.Exec("INSERT INTO absensi (user_id, type) VALUES ($1, 'in')", userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal check-in"})
		return
	}
	c.JSON(200, gin.H{"message": "Check-in berhasil!"})
}

func CheckOut(c *gin.Context) {
	userID := c.GetInt("user_id")
	_, err := database.DB.Exec("INSERT INTO absensi (user_id, type) VALUES ($1, 'out')", userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal check-out"})
		return
	}
	c.JSON(200, gin.H{"message": "Check-out berhasil!"})
}

func AbsensiHistory(c *gin.Context) {
	userID := c.GetInt("user_id")
	rows, err := database.DB.Query("SELECT id, type, timestamp FROM absensi WHERE user_id=$1 ORDER BY timestamp DESC LIMIT 20", userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal ambil riwayat"})
		return
	}
	defer rows.Close()
	var history []gin.H
	for rows.Next() {
		var id int
		var tipe, ts string
		rows.Scan(&id, &tipe, &ts)
		history = append(history, gin.H{"id": id, "type": tipe, "timestamp": ts})
	}
	c.JSON(200, gin.H{"history": history})
}