package main

import (
	"net/http"

	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/shoppehub/shoppe"
)

//GetAllData all list
func GetAllData(c *gin.Context) {
	posts := []string{
		"Larry Ellison",
		"Carlos Slim Helu",
		"Mark Zuckerberg",
		"Amancio Ortega ",
		"Jeff Bezos",
		" Warren Buffet ",
		"Bill Gates",
		"selman tunç",
	}
	// Call the HTML method of the Context to render a template
	c.HTML(http.StatusOK, "index.html",
		pongo2.Context{
			"title": "hello pongo",
			"posts": posts,
		},
	)
}

func main() {
	r := shoppe.New()

	r.GET("home", GetAllData)

	shoppe.Start(r)
}
