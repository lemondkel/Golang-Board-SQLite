package main

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

var Db *gorp.DbMap

func main() {
	// Settings
	db, _ := sql.Open("sqlite3", "./foo.db")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(BoardModel{}, "boards").SetKeys(true, "uid")

	m := martini.Classic()
	m.Map(db)
	assets := martini.Static("assets", martini.StaticOptions{Fallback: "/error.html", Exclude: "/api/v"})
	m.Use(martini.Static("assets"))
	m.NotFound(assets, http.NotFound)

	m.Get("/", func() string {
		return "index"
	})

	m.Get("/insert", func() string {
		board := BoardModel{1, "testuser", "password", time.Now()}
		err := dbmap.Insert(&board)
		if err != nil {
			log.Fatalln("Could not insert test user", err)
		}
		fmt.Println(err)
		m.Map(db)
		return "true"
	})

	// Run
	m.Run()
}
