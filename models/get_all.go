package models
func GetAll(id int64) (todos []Todo, err error){
	conn, err, := db.OpenConnection()

	if err != nil {
		return
	}
	 defer conn.Close()


	 row := conn.QueryRow("SELECT * FROM todos WHERE id = $1", id)
	 err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	 return
	}