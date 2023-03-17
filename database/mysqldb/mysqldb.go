package mysqldb

import (
	"database/sql"
	"db/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sql.DB

func InitMySqlDb() (*gorm.DB, error) {
	fmt.Println("Go MySQL Tutorial")

	config := config.GetConfig()

	// dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME

	dsn := config.Sever.ServerMySql.DBUserName + ":" + config.Sever.ServerMySql.DBPassword + "@tcp" +
		"(" + config.Sever.ServerMySql.DBHost + ":" + config.Sever.ServerMySql.DBPort + ")/" + config.Sever.ServerMySql.DBName

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
