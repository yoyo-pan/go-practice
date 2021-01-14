package routers


import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"yoyoserver/routers/api/v1/users"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", defaultHandler)
	
	userV1 := r.Group("/users")
	{
		userV1.GET("/", users.GetUsers)
	}

	return r
}

func defaultHandler(c *gin.Context) {
	body, err := c.GetRawData()

	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println("Body %v", body)
	c.JSON(http.StatusOK, "success")
}