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
	Db, err := sql.Open("mysql", dsn)
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
	SetupDB()

	r := gin.Default()

	r.POST("/createSize", func(c *gin.Context) { controller.CreateSize(c, Db) })
	r.GET("/getSizes", func(c *gin.Context) { controller.GetSizes(c, Db) })

	r.Run(":8080")
}
