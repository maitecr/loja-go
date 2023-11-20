package models

import (
	db "loja/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetProdutos() []Produto {
	db := db.ConectarBD()

	selectProdutos, err := db.Query("select * from produtos order by id asc")

	if err != nil {
		panic(err)
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	defer db.Close()

	return produtos
}

func CreateProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectarBD()

	insertScript, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertScript.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeleteProduto(id string) {
	db := db.ConectarBD()

	deleteScript, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}

	deleteScript.Exec(id)

	defer db.Close()
}

func EditProduto(id string) Produto {
	db := db.ConectarBD()

	infoBD, err := db.Query("select * from produtos where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoUpdate := Produto{}

	for infoBD.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = infoBD.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoUpdate.Id = id
		produtoUpdate.Nome = nome
		produtoUpdate.Descricao = descricao
		produtoUpdate.Preco = preco
		produtoUpdate.Quantidade = quantidade
	}

	defer db.Close()

	return produtoUpdate
}

func UpdateProduto(nome, descricao string, preco float64, quantidade, id int) {
	db := db.ConectarBD()

	updateScrip, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateScrip.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
