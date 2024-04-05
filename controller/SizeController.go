package controller

import (
	"database/sql"
	"log"
	"net/http"
	"web-project/models"

	"github.com/gin-gonic/gin"
) 

var sizes []models.Size

func CreateSize(c *gin.Context, db *sql.DB) {
	var size models.Size

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert into database
	insertQuery := "INSERT INTO size (Size_name_th, Size_name_en, Size_price, Size_Stock) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(insertQuery, size.Size_name_th, size.Size_name_en, size.Size_price, size.Size_Stock)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Size created successfully"})
}

func GetSizes(c *gin.Context, db *sql.DB) {
	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Query the database
	rows, err := db.Query("SELECT * FROM size")
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

func GetSize(c *gin.Context, db *sql.DB) {
	sizeID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var size models.Size
	err := db.QueryRow("SELECT * FROM size WHERE Size_ID = ?", sizeID).Scan(&size.Size_ID, &size.Size_name_th, &size.Size_name_en, &size.Size_price, &size.Size_Stock)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}

	c.JSON(http.StatusOK, size)
}

func UpdateSize(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var size models.Size

	if err := c.ShouldBindJSON(&size); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update database
	updateQuery := "UPDATE size SET Size_name_th=?, Size_name_en=?, Size_price=?, Size_Stock=? WHERE Size_ID=?"
	_, err := db.Exec(updateQuery, size.Size_name_th, size.Size_name_en, size.Size_price, size.Size_Stock, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Size updated successfully"})
}

func DeleteSize(c *gin.Context, db *sql.DB) {
	sizeID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	deleteQuery := "DELETE FROM size WHERE Size_ID = ?"
	_, err := db.Exec(deleteQuery, sizeID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Size deleted successfully"})
}
