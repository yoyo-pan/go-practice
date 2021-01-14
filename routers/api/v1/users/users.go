package users

import (
	"net/http"

	"yoyoserver/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	response := models.User{
		Username: "Yoyo",
		Email: "yoyo@yoyo.com",
	}

	c.JSON(http.StatusOK, response)
}