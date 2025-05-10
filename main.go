package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Email struct {
	Email string `json:"email"`
}

func main() {
	InitDB() // DB初期化

	r := gin.Default()

	r.GET("/emails", func(c *gin.Context) {
		rows, err := DB.Query("SELECT email FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "データ取得エラー"})
			return
		}
		defer rows.Close()

		var emails []Email
		for rows.Next() {
			var e Email
			if err := rows.Scan(&e.Email); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "スキャンエラー"})
				return
			}
			emails = append(emails, e)
		}

		c.JSON(http.StatusOK, emails)
	})

	r.Run(":8080") // localhost:8080で起動
}