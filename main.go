package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		htmlContent := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Daucu Cloud</title>
        </head>
        <body>
            <h1>Daucu Cloud</h1>
            <p>This is a simple HTML page.</p>
			<p>https://daucu.com</p>
        </body>
        </html>
        `
		c.Data(200, "text/html; charset=utf-8", []byte(htmlContent))
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
