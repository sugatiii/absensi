package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sugatiii/absensi/database"
)

func DashboardSummary(c *gin.Context) {
	var totalKaryawan, hadirHariIni int
	database.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&totalKaryawan)
	database.DB.QueryRow("SELECT COUNT(DISTINCT user_id) FROM absensi WHERE type='in' AND DATE(timestamp)=CURRENT_DATE").Scan(&hadirHariIni)
	c.JSON(200, gin.H{"total_karyawan": totalKaryawan, "hadir_hari_ini": hadirHariIni})
}