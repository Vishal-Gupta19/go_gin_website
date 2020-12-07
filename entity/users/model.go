package users

import (
	"github.com/Vishal-Gupta19/go_gin_website/entity/database"
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
)

// User : Model for user credentials
type User struct {
	gorm.Model
	Username   string  `gorm:"column:username;primary_key"`
	Email      string  `gorm:"column:email;unique"`
	Bio        string  `gorm:"column:bio;size:1024"`
	ProfilePic *string `gorm:"column:image"`
	Password   string  `gorm:"column:password;not null"`
	FirstName  string  `gorm:"column:firstname"`
	LastName   string  `gorm:"column:lastname"`
	City       string  `gorm:"column:city"`
	// Wallpapers []wall.Wallpaper `gorm:"references:Username"` // User model has 'One to many' relation with Wallpaper model
}

// Other utility functions

// SetPassword : Set password hash
func setPassword(password string) string {
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(passwordHash)
}

// Add all db related operations for this model below

// GetUserProfile : Handler to GET all users from DB
func GetUserProfile(data interface{}) (User, error) {
	db := database.GetDB()
	var user User
	err := db.Find(&user, data).Error
	return user, err
}

// SaveNewUserToDB : Save new user information to database
func SaveNewUserToDB(data interface{}) error {
	db := database.GetDB()
	err := db.Create(data).Error
	return err
}
