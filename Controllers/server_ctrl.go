package srvCtrl

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	mydb "github.com/tknott95/MasterGo/Controllers/DbCtrl"
	globals "github.com/tknott95/MasterGo/Globals"
)

var tmpl = template.Must(template.ParseGlob("./Views/*"))

func InitServer() {
	mux := httprouter.New()

	mux.GET("/", index)
	mux.GET("/1", lang_control)

	http.ListenAndServe(globals.PortNumber, mux)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Println("📝 Currently on Index page.")

	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func lang_control(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Println("📝 Currently on Language Control page.")

	tmpl.ExecuteTemplate(w, "langs_fetch.gohtml", mydb.FetchLangs())
}
