package db

import (
	"database/sql"
	"fmt"

	"github.com/Amad3eu/api-pstg-go/configs"
	
	_ "github.com/lib/pq"
)
func openConnection(*sql.DB, errot) {
	conf := configs.GetDB()
}

sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
conf.host, conf.port, conf.user, conf.pass, conf.database)

conn, err := sql.Open("postgres", sc)
if err != nil {
	panic(err)
}

err = conn.Ping()

return conn, err
