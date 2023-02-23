package levedb

import (
	"db/config"
	"encoding/json"
	"fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

type Info struct {
	Name     string `json:"name"`
	FullName string `json:"fullname"`
}

var db *leveldb.DB

func InitLeveDb() (*leveldb.DB, error) {
	config := config.GetConfig()
	db, err := leveldb.OpenFile(config.Sever.ServerLevelDB.PathFile, nil)
	if err != nil {
		log.Fatal("Yikes!")
	}
	defer db.Close()
	return db, err
}

func PutData() {

	byteArray, err := json.Marshal(Info{Name: "Nhat", FullName: "NTN"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
	err = db.Put([]byte("data_test"), byteArray, nil)
	// err = db.Put([]byte("fizz2"), []byte("buzz2"), nil)
	// err = db.Put([]byte("fizz3"), []byte("buzz3"), nil)
	// err = db.Delete([]byte("fizz"), nil)
	// err = db.Delete([]byte("fizz3"), nil)
}

func GetData() {
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		var temp Info
		err := json.Unmarshal([]byte(value), &temp)
		if err != nil {
			panic(err)
		}
		fmt.Println("key: ", key)
		fmt.Println("temp: ", temp)
	}

	// fmt.Println("\n")

	// for ok := iter.Seek([]byte("fizz2")); ok; ok = iter.Next() {
	// 	key := iter.Key()
	// 	value := iter.Value()
	// 	fmt.Printf("key: %s | value: %s\n", key, value)
	// }

	// fmt.Println("\n")

	// for ok := iter.First(); ok; ok = iter.Next() {
	// 	key := iter.Key()
	// 	value := iter.Value()
	// 	fmt.Printf("key: %s | value: %s\n", key, value)
	// }

	// iter.Release()
	// err = iter.Error()
}
