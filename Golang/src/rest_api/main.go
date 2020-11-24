package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"io/ioutil"
)

//HomePage test get
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

// PostHomePage test Post page
func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"message": string(value),
	})
}

// QueryStrings accepts values to be collected
func QueryStrings(c *gin.Context) {
	name := c.Query("name") // name = elliot
	age := c.Query("age")   // age = 24

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})

}

// PathParameters accepts params in url path
func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings)             // /query?name = elliot&age = 24
	r.GET("/path/:name/:age", PathParameters) // /path/elliot/24/
	r.Run()
}
