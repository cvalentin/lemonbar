package main

import (
  "log"
  "fmt"
	"database/sql"
  _ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(rows)
}
