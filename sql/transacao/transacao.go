package main

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cSenha, _ := ioutil.ReadFile("c:\\des\\golang\\senha.txt")

	db, err := sql.Open("mysql", "root:"+string(cSenha)+"@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios (id,nome) values(?,?)")

	stmt.Exec(4000, "Bia")
	stmt.Exec(4001, "Carlos")
	_, err = stmt.Exec(1, "Tiago")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		// panic(err)
	}

	tx.Commit()
}
