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
            <style>
                body {
                    font-family: Arial, sans-serif;
                    background-color: #f4f4f9;
                    margin: 0;
                    padding: 0;
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                    justify-content: center;
                    height: 100vh;
                    color: #333;
                }
                h1 {
                    color: #4CAF50;
                }
                p {
                    font-size: 1.2em;
                }
                a {
                    color: #4CAF50;
                    text-decoration: none;
                    font-weight: bold;
                }
                a:hover {
                    text-decoration: underline;
                }
                .container {
                    text-align: center;
                    padding: 20px;
                    background-color: #fff;
                    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
                    border-radius: 8px;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Daucu Cloud</h1>
                <p><a href="https://daucu.com" target="_blank">https://daucu.com</a></p>
				<p>This is a simple http server</p>
            </div>
        </body>
        </html>
        `
        c.Data(200, "text/html; charset=utf-8", []byte(htmlContent))
    })

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
