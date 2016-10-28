package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		// v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		// v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		// v2.POST("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		// v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

func submitEndpoint(c *gin.Context) {
	body := c.Request.Body
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	defer body.Close()
	x, _ := ioutil.ReadAll(body)
	fmt.Printf("%s \n", string(x))

	// c.Next()
	// if c.Bind(&test) == nil {
	// 	c.JSON(http.StatusOK, gin.H{"status": test})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	// }

}
