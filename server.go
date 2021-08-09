package main

import (
	"database/sql"
	"github.com/go-martini/martini"
	"net/http"
)

func main() {
	db, _ := sql.Open("sqlite3", "./foo.db")
	m := martini.Classic()
	m.Map(db)
	m.Get("/", func() string {
		return "index"
	})
	assets := martini.Static("assets", martini.StaticOptions{Fallback: "/error.html", Exclude: "/api/v"})
	m.Use(martini.Static("assets"))
	m.NotFound(assets, http.NotFound)
	m.Run()
}
