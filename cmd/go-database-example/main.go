package main

import (
	"log"
	"os"
)

var (
	persistence persistence.Persistence
)

func init() {
	dbConn := os.Getenv("DB_CONN")
	persistence, err := NewMysqlPersistence(dbConn)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

}
