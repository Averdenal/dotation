package Controller

import (
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	p := Page{"Titre de ma page", []string{"Item 1", "Item 2", "Item 3"}}

	t := template.New("home page")

	t = template.Must(t.ParseFiles("./Template/layout.tmpl", "./Template/content.tmpl"))

	err := t.ExecuteTemplate(c.Writer, "layout", p)
	if err != nil {
		log.Fatalf("Template execution: %s", err)
	}
}

type Page struct {
	Title    string
	Articles []string
}
