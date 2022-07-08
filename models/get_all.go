package models

import "github.com/Amad3eu/api-pstg-go/db"

func GetAll(id int64) (todos []Todo, err error){
	conn, err, := db.OpenConnection()

	if err != nil {
		return
	}
	 defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
	if err != nil {
		continue
	}

	todos = append(todos, todo)
	}

	return 

	 row := conn.QueryRow("SELECT * FROM todos WHERE id = $1", id)
	 err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	 return
	}
