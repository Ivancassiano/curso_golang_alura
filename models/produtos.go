package models

import (
	"curso_golang_alura/db"
	"database/sql"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Id, Quantidade  int
}

func BuscaTodosOsProdutos() []Produto {
	dbPostgres := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := dbPostgres.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	var produtos []Produto
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

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(dbPostgres)

	return produtos
}

func CriarNovoproduto(nome, descricao string, preco float64, quantidade int) {
	dbPostgres := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := dbPostgres.Prepare("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	_, err = insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(dbPostgres)
}

func DeletaProduto(id string) {
	dbPostgres := db.ConectaComBancoDeDados()

	deletarProduto, err := dbPostgres.Prepare("DELETE FROM produtos WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	_, err = deletarProduto.Exec(id)
	if err != nil {
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(dbPostgres)
}
