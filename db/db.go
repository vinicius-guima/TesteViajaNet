package db

import (
	"database/sql"

	_ "github.com/lib/pq" //bliclioteca a ser usada somente na conexão com o banco
)

func Conecta() *sql.DB { //função para conectar com o banco
	conn := "user=postgres dbname=viajanet password=vini host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
