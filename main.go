package main

import (
	"net/http"

	r "github.com/marceloagmelo/go-teste-conexao/routers"
	"github.com/marceloagmelo/pongor-echo"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	e := r.App
	e.Static("/static", "./static")

	p := pongor.GetRenderer()
	p.Directory = "views"

	e.Renderer = p

	e.Start(":8080")
}
