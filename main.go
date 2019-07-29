package main

import (
	"fmt"

	"github.com/dnivra26/things_list/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=thingslist dbname=thingslist password=password sslmode=disable")
	if err != nil {
		fmt.Println("Couldn't connect to DB", err)
	}
	defer db.Close()

	db.AutoMigrate(&models.List{})
	db.AutoMigrate(&models.Thing{}).AddForeignKey("list_id", "lists(id)", "RESTRICT", "RESTRICT")

	r := gin.Default()
	r.GET("/ping", ping)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
