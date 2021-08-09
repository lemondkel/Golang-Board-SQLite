package main

import (
	"database/sql"
	"github.com/go-martini/martini"
)

func main() {
	db, _ := sql.Open("sqlite3", "./foo.db")
	m := martini.Classic()
	m.Map(db)
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
