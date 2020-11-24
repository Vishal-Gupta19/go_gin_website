package users

import (
	wall "github.com/Vishal-Gupta19/go_gin_website/entity/wallpapers"
	"github.com/jinzhu/gorm"
)

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
