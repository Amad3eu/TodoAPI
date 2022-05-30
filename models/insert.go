package models

import (
	"github.com/Amad3eu/api-pstg-go/db"
)

func Insert(todo Todo)(id int64, err error){
	conn, err := db.OpenConnection()
	
	return
}

defer conn.Close()

sql := "INSERT INTO todos (title, description, done) values ($1, $2, $3) RETURNING id"
err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)
return
