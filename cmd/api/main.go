package main

import (
	"github.com/dositadi/groupie-tracker/cmd/api/app"
)

func main() {
	app := &app.App{}
	app.Run()

	/* db := new(pgx.Conn)

	m := models.New(db)
	m.UserModel.InsertUser(data.User{}) */
}
