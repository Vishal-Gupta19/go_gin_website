package main

import (
	"github.com/Vishal-Gupta19/go_gin_website/entity/database"
	"github.com/Vishal-Gupta19/go_gin_website/entity/users"
	wall "github.com/Vishal-Gupta19/go_gin_website/entity/wallpapers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// InitialMigration : Initialize the postgres db and Migrate User model
func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
}

func main() {
	// Initialize Postgres DB
	db := database.Init()
	// Migrate models to DB
	InitialMigration(db)

	server := gin.Default()

	// Creating api groups, starting with <domain>/api/. . .
	v1 := server.Group("/api")

	// User group apis, <domain>/api/user/. . .
	users.UserRegisterGroup(v1.Group("/user"))
	users.UserProfileGroup(v1.Group("/user"))

	// Wallpaper group apis, <domain>/api/wall/. . .
	wall.WallpaperGroup(v1.Group("/wall"))
	wall.WallpaperOtherGroup(v1.Group("/wall"))

	server.Run(":8080")
}
