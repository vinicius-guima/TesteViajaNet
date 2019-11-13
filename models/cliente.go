package models

import (
	"github.com/vinicius/ViajaNet/db"
)

type Cliente struct { // etrututa de variavel cliente
	Id       int
	Nome     string
	Telefone string
	Email    string
}

func BuscaCliente() []Cliente { // select no banco com tbl clienetes
	db := db.Conecta()
	selectSql, err := db.Query("select * from cliente")
	if err != nil {
		panic(err.Error())
	}

	c := Cliente{}
	clientes := []Cliente{}

	for selectSql.Next() {
		var id int
		var nome, telefone, email string

		err = selectSql.Scan(&id, &nome, &telefone, &email)
		if err != nil {
			panic(err.Error())
		}
		c.Id = id
		c.Nome = nome
		c.Telefone = telefone
		c.Email = email

		clientes = append(clientes, c)
	}
	defer db.Close()
	return clientes
}

func AdicionaCliente(nome string, telefone string, email string) { //adicionar
	db := db.Conecta()
	insereDadosNoBanco, err := db.Prepare("insert into cliente(nome, telefone, email) values($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, telefone, email)
	defer db.Close()

}

func DeletaCliente(id string){
	db := db.Conecta()
	deleteSql, err := db.Prepare("delete from cliente where id=$1")
	if err != nil{
		panic(err.Error())
	}
	deleteSql.Exec(id)
	defer db.Close()
}


func EditaCliente(id string) Cliente {
	db := db.Conecta()

	cliDoBanco, err := db.Query("select * from cliente where id=$1",id)
	if err != nil {
		panic(err.Error())
	}

	cliParaAtualizar := Cliente{}

	for cliDoBanco.Next() {
		var id int
		var nome, telefone, email string

		err = cliDoBanco.Scan(&id, &nome, &telefone, &email)
		if err != nil {
			panic(err.Error())
		}
		cliParaAtualizar.Id = id
		cliParaAtualizar.Nome = nome
		cliParaAtualizar.Telefone = telefone
		cliParaAtualizar.Email = email
	}
	defer db.Close()
	return cliParaAtualizar
}

  func AtualizaCliente(id int, nome string, telefone string, email string) {
	db := db.Conecta()

	AtualizaCliente, err := db.Prepare("update cliente set nome=$1, telefone=$2, email=$3 where id=$4")
	if err != nil {
		panic(err.Error())
	}
	AtualizaCliente.Exec(nome, telefone, email, id)
	defer db.Close()
} 
