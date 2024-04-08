
package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"web-project/controller"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/createSize", func(c *gin.Context) { controller.CreateSize(c, db) })
	r.GET("/getSizes", func(c *gin.Context) { controller.GetSizes(c, db) })
	r.GET("/getSize/:id", func(c *gin.Context) { controller.GetSize(c, db) })
	r.PUT("/updateSize/:id", func(c *gin.Context) { controller.UpdateSize(c, db) })
	r.DELETE("/deleteSize/:id", func(c *gin.Context) { controller.DeleteSize(c, db) })

	return r
}
