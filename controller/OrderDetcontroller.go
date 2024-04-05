package controller

import (
	"database/sql"
	"log"
	"net/http"
	"web-project/models"

	"github.com/gin-gonic/gin"
)


func CreateOrderDetail(c *gin.Context, db *sql.DB) {
	var orderDetail models.Order_detail

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&orderDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertQuery := "INSERT INTO order_detail (Order_id, Size_name_en, Flavor_name_en, Topping_name_en, Sauce_name_en, Sum_Price) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(insertQuery, orderDetail.Order_id, orderDetail.Size_name_en, orderDetail.Flavor_name_en, orderDetail.Topping_name_en, orderDetail.Sauce_name_en, orderDetail.Sum_Price)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order detail created successfully"})
}

func GetOrderDetails(c *gin.Context, db *sql.DB) {
	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	rows, err := db.Query("SELECT * FROM order_detail")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
	defer rows.Close()

	var orderDetails []models.Order_detail
	for rows.Next() {
		var orderDetail models.Order_detail
		err := rows.Scan(&orderDetail.Order_id, &orderDetail.Size_name_en, &orderDetail.Flavor_name_en, &orderDetail.Topping_name_en, &orderDetail.Sauce_name_en, &orderDetail.Sum_Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		orderDetails = append(orderDetails, orderDetail)
	}

	c.JSON(http.StatusOK, orderDetails)
}

func GetOrderDetail(c *gin.Context, db *sql.DB) {
	detailID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var orderDetail models.Order_detail
	err := db.QueryRow("SELECT * FROM order_detail WHERE Order_id = ?", detailID).Scan(&orderDetail.Order_id, &orderDetail.Size_name_en, &orderDetail.Flavor_name_en, &orderDetail.Topping_name_en, &orderDetail.Sauce_name_en, &orderDetail.Sum_Price)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}

	c.JSON(http.StatusOK, orderDetail)
}

func UpdateOrderDetail(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var orderDetail models.Order_detail

	if err := c.ShouldBindJSON(&orderDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update database
	updateQuery := "UPDATE order_detail SET Size_ID=?, Flavor_ID=?, Topping_ID=?, Sauce_ID=?, Price=? WHERE Order_id=?"
	_, err := db.Exec(updateQuery, orderDetail.Size_name_en, orderDetail.Flavor_name_en, orderDetail.Topping_name_en, orderDetail.Sauce_name_en, orderDetail.Sum_Price, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order detail updated successfully"})
}

func DeleteOrderDetail(c *gin.Context, db *sql.DB) {
	detailID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	deleteQuery := "DELETE FROM order_detail WHERE Order_id = ?"
	_, err := db.Exec(deleteQuery, detailID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order detail deleted successfully"})
}
