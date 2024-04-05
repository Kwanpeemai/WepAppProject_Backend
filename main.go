package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"web-project/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func SetupDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "kwanpeemai0101", "127.0.0.1", "3306", "YogurtShop")
	var err error
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to database!")

	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

	return Db
}

func main() {
	Db = SetupDB()

	r := gin.Default()
	//Size
	r.POST("/createSize", func(c *gin.Context) { controller.CreateSize(c, Db) })
	r.GET("/getSizes", func(c *gin.Context) { controller.GetSizes(c, Db) })
	r.GET("/getSize/:id", func(c *gin.Context) { controller.GetSize(c, Db) })
	r.PUT("/updateSize/:id", func(c *gin.Context) { controller.UpdateSize(c, Db) })
	r.DELETE("/deleteSize/:id", func(c *gin.Context) { controller.DeleteSize(c, Db) })

	//Flavor
	r.POST("/createFlavor", func(c *gin.Context) { controller.CreateFlavor(c, Db) })
	r.GET("/getFlavors", func(c *gin.Context) { controller.GetFlavors(c, Db) })
	r.GET("/getFlavor/:id", func(c *gin.Context) { controller.GetFlavor(c, Db) })
	r.PUT("/updateFlavor/:id", func(c *gin.Context) { controller.UpdateFlavor(c, Db) })
	r.DELETE("/deleteFlavor/:id", func(c *gin.Context) { controller.DeleteFlavor(c, Db) })

	//Topping
	r.POST("/createTopping", func(c *gin.Context) { controller.CreateTopping(c, Db) })
	r.GET("/getToppings", func(c *gin.Context) { controller.GetToppings(c, Db) })
	r.GET("/getTopping/:id", func(c *gin.Context) { controller.GetTopping(c, Db) })
	r.PUT("/updateTopping/:id", func(c *gin.Context) { controller.UpdateTopping(c, Db) })
	r.DELETE("/deleteTopping/:id", func(c *gin.Context) { controller.DeleteTopping(c, Db) })

	//Sauce
	r.POST("/createSauce", func(c *gin.Context) { controller.CreateSauce(c, Db) })
	r.GET("/getSauces", func(c *gin.Context) { controller.GetSauces(c, Db) })
	r.GET("/getSauce/:id", func(c *gin.Context) { controller.GetSauce(c, Db) })
	r.PUT("/updateSauce/:id", func(c *gin.Context) { controller.UpdateSauce(c, Db) })
	r.DELETE("/deleteSauce/:id", func(c *gin.Context) { controller.DeleteSauce(c, Db) })

	//orderdetail
	r.POST("/createOrderDetail", func(c *gin.Context) { controller.CreateOrderDetail(c, Db) })
	r.GET("/getOrderDetails", func(c *gin.Context) { controller.GetOrderDetails(c, Db) })
	r.GET("/getOrderDetail/:id", func(c *gin.Context) { controller.GetOrderDetail(c, Db) })
	r.PUT("/updateOrderDetail/:id", func(c *gin.Context) { controller.UpdateOrderDetail(c, Db) })
	r.DELETE("/deleteOrderDetail/:id", func(c *gin.Context) { controller.DeleteOrderDetail(c, Db) })

	//payment
	r.POST("/createPayment", func(c *gin.Context) { controller.CreatePayment(c, Db) })
	r.GET("/getPayments", func(c *gin.Context) { controller.GetPayments(c, Db) })
	r.GET("/getPayment/:id", func(c *gin.Context) { controller.GetPayment(c, Db) })
	r.PUT("/updatePayment/:id", func(c *gin.Context) { controller.UpdatePayment(c, Db) })
	r.DELETE("/deletePayment/:id", func(c *gin.Context) { controller.DeletePayment(c, Db) })

	r.Run(":8080")
}
