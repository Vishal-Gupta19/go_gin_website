package users

import (
	"encoding/json"

	"github.com/Vishal-Gupta19/go_gin_website/entity/database"
	wall "github.com/Vishal-Gupta19/go_gin_website/entity/wallpapers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Add all db related operations for this model below

// User : Model for user credentials
type User struct {
	gorm.Model
	Username     string           `gorm:"column:username"`
	Email        string           `gorm:"column:email;unique_index"`
	Bio          string           `gorm:"column:bio;size:1024"`
	ProfilePic   *string          `gorm:"column:image"`
	PasswordHash string           `gorm:"column:password;not null"`
	FirstName    string           `gorm:"column:firstname"`
	LastName     string           `gorm:"column:lastname"`
	City         string           `gorm:"column:city"`
	Wallpapers   []wall.Wallpaper `gorm:"references:Username"` // User model has 'One to many' relation with Wallpaper model
}

// GetUsers : Handler to GET all users from DB
func GetUsers(g *gin.Context) {
	db := database.GetDB()
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(g.Writer).Encode(users)
}
