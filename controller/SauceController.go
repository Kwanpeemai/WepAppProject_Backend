package controller

import (
	"database/sql"
	"log"
	"net/http"
	"web-project/models"

	"github.com/gin-gonic/gin"
)
type Sauces struct {
	Sauces []models.Sauce `json:"sauces"`
}

func CreateSauce(c *gin.Context, db *sql.DB) {
	var sauces Sauces

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&sauces); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(sauces.Sauces) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No sauces provided"})
		return
	}

	var errors []error

	for _, sauce := range sauces.Sauces {
		insertQuery := "INSERT INTO sauce (Sauce_name_th, Sauce_name_en, Sauce_price, Sauce_Stock) VALUES (?, ?, ?, ?)"
		_, err := db.Exec(insertQuery, sauce.Sauce_name_th, sauce.Sauce_name_en, sauce.Sauce_price, sauce.Sauce_Stock)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data", "details": errors})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sauces created successfully"})
}

func GetSauces(c *gin.Context, db *sql.DB) {
	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Query the database
	rows, err := db.Query("SELECT * FROM sauce")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
	defer rows.Close()

	var sauces []models.Sauce
	for rows.Next() {
		var sauce models.Sauce
		err := rows.Scan(&sauce.Sauce_ID, &sauce.Sauce_name_th, &sauce.Sauce_name_en, &sauce.Sauce_price, &sauce.Sauce_Stock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		sauces = append(sauces, sauce)
	}

	c.JSON(http.StatusOK, sauces)
}

func GetSauce(c *gin.Context, db *sql.DB) {
	sauceID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var sauce models.Sauce
	err := db.QueryRow("SELECT * FROM sauce WHERE Sauce_ID = ?", sauceID).Scan(&sauce.Sauce_ID, &sauce.Sauce_name_th, &sauce.Sauce_name_en, &sauce.Sauce_price, &sauce.Sauce_Stock)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}

	c.JSON(http.StatusOK, sauce)
}

func UpdateSauce(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var sauce models.Sauce

	if err := c.ShouldBindJSON(&sauce); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update database
	updateQuery := "UPDATE sauce SET Sauce_name_th=?, Sauce_name_en=?, Sauce_price=?, Sauce_Stock=? WHERE Sauce_ID=?"
	_, err := db.Exec(updateQuery, sauce.Sauce_name_th, sauce.Sauce_name_en, sauce.Sauce_price, sauce.Sauce_Stock, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sauce updated successfully"})
}

func DeleteSauce(c *gin.Context, db *sql.DB) {
	sauceID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	deleteQuery := "DELETE FROM sauce WHERE Sauce_ID = ?"
	_, err := db.Exec(deleteQuery, sauceID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	resetQuery := "ALTER TABLE sauce AUTO_INCREMENT = 1"
	_, err = db.Exec(resetQuery)
	if err != nil {
		log.Printf("Error resetting auto-increment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error resetting auto-increment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sauce deleted successfully"})
}
