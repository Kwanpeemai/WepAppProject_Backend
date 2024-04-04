package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-project/models"
)

var orders []models.Order

// CreateOrder creates a new order
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order
	json.NewDecoder(r.Body).Decode(&newOrder)
	orders = append(orders, newOrder)
	json.NewEncoder(w).Encode(newOrder)
}

// GetOrders returns all orders
func GetOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(orders)
}

// GetOrder returns a single order
func GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("id")
	for _, order := range orders {
		if fmt.Sprint(order.Order_ID) == orderID {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Order{})
}

// UpdateOrder updates an existing order
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var updatedOrder models.Order
	json.NewDecoder(r.Body).Decode(&updatedOrder)
	for i, order := range orders {
		if order.Order_ID == updatedOrder.Order_ID {
			orders[i] = updatedOrder
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Order{})
}

// DeleteOrder deletes an order
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("id")
	for i, order := range orders {
		if fmt.Sprint(order.Order_ID) == orderID {
			orders = append(orders[:i], orders[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(orders)
}
