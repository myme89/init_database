package postgresdb

import (
	sql "database/sql"
	"db/config"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once
var err error

func InitPostgresDb() (*sql.DB, error) {
	once.Do(func() {

		config := config.GetConfig()

		host := config.Sever.ServerPostgersSql.DBHost
		port := config.Sever.ServerPostgersSql.DBPort
		user := config.Sever.ServerPostgersSql.DBUserName
		password := config.Sever.ServerPostgersSql.DBPassword
		dbname := config.Sever.ServerPostgersSql.DBName

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err = sql.Open("postgres", psqlInfo)

	})
	return db, err
}
