package main

import (
	"groupie/routeur"
	"groupie/templates"
)

func main() {
	templates.InitTemplates()
	routeur.Initserv()
}