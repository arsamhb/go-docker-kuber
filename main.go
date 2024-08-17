package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"message": "SUCCESS",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// WHEN YOU WRITE YOUR DOCKER RUN YOU HAVE TO TAKE A BUILD FROM IT
// USING docker build -t my-go-server .
// THEN YOU'LL BE ABLE TO USE THE IMAGE YOU HAVE
// docker run -d -p 8080:8080 --name go-server my-go-server
// NOW YOUR CONTAINER IS RUNNING
