package main

import (
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func main() {
	connectDb()
	defer clossDb()
	save()
	// selectByKey()

	// n := 100
	// bulbs := make([]bool, n)
	// for i := 1; i <= n; i++ {
	// 	for j := i - 1; j < n; j += i {
	// 		bulbs[j] = !bulbs[j]
	// 	}
	// }
	// for i, bulb := range bulbs {
	// 	if bulb {
	// 		fmt.Printf("燈泡 %d 是開的\n", i+1)
	// 	} else {
	// 		fmt.Printf("燈泡 %d 是關的\n", i+1)
	// 	}
	// }
}

var db *sqlx.DB

func connectDb() {
	// DSN or postgres://postgres:ppassword@localhost:5432/postgres?sslmode=disable
	dbx, err := sqlx.Connect("pgx", "user=postgres password=ppassword host=localhost port=5432 database=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}

	db = dbx
}

func clossDb() {
	db.Close()
}

func save() {
	insertSql := `INSERT INTO backmessage(key, value) VALUES($1, $2)`

	tx, _ := db.Begin()
	t := test{Key: "aaaaasd", Value: "asdafa"}

	_, err := tx.Exec(insertSql, t.Key, t.Value)

	if err != nil {
		tx.Rollback()
		log.Err(err).Msg("BackMessage Insert ERROR")
	}
	tx.Commit()

}

func selectByKey() {
	stmt, err := db.PrepareNamed(`SELECT * FROM TEST WHERE KEY = :key`)
	if err != nil {
		fmt.Println(err)
	}

	results := []test{}
	args := test{Key: "2"}

	err = stmt.Select(&results, args)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("results: %+v\n", results)
}

type test struct {
	Id    int
	Key   string
	Value string
}
