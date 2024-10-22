package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...

func main() {
	// log.Println("Starting Railroad-Central-Command-Server, have fun training")
	// server.ListenAndServeServer(&config.Config{})
	db, err := sql.Open("mysql", "railroad_central_command_admin:fIqaSaspumusubAphu4e@tcp(zduvewarehouse.local:3307)/")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Println("We pinged the database")
	}

	res, err := db.Query("CREATE TALE * FROM USERS")
	if err != nil {
		panic(err)
	} else {
		log.Println("We got results")
	}

	log.Println(res.Columns())
}
