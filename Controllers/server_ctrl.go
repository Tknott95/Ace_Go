package srvCtrl

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	globals "github.com/tknott95/MasterGo/Globals"
)

var tmpl = template.Must(template.ParseGlob("./Views/*"))

func InitServer() {
	mux := httprouter.New()

	mux.GET("/", index)

	http.ListenAndServe(globals.PortNumber, mux)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Println("📝 Currently on Index page.")

	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}
