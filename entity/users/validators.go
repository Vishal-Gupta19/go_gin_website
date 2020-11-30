package users

import (
	"github.com/Vishal-Gupta19/go_gin_website/entity/common"
	"github.com/gin-gonic/gin"
)

// UserModelValidator : Struct to validate User Model
type UserModelValidator struct {
	User struct {
		Username   string `form:"username" json:"username" binding:"min=4,max=255"`
		Email      string `form:"email" json:"email" binding:"email"`
		Password   string `form:"password" json:"password" binding:"min=8,max=255"`
		Bio        string `form:"bio" json:"bio" binding:"max=1024"`
		ProfilePic string `form:"image" json:"image" binding:"omitempty,url"`
		FirstName  string `form:"firstname" json:"firstname" binding:"max=100"`
		LastName   string `form:"lastname" json:"lastname" binding:"max=100"`
		City       string `form:"city" json:"city" binding:"max=255"`
	} `json:"user"`
	userModel User `json:"-"`
}

// NewUserModelValidator : Constructor for UserModelValidator
func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	return userModelValidator
}

// Validate : This will validate the inputs using binding
func (u *UserModelValidator) Validate(c *gin.Context) error {
	err := common.Bind(c, u)
	if err != nil {
		return err
	}
	u.userModel.Username = u.User.Username
	u.userModel.Email = u.User.Email
	// u.userModel.Bio = u.User.Bio
	u.userModel.Password = u.User.Password
	u.userModel.FirstName = u.User.FirstName
	u.userModel.LastName = u.User.LastName

	// if self.User.Password != common.NBRandomPassword {
	// 	self.userModel.setPassword(self.User.Password)
	// }
	// if self.User.Image != "" {
	// 	self.userModel.Image = &self.User.Image
	// }
	return nil
}