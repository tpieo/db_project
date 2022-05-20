package main

import (
	"SZUCourse/global"
	"SZUCourse/views"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	dsn := "cj:666666@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	global.Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	router.GET("/", views.Index)
	router.GET("/query", views.Query)

	router.Run(":80")
}
