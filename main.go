package main

import (
	"fmt"
	"net/http"

	"github.com/dnivra26/things_list/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=thingslist dbname=thingslist password=password sslmode=disable")
	if err != nil {
		fmt.Println("Couldn't connect to DB", err)
	}
	defer db.Close()

	db.AutoMigrate(&models.List{})
	db.AutoMigrate(&models.Thing{}).AddForeignKey("list_id", "lists(id)", "RESTRICT", "RESTRICT")

	r := gin.Default()
	r.GET("/ping", ping)
	v1Lists := r.Group("api/v1/lists")
	{
		v1Lists.POST("/", createList)
	}

	v1Things := r.Group("api/v1/things")
	{
		v1Things.POST("/", createThing)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}

func createThing(c *gin.Context) {
	var thing models.Thing
	if err := c.ShouldBind(&thing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&thing)
	c.JSON(http.StatusCreated, gin.H{"message": "Thing created successfully"})
}
func createList(c *gin.Context) {
	var list models.List
	if err := c.ShouldBind(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&list)
	c.JSON(http.StatusCreated, gin.H{"message": "List created successfully"})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
