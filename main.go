package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Defaults()

    r.GET("/", func(c *gin.Context) {
        // Retrieve client information
        clientIP := c.ClientIP()
        userAgent := c.Request.UserAgent()
        headers := c.Request.Header

        // Generate HTML content
        htmlContent := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Daucu Cloud</title>
            <style>
                body {
                    font-family: 'Arial', sans-serif;
                    background-color: #f4f4f9;
                    margin: 0;
                    padding: 0;
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                    justify-content: center;
                    min-height: 100vh;
                    color: #333;
                }
                h1 {
                    color: #4CAF50;
                    font-size: 3em;
                    margin-bottom: 20px;
                }
                p {
                    font-size: 1.2em;
                    margin: 10px 0;
                }
                a {
                    color: #4CAF50;
                    text-decoration: none;
                    font-weight: bold;
                    font-size: 1.2em;
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
                    max-width: 800px;
                    width: 90%;
                    margin: 20px;
                    overflow: auto;
                }
                .info {
                    text-align: left;
                    margin-top: 20px;
                    font-size: 1em;
                    word-break: break-word;
                }
                .info p {
                    margin: 5px 0;
                }
                ul {
                    padding-left: 20px;
                }
                li {
                    margin-bottom: 5px;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Daucu Cloud</h1>
                <p><a href="https://daucu.com" target="_blank">https://daucu.com</a></p>
                <div class="info">
                    <p><strong>Client IP:</strong> ` + clientIP + `</p>
                    <p><strong>User Agent:</strong> ` + userAgent + `</p>
                    <p><strong>Headers:</strong></p>
                    <ul>
        `
        for key, values := range headers {
            for _, value := range values {
                htmlContent += `<li><strong>` + key + `:</strong> ` + value + `</li>`
            }
        }
        htmlContent += `
                    </ul>
                </div>
            </div>
        </body>
        </html>
        `
        c.Data(200, "text/html; charset=utf-8", []byte(htmlContent))
    })

    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
