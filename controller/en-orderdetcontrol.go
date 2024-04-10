package controller

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"web-project/models"

	"github.com/gin-gonic/gin"
)
//version eng
func CreateOrderDetail_en(c *gin.Context, db *sql.DB) {
	var orderDetail_en models.Order_detail_en

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if err := c.ShouldBindJSON(&orderDetail_en); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แปลง array ของ topping เป็นสตริงที่แยกด้วย comma
	toppings := strings.Join(orderDetail_en.Topping_name_en, "")
	
	insertQuery := "INSERT INTO order_detail (Order_id, Size_name_en, Flavor_name_en, Topping_name_en, Sauce_name_en, sum_Price) VALUES (?, ?, ?, ?, ?, ?)"
result, err := db.Exec(insertQuery, orderDetail_en.Order_id, orderDetail_en.Size_name_en, orderDetail_en.Flavor_name_en, toppings, orderDetail_en.Sauce_name_en, orderDetail_en.Sum_Price)
if err != nil {
    log.Printf("Error executing query: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
    return
}

lastInsertId, err := result.LastInsertId()
if err != nil {
    log.Printf("Error getting last insert ID: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting last insert ID"})
    return
}

c.JSON(http.StatusOK, gin.H{
    "Order_id": lastInsertId,
    "message":  "Data inserted successfully",
})



	// ลดสต็อกของไซส์
	_, err = db.Exec("UPDATE size SET Size_Stock = Size_Stock - 1 WHERE Size_name_en = ?", orderDetail_en.Size_name_en)
	if err != nil {
		log.Printf("Error updating size stock: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating size stock"})
		return
	}

	// ลดสต็อกของรส
	_, err = db.Exec("UPDATE flavor SET Flavor_Stock = Flavor_Stock - 1 WHERE Flavor_name_en = ?", orderDetail_en.Flavor_name_en)
	if err != nil {
		log.Printf("Error updating flavor stock: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating flavor stock"})
		return
	}

	// ลดสต็อกของท็อปปิ้ง
	for _, t := range orderDetail_en.Topping_name_en {
		_, err = db.Exec("UPDATE topping SET Topping_Stock = Topping_Stock - 1 WHERE Topping_name_en = ?", t)
		if err != nil {
			log.Printf("Error updating topping stock: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating topping stock"})
			return
		}
	}

	// ลดสต็อกของซอส
	_, err = db.Exec("UPDATE sauce SET Sauce_Stock = Sauce_Stock - 1 WHERE Sauce_name_en = ?", orderDetail_en.Sauce_name_en)
	if err != nil {
		log.Printf("Error updating sauce stock: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating sauce stock"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order detail created successfully"})
}


// func GetOrderDetail_en(c *gin.Context, db *sql.DB) {
// 	detailID := c.Param("id")

// 	if db == nil {
// 		log.Fatalf("DB connection is nil")
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 		return
// 	}

// 	var orderDetail models.Order_detail_en
// 	// var toppings string
// 	toppings := strings.Join(orderDetail.Topping_name_en, ",")

// 	err := db.QueryRow("SELECT Size_id,Size_name_en, Flavor_name_en, Topping_name_en, Sauce_name_en FROM order_detail WHERE Order_id = ?", detailID).Scan(&orderDetail.Size_name_en, &orderDetail.Flavor_name_en, &toppings, &orderDetail.Sauce_name_en)
// 	if err != nil {
// 		log.Printf("Error querying data: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
// 		return
// 	}	

// 	c.JSON(http.StatusOK, orderDetail)
// }

func GetOrderDetail_en(c *gin.Context, db *sql.DB) {
	detailID := c.Param("id")

	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var orderDetail models.Order_detail_en
	var toppings string
	
	err := db.QueryRow("SELECT Order_id, Size_name_en, Flavor_name_en, Topping_name_en, Sauce_name_en FROM order_detail WHERE Order_id = ?", detailID).Scan(&orderDetail.Order_id, &orderDetail.Size_name_en, &orderDetail.Flavor_name_en, &toppings, &orderDetail.Sauce_name_en)
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}	

	c.JSON(http.StatusOK, gin.H{
		"Order_id":        orderDetail.Order_id,
		"Size_name_en":    orderDetail.Size_name_en,
		"Flavor_name_en":  orderDetail.Flavor_name_en,
		"Topping_name_en": strings.Split(toppings, ","),
		"Sauce_name_en":   orderDetail.Sauce_name_en,
	})
}


func GetOrderDetails_en(c *gin.Context, db *sql.DB) {
	if db == nil {
		log.Fatalf("DB connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	rows, err := db.Query("SELECT Order_id, Size_name_en, Flavor_name_en, Topping_name_en, Sauce_name_en, sum_Price FROM order_detail")
	if err != nil {
		log.Printf("Error querying data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return
	}
	defer rows.Close()

	var orderDetails []models.Order_detail_en
	for rows.Next() {
		var orderDetail models.Order_detail_en
		var toppings string
		err := rows.Scan(&orderDetail.Order_id, &orderDetail.Size_name_en, &orderDetail.Flavor_name_en, &toppings, &orderDetail.Sauce_name_en, &orderDetail.Sum_Price)
		if err != nil {
			log.Printf("Error scanning data: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		orderDetail.Topping_name_en = strings.Split(toppings, ",")
		orderDetails = append(orderDetails, orderDetail)
	}

	c.JSON(http.StatusOK, orderDetails)
}


func UpdateOrderDetail_en(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var orderDetail models.Order_detail_en

	if err := c.ShouldBindJSON(&orderDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update database
	updateQuery := "UPDATE order_detail SET Size_name_en=?, Flavor_name_en=?, Topping_name_en=?, Sauce_name_en=?, Sum_Price=? WHERE Order_id=?"
	_, err := db.Exec(updateQuery, orderDetail.Size_name_en, orderDetail.Flavor_name_en, strings.Join(orderDetail.Topping_name_en, ","), orderDetail.Sauce_name_en, orderDetail.Sum_Price, id)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order detail updated successfully"})
}

func DeleteOrderDetail_en(c *gin.Context, db *sql.DB) {
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

