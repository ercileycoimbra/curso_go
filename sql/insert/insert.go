package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func connectToMySQL() *sql.DB {
	cSenha, _ := ioutil.ReadFile("c:\\des\\golang\\senha.txt")

	db, err := sql.Open("mysql", "root:"+string(cSenha)+"@/cursogo")
	if err != nil {
		panic(err)
	}

	c1 := make(chan interface{})
	go func() {
		err = db.Ping()
		if err != nil {
			c1 <- err
		} else {
			c1 <- nil
		}
	}()

	select {
	case a := <-c1:
		if a != nil {
			panic(a)
		}
	case <-time.After(time.Second * 3):
		fmt.Println("Erro de conexão - Timeout")
		os.Exit(1)
	}
	return db
}

func main() {
	db := connectToMySQL()
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO USUARIOS (NOME) VALUES(?)")

	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	nID, _ := res.LastInsertId()
	nRows, _ := res.RowsAffected()

	fmt.Printf("ID inserido %d, linhas afetadas:%d\n", nID, nRows)
}
