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
	// router.GET("/", UserRetrieve)
	// router.PUT("/", UserUpdate)
}

// UserProfileGroup : Other users apis
func UserProfileGroup(router *gin.RouterGroup) {
	// router.GET("/:username", ProfileRetrieve)
	// router.POST("/:username/follow", ProfileFollow)
	// router.DELETE("/:username/follow", ProfileUnfollow)
}

// UsersRegistration : Handle for UsersRegistration
func usersRegistration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	err := userModelValidator.Validate(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	err = SaveNewUserToDB(userModelValidator.userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	}

	c.JSON(http.StatusCreated, gin.H{"user": userModelValidator.userModel})
}
