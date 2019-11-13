package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"log"

	"github.com/vinicius/ViajaNet/models"
)

var tempx = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsClientes := models.BuscaCliente()
	tempx.ExecuteTemplate(w, "Index", todosOsClientes)
}

func New(w http.ResponseWriter, r *http.Request) {
	tempx.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		telefone := r.FormValue("telefone")
		email := r.FormValue("email")

		models.AdicionaCliente(nome, telefone, email)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete (w http.ResponseWriter, r *http.Request){
	idDoCliente := r.URL.Query().Get("id")
	models.DeletaCliente(idDoCliente)
	http.Redirect(w,r,"/",301)
}

func Edit (w http.ResponseWriter, r *http.Request){
	idDocli := r.URL.Query().Get("id")
	cli := models.EditaCliente(idDocli)
	tempx.ExecuteTemplate(w, "Edit", cli)
}

func Update (w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		telefone := r.FormValue("telefone")
		email := r.FormValue("email")

		conversaoId, err := strconv.Atoi(id)
		if err != nil{
			log.Println("erro na conversão do id" , err)
		}
		models.AtualizaCliente(conversaoId, nome, telefone, email)
}
http.Redirect(w, r, "/", 301)
}
