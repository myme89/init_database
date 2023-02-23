package main

import (
	"db/database/levedb"
	"db/database/mongodb"
	"db/database/mysqldb"
	"db/database/postgresdb"
	"fmt"
)

func main() {
	var sel int
	fmt.Println("/************************************************************************/")
	fmt.Println("Moi ban chon:")
	fmt.Println("\t 1: Leve DB")
	fmt.Println("\t 2: My Sql")
	fmt.Println("\t 3: Postgress DB")
	fmt.Println("\t 4: Mongo DB")
	fmt.Printf("Nhap lua chon: ")
	fmt.Scan(&sel)
	fmt.Println("/************************************************************************/")

	switch sel {
	case 1:
		fmt.Println()
		fmt.Println("/************************************************************************/")
		fmt.Println("Ban chon: leve DB")

		_, err := levedb.InitLeveDb()

		if err != nil {
			fmt.Println("Error connecting to database : error=% ", err)
		}
		fmt.Println("Successfully Leve DB connected!")

		fmt.Println("/************************************************************************/")

	case 2:
		fmt.Println()
		fmt.Println("/************************************************************************/")
		fmt.Println("Ban chon: My Sql DB")

		_, err := mysqldb.InitMySqlDb()

		if err != nil {
			fmt.Println("Error connecting to database : error=% ", err)
		}

		fmt.Println("Successfully MySql DB connected!")

		fmt.Println("/************************************************************************/")

	case 3:
		fmt.Println()
		fmt.Println("/************************************************************************/")
		fmt.Println("Ban chon: Postgres DB")

		db, err := postgresdb.InitPostgresDb()

		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully Postgres DB connected!")

		fmt.Println("/************************************************************************/")

	case 4:
		fmt.Println()
		fmt.Println("/************************************************************************/")
		fmt.Println("Ban chon: Mongo DB")
		mongodb.Init()
		fmt.Println("/************************************************************************/")
	default:
		fmt.Println("Hello")
		fmt.Println("/************************************************************************/")

	}
}
