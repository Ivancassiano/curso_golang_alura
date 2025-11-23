package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=cassiano dbname=alura_loja password=cassiano host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Id, Quantidade  int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}
	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Id = id
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	temp.ExecuteTemplate(w, "Index", produtos)
}
