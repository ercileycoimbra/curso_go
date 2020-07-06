package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Usuario struct de usuário
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	cSenha, _ := ioutil.ReadFile("c:/des/golang/senha.txt")

	db, err := sql.Open("mysql", "root:"+string(cSenha)+"@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var usu Usuario
	db.QueryRow("select id,nome from usuarios where id = ?", id).Scan(&usu.ID, &usu.Nome)

	json, _ := json.Marshal(usu)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	cSenha, _ := ioutil.ReadFile("c:/des/golang/senha.txt")

	db, err := sql.Open("mysql", "root:"+string(cSenha)+"@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var usu Usuario
	var usuarios []Usuario

	q, _ := db.Query("select id,nome from usuarios")
	defer q.Close()

	for q.Next() {
		q.Scan(&usu.ID, &usu.Nome)

		usuarios = append(usuarios, usu)
	}

	json, _ := json.Marshal(usuarios)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))

}
