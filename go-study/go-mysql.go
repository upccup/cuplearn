package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:test@tcp(127.0.01:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	fi, err := os.Open("./docker-host.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer fi.Close()

	br := bufio.NewReader(fi)

	for {
		content, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		if _, err := tx.Exec(`INSERT INTO docker_host (ip) VALUES (?)`, string(content)); err != nil {
			fmt.Printf("write %s to db failed. Err: %s \n", string(content), err.Error())
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal("err")
	}
}
