package main

import (
	"bytes"
	"fmt"
	"io"
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
		v1.POST("/test", testEndpoint)
		v1.POST("/test2", testEndpoint2)
		v1.POST("/test3", testEndpoint3)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		// v2.POST("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		v2.POST("/test", testEndpoint)
	}

	router.Run(":8080")
}

func submitEndpoint(c *gin.Context) {
	body := c.Request.Body
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	defer body.Close()
	x, _ := ioutil.ReadAll(body)
	fmt.Printf("X: %s \n", string(x))

	y, _ := ioutil.ReadAll(body)
	fmt.Printf("Y: %s \n", string(y))

	// c.Next()
	// if c.Bind(&test) == nil {
	// 	c.JSON(http.StatusOK, gin.H{"status": test})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	// }

}

// Output
// X: aaaaaaaa
// Y:

func testEndpoint(c *gin.Context) {
	body := c.Request.Body
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	defer body.Close()

	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(body)
	}

	//body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	fmt.Printf("X: %s \n", string(bodyBytes))

	y, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("Y: %s \n", string(y))
}

// Output
// X: aaaaaaaa
// Y: aaaaaaaa

func testEndpoint2(c *gin.Context) {
	body := c.Request.Body
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	defer body.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(body, &buf)

	y, _ := ioutil.ReadAll(tee)
	fmt.Printf("Y: %s \n", string(y))

	x, _ := ioutil.ReadAll(&buf)
	fmt.Printf("X: %s \n", string(x))

	z, _ := ioutil.ReadAll(body)
	fmt.Printf("Z: %s \n", string(z))
}

// Output
// Y: aaaaaaaa
// X: aaaaaaaa
// Z:

func testEndpoint3(c *gin.Context) {
	body := c.Request.Body
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	defer body.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(body, &buf)

	x, _ := ioutil.ReadAll(&buf)
	fmt.Printf("X: %s \n", string(x))

	y, _ := ioutil.ReadAll(tee)
	fmt.Printf("Y: %s \n", string(y))

	z, _ := ioutil.ReadAll(body)
	fmt.Printf("Z: %s \n", string(z))
}

//  Output
// X:
// Y: aaaaaaaa
// Z:
