package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"web-project/models"

	"github.com/gin-gonic/gin"
)

var sizes []models.Size

func CreateSize(c *gin.Context, db *sql.DB) {
	var size models.Size
	fmt.Println("🦏🦏🦏🦏🦏🦏🦏🦏🦏🦏")
	if err := c.ShouldBindJSON(&size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("🦏🦏🦏🦏🦏🦏🦏🦏🦏🦏")

	// Insert into database
	insertQuery := "INSERT INTO sizes (Size_name_th, Size_name_en, Size_price, Size_Stock) VALUES (?, ?, ?, ?)"
	fmt.Println("🦏🦏🦏🦏🦏🦏🦏🦏🦏🦏")
	_, err := db.Exec(insertQuery, size.Size_name_th, size.Size_name_en, size.Size_price, size.Size_Stock)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Size created successfully"})
}

func GetSizes(c *gin.Context, db *sql.DB) {
	// Query the database
	rows, err := db.Query("SELECT * FROM sizes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
	defer rows.Close()

	var sizes []models.Size
	for rows.Next() {
		var size models.Size
		err := rows.Scan(&size.Size_ID, &size.Size_name_th, &size.Size_name_en, &size.Size_price, &size.Size_Stock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		sizes = append(sizes, size)
	}

	c.JSON(http.StatusOK, sizes)
}
