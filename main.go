package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

func main() {
	const tpl = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta http-equiv="x-ua-compatible" content="ie=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />

    <title>f5 supportid checker</title>
  </head>

  <body>
<p>This is the Support ID Checker, please fill the supportid in the form</p>
<form action="/">
  <label for="supportid">SupportID:</label><br>
  <input type="text" name="supportid" ><br>
  <input type="submit" value="Submit">
</form> 
  </body>
</html>`
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)
	r := gin.Default()
	r.SetHTMLTemplate(t)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/temp", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"webpage",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
