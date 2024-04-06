package controller

import (
	"database/sql"
	"log"
	"net/http"
	"web-project/models"

	"github.com/gin-gonic/gin"
)

type Flavors struct {
	Flavors []models.Flavor `json:"flavors"`
}

var flavors []models.Flavor

func CreateFlavor(c *gin.Context, db *sql.DB) {
	var flavors Flavors

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&flavors); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(flavors.Flavors) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No flavors provided"})
		return
	}

	var errors []error

	for _, flavor := range flavors.Flavors {
		insertQuery := "INSERT INTO flavor (Flavor_name_th, Flavor_name_en, Flavor_price, Flavor_Stock) VALUES (?, ?, ?, ?)"
		_, err := db.Exec(insertQuery, flavor.Flavor_name_th, flavor.Flavor_name_en, flavor.Flavor_price, flavor.Flavor_Stock)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data", "details": errors})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flavors created successfully"})
}

func GetFlavors(c *gin.Context, db *sql.DB) {
	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Query the database
	rows, err := db.Query("SELECT * FROM flavor")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
	defer rows.Close()

	var flavor models.Flavor
	for rows.Next() {
		err := rows.Scan(&flavor.Flavor_ID, &flavor.Flavor_name_th, &flavor.Flavor_name_en, &flavor.Flavor_price, &flavor.Flavor_Stock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		flavors = append(flavors, flavor)
	}

	c.JSON(http.StatusOK, flavors)
}

func GetFlavor(c *gin.Context, db *sql.DB) {
	flavorID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var flavor models.Flavor
	err := db.QueryRow("SELECT * FROM flavor WHERE Flavor_ID = ?", flavorID).Scan(&flavor.Flavor_ID, &flavor.Flavor_name_th, &flavor.Flavor_name_en, &flavor.Flavor_price, &flavor.Flavor_Stock)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}

	c.JSON(http.StatusOK, flavor)
}

func UpdateFlavor(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var flavor models.Flavor

	if err := c.ShouldBindJSON(&flavor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateQuery := "UPDATE flavor SET Flavor_name_th=?, Flavor_name_en=?, Flavor_price=?, Flavor_Stock=? WHERE Flavor_ID=?"
	_, err := db.Exec(updateQuery, flavor.Flavor_name_th, flavor.Flavor_name_en, flavor.Flavor_price, flavor.Flavor_Stock, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flavor updated successfully"})
}

func DeleteFlavor(c *gin.Context, db *sql.DB) {
	flavorID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	deleteQuery := "DELETE FROM flavor WHERE Flavor_ID = ?"
	_, err := db.Exec(deleteQuery, flavorID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flavor deleted successfully"})
}

