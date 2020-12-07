package users

import (
	"github.com/Vishal-Gupta19/go_gin_website/entity/common"
	// wall "github.com/Vishal-Gupta19/go_gin_website/entity/wallpapers"
	"github.com/gin-gonic/gin"
)

// UserModelValidator : Struct to validate User Model
type UserModelValidator struct {
	Username   string  `form:"username" json:"username" binding:"min=8,max=255"`
	Email      string  `form:"email" json:"email"`
	Password   string  `form:"password" json:"password" binding:"max=255"`
	Bio        string  `form:"bio" json:"bio" binding:"max=1024"`
	ProfilePic *string `form:"image" json:"image" binding:"omitempty,url"`
	FirstName  string  `form:"firstname" json:"firstname" binding:"max=100"`
	LastName   string  `form:"lastname" json:"lastname" binding:"max=100"`
	City       string  `form:"city" json:"city" binding:"max=255"`
	// Wallpapers []wall.Wallpaper `form:"wallpapers" json:"wallpapers"`
}

// NewUserModelValidator : Constructor for UserModelValidator
func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	return userModelValidator
}

// Validate : This will validate the inputs using bindings
func (u *UserModelValidator) Validate(c *gin.Context) error {
	err := common.Bind(c, u)
	if err != nil {
		return err
	}

	return nil
}
