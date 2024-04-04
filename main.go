package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"web-project/controller"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func SetupDB() {
	Db, _ = sql.Open("mysql", "root:kwanpeemai0101@tcp(127.0.0.1:3306)/YogurtShop")

	fmt.Println(Db)
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}

func main() {
	SetupDB()
	// http.HandleFunc("/orders", controller.GetOrders)
	// http.HandleFunc("/order", controller.GetOrder)
	// http.HandleFunc("/createOrder", controller.CreateOrder)
	// http.HandleFunc("/updateOrder", controller.UpdateOrder)
	// http.HandleFunc("/deleteOrder", controller.DeleteOrder)

	http.HandleFunc("/sizes", controller.GetSizes)
	http.HandleFunc("/size", controller.GetSize)
	http.HandleFunc("/createSize", controller.CreateSize)
	http.HandleFunc("/updateSize", controller.UpdateSize)
	http.HandleFunc("/deleteSize", controller.DeleteSize)

	// http.HandleFunc("/flavors", controller.GetFlavors)
	// http.HandleFunc("/flavor", controller.GetFlavor)
	// http.HandleFunc("/createFlavor", controller.CreateFlavor)
	// http.HandleFunc("/updateFlavor", controller.UpdateFlavor)
	// http.HandleFunc("/deleteFlavor", controller.DeleteFlavor)

	// http.HandleFunc("/toppings", controller.GetToppings)
	// http.HandleFunc("/topping", controller.GetTopping)
	// http.HandleFunc("/createTopping", controller.CreateTopping)
	// http.HandleFunc("/updateTopping", controller.UpdateTopping)
	// http.HandleFunc("/deleteTopping", controller.DeleteTopping)

	// http.HandleFunc("/sauces", controller.GetSauces)
	// http.HandleFunc("/sauce", controller.GetSauce)
	// http.HandleFunc("/createSauce", controller.CreateSauce)
	// http.HandleFunc("/updateSauce", controller.UpdateSauce)
	// http.HandleFunc("/deleteSauce", controller.DeleteSauce)

	http.HandleFunc("/payments", controller.GetPayments)
	http.HandleFunc("/payment", controller.GetPayment)
	http.HandleFunc("/createPayment", controller.CreatePayment)
	http.HandleFunc("/updatePayment", controller.UpdatePayment)
	http.HandleFunc("/deletePayment", controller.DeletePayment)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
