package users

/*
 * Collection of all GET/SET apis and handlers for User model
 */

import (
	"net/http"

	"github.com/Vishal-Gupta19/go_gin_website/entity/common"
	"github.com/gin-gonic/gin"
)

// Mapping paths to api groups

// UserRegisterGroup : Personal user apis
func UserRegisterGroup(router *gin.RouterGroup) {
	router.POST("/", usersRegistration)
	// router.POST("/login", UsersLogin)
	// router.PUT("/", UserUpdate)
}

// UserProfileGroup : Other users apis
func UserProfileGroup(router *gin.RouterGroup) {
	router.GET("/:username", ProfileRetrieve)
	// router.POST("/:username/follow", ProfileFollow)
	// router.DELETE("/:username/follow", ProfileUnfollow)
}

// UsersRegistration : Handle for UsersRegistration
func usersRegistration(c *gin.Context) {
	// Get the user details from context and store in Validator model
	u := NewUserModelValidator()
	err := u.Validate(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	// Copy data in User model
	user := User{
		Username:  u.Username,
		Email:     u.Email,
		Password:  setPassword(u.Password),
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}

	err = SaveNewUserToDB(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}

	c.Set("my_user_model", user)
	serializer := UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

// ProfileRetrieve : Get user from db using username
func ProfileRetrieve(c *gin.Context) {
	username := c.Param("username")
	var user User
	user, err := GetUserProfile(&User{Username: username})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}
	c.Set("my_user_model", user)
	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
