package routes

import (
	"net/http"

	"github.com/vinicius/ViajaNet/controllers"
)

func CarregaRotas() { //CARREGA AS ROTAS DE URL
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
