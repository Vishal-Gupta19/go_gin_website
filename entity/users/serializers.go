package users

import (
	// wall "github.com/Vishal-Gupta19/go_gin_website/entity/wallpapers"
	"github.com/gin-gonic/gin"
)

// UserResponse : Model for serializing the data in JSON format
type UserResponse struct {
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Bio        string  `json:"bio"`
	ProfilePic *string `json:"image"`
	FirstName  string  `json:"firstname"`
	LastName   string  `json:"lastname"`
	City       string  `json:"city"`
	// Wallpapers []wall.Wallpaper `json:"wallpaper"`
}

// UserSerializer : Wrapper over gin context
type UserSerializer struct {
	c *gin.Context
}

// Response : After filling the data, send the JSON model back
func (s *UserSerializer) Response() UserResponse {
	myUserModel := s.c.MustGet("my_user_model").(User)
	user := UserResponse{
		Username:   myUserModel.Username,
		Email:      myUserModel.Email,
		FirstName:  myUserModel.FirstName,
		LastName:   myUserModel.LastName,
		Password:   myUserModel.Password,
		Bio:        myUserModel.Bio,
		ProfilePic: myUserModel.ProfilePic,
		City:       myUserModel.City,
		// Wallpapers: myUserModel.Wallpapers,
	}
	return user
}
