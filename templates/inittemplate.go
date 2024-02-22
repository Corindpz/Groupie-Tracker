package templates

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template


func InitTemplates() {
	// Parse the templates
	temp, err := template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Temp = temp
}
