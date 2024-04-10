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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func main() {
	Db = SetupDB()

	r := gin.Default()

	r.Use(CORSMiddleware())

	//Size
	r.POST("/createSize", func(c *gin.Context) { controller.CreateSize(c, Db) })
	r.GET("/getSizes", func(c *gin.Context) { controller.GetSizes(c, Db) })
	r.GET("/getSize/:id", func(c *gin.Context) { controller.GetSize(c, Db) })
	r.PUT("/updateSize/:id", func(c *gin.Context) { controller.UpdateSize(c, Db) })
	r.DELETE("/deleteSize/:id", func(c *gin.Context) { controller.DeleteSize(c, Db)})

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

	//orderdetail_en
	r.POST("/createOrderDetail-en", func(c *gin.Context) { controller.CreateOrderDetail_en(c, Db) })
	r.GET("/getOrderDetails-en", func(c *gin.Context) { controller.GetOrderDetails_en(c, Db) })
	r.GET("/getOrderDetail-en/:id", func(c *gin.Context) { controller.GetOrderDetail_en(c, Db) })
	r.PUT("/updateOrderDetail-en/:id", func(c *gin.Context) { controller.UpdateOrderDetail_en(c, Db) })
	r.DELETE("/deleteOrderDetail-en/:id", func(c *gin.Context) { controller.DeleteOrderDetail_en(c, Db) })

	//orderdetail_th
	r.POST("/createOrderDetail-th", func(c *gin.Context) { controller.CreateOrderDetail_th(c, Db) })
	r.GET("/getOrderDetails-th", func(c *gin.Context) { controller.GetOrderDetails_th(c, Db) })
	r.GET("/getOrderDetail-th/:id", func(c *gin.Context) { controller.GetOrderDetail_th(c, Db) })
	r.PUT("/updateOrderDetail-th/:id", func(c *gin.Context) { controller.UpdateOrderDetail_th(c, Db) })
	r.DELETE("/deleteOrderDetail-th/:id", func(c *gin.Context) { controller.DeleteOrderDetail_th(c, Db) })

	r.Run(":8080")
}
