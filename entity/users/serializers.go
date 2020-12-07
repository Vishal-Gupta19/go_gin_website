package users

import "github.com/gin-gonic/gin"

// UserResponse : Model for serializing the data in JSON format
type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//Bio        string `form:"bio" json:"bio" binding:"max=1024"`
	//ProfilePic string `form:"image" json:"image" binding:"omitempty,url"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	//City       string `form:"city" json:"city" binding:"max=255"`
}

// UserSerializer : Wrapper over gin context
type UserSerializer struct {
	c *gin.Context
}

// Response : After filling the data, send the JSON model back
func (s *UserSerializer) Response() UserResponse {
	myUserModel := s.c.MustGet("my_user_model").(User)
	user := UserResponse{
		Username:  myUserModel.Username,
		Email:     myUserModel.Email,
		FirstName: myUserModel.FirstName,
		LastName:  myUserModel.LastName,
		Password:  myUserModel.Password,
	}
	return user
}
