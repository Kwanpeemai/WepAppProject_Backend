package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"web-project/models"
)

var payments []models.Payment

// CreatePayment creates a new payment
func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var newPayment models.Payment
	newPayment.Payment_date = time.Now()
	newPayment.Payment_time = time.Now()
	json.NewDecoder(r.Body).Decode(&newPayment)
	payments = append(payments, newPayment)
	json.NewEncoder(w).Encode(newPayment)
}

// GetPayments returns all payments
func GetPayments(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(payments)
}

// GetPayment returns a single payment
func GetPayment(w http.ResponseWriter, r *http.Request) {
	paymentID := r.URL.Query().Get("id")
	for _, payment := range payments {
		if fmt.Sprint(payment.Payment_id) == paymentID {
			json.NewEncoder(w).Encode(payment)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Payment{})
}

// UpdatePayment updates an existing payment
func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	var updatedPayment models.Payment
	json.NewDecoder(r.Body).Decode(&updatedPayment)
	for i, payment := range payments {
		if payment.Payment_id == updatedPayment.Payment_id {
			payments[i] = updatedPayment
			json.NewEncoder(w).Encode(updatedPayment)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Payment{})
}

// DeletePayment deletes a payment
func DeletePayment(w http.ResponseWriter, r *http.Request) {
	paymentID := r.URL.Query().Get("id")
	for i, payment := range payments {
		if fmt.Sprint(payment.Payment_id) == paymentID {
			payments = append(payments[:i], payments[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(payments)
}
