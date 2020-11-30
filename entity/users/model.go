package users

import (
	"encoding/json"

	"github.com/Vishal-Gupta19/go_gin_website/entity/database"
	wall "github.com/Vishal-Gupta19/go_gin_website/entity/wallpapers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// User : Model for user credentials
type User struct {
	gorm.Model
	Username   string           `gorm:"column:username"`
	Email      string           `gorm:"column:email;unique"`
	Bio        string           `gorm:"column:bio;size:1024"`
	ProfilePic *string          `gorm:"column:image"`
	Password   string           `gorm:"column:password;not null"`
	FirstName  string           `gorm:"column:firstname"`
	LastName   string           `gorm:"column:lastname"`
	City       string           `gorm:"column:city"`
	Wallpapers []wall.Wallpaper `gorm:"references:Username"` // User model has 'One to many' relation with Wallpaper model
}

// Add all db related operations for this model below

// GetUsers : Handler to GET all users from DB
func GetUsers(g *gin.Context) {
	db := database.GetDB()
	// defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(g.Writer).Encode(users)
}

// SaveNewUserToDB : Save new user information to database
func SaveNewUserToDB(data interface{}) error {
	db := database.GetDB()
	err := db.Save(data).Error
	return err
}
