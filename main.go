package main

import (
	"embed"
	"go-portal/routes"
)

//go:embed view/*
var f embed.FS

func main() {
	routes.F = f
	r := routes.InitRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
