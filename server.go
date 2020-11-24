package main

import (
	"fmt"

	"github.com/Vishal-Gupta19/go_gin_website/entity"
	"github.com/Vishal-Gupta19/go_gin_website/entity/users"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// InitialMigration : Initialize the postgres db and Migrate User model
func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
}

func main() {

	// Initialize Postgres DB
	db := entity.Init()
	// Migrate models to DB
	InitialMigration(db)

	server := gin.Default()

	server.GET("/", HomePage)

	// server.GET("/wall", db.GetWallpapers)
	// server.POST("/wall", db.GetWallpapers)

	server.Run(":8080")
}

// HomePage : Dummy knob for '/'
func HomePage(g *gin.Context) {
	fmt.Fprintln(g.Writer, "Welcome!!")
}
