package main

import (
	"time"
)

type BoardModel struct {
	uid        int64     `form:"uid" db:"uid"`
	username   string    `form:"username" db:"username"`
	departname string    `form:"departname" db:"departname"`
	created    time.Time `form:"-" db:"created"`
}

func (u *BoardModel) GetById(id interface{}) error {
	err := Db.SelectOne(u, "SELECT * FROM boards WHERE uid = $1", id)
	if err != nil {
		return err
	}

	return nil
}
