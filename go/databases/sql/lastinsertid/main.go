package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Main is a test to see what the LastInsertId() function on a *sql.Stmt being executed is.
// The expectation was that LastInsertId() would not be thread-safe.
// However, it appears to be threadsafe according to this test.
func main() {
	db, err := sql.Open("sqlite3", "file:sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	defer os.Remove("sqlite.db")

	makeTable := "CREATE TABLE IF NOT EXISTS lastid (id INTEGER PRIMARY KEY, words TEXT);"
	_, err = db.Exec(makeTable)
	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO lastid (words) VALUES ($1);"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	afterGoExec := make(chan struct{}, 1)
	waitForSecondInsert := make(chan struct{}, 1)

	go func() {
		result, _ := stmt.Exec("words")
		afterGoExec <- struct{}{}
		<-waitForSecondInsert
		resultId, _ := result.LastInsertId()
		fmt.Printf("first insert result last id is %d\n", resultId)
		afterGoExec <- struct{}{}
	}()

	<-afterGoExec
	res, _ := stmt.Exec("words")
	waitForSecondInsert <- struct{}{}
	<-afterGoExec
	resId, _ := res.LastInsertId()
	fmt.Printf("second insert result last id is %d\n", resId)
}
