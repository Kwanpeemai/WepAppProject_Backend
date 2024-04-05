package controller

import (
	"database/sql"
	"log"
	"net/http"
	// "time"
	"web-project/models"

	"github.com/gin-gonic/gin"
)

// var payments []models.Payment

func CreatePayment(c *gin.Context, db *sql.DB) {
	var payment models.Payment

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert into database
	insertQuery := "INSERT INTO payment (Payment_method, Payment_amount, Order_id, Payment_date, Payment_time) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(insertQuery, payment.Payment_method, payment.Payment_amount, payment.Order_id, payment.Payment_date, payment.Payment_time)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment created successfully"})
}

func GetPayments(c *gin.Context, db *sql.DB) {
	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Query the database
	rows, err := db.Query("SELECT * FROM payment")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(&payment.Payment_id, &payment.Payment_method, &payment.Payment_amount, &payment.Order_id, &payment.Payment_date, &payment.Payment_time)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		payments = append(payments, payment)
	}

	c.JSON(http.StatusOK, payments)
}

func GetPayment(c *gin.Context, db *sql.DB) {
	paymentID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var payment models.Payment
	err := db.QueryRow("SELECT * FROM payment WHERE Payment_id = ?", paymentID).Scan(&payment.Payment_id, &payment.Payment_method, &payment.Payment_amount, &payment.Order_id, &payment.Payment_date, &payment.Payment_time)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

func UpdatePayment(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update database
	updateQuery := "UPDATE payment SET Payment_method=?, Payment_amount=?, Order_id=?, Payment_date=?, Payment_time=? WHERE Payment_id=?"
	_, err := db.Exec(updateQuery, payment.Payment_method, payment.Payment_amount, payment.Order_id, payment.Payment_date, payment.Payment_time, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment updated successfully"})
}

func DeletePayment(c *gin.Context, db *sql.DB) {
	paymentID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	deleteQuery := "DELETE FROM payment WHERE Payment_id = ?"
	_, err := db.Exec(deleteQuery, paymentID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}
