package users

/*
 * Collection of all GET/SET handlers for User model
 */

import (
	"encoding/json"

	"github.com/Vishal-Gupta19/go_gin_website/entity"
	"github.com/gin-gonic/gin"
)

// GetUsers : Handler to GET all users from DB
func GetUsers(g *gin.Context) {
	db := entity.GetDB()
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(g.Writer).Encode(users)
}
