package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuarios struct {
	id   int
	nome string
}

func main() {
	cSenha, _ := ioutil.ReadFile("c:/des/golang/senha.txt")

	db, err := sql.Open("mysql", "root:"+string(cSenha)+"@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id,nome from usuarios where id between ? and ?", 1, 5000)
	defer rows.Close()

	var oUser usuarios

	for rows.Next() {
		rows.Scan(&oUser.id, &oUser.nome)

		fmt.Printf("Usuário: %s (%d).\n", oUser.nome, oUser.id)
	}

	//asql,_:=db.("")

}
