package users

/*
 * Collection of all GET/SET apis and handlers for User model
 */

import (
	"github.com/gin-gonic/gin"
)

// Mapping paths to api groups

// UserRegisterGroup : Personal user apis
func UserRegisterGroup(router *gin.RouterGroup) {
	// router.POST("/", UsersRegistration)
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
