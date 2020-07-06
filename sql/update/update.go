package main

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cSenha, _ := ioutil.ReadFile("c:/des/golang/senha.txt")

	db, err := sql.Open("mysql", "root:"+string(cSenha)+"@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Uóshinton Cleber", 1)
	stmt.Exec("Cheron Istôm", 2)

	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}
