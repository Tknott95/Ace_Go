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
	mux.GET("/pc_langs", langAdd)
	mux.POST("/pc_langs/delete/:pc_lang", langDelete)

	http.ListenAndServe(globals.PortNumber, mux)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Println("📝 Currently on Index page.")

	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func langAdd(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Println("📝 Currently on Language Control page.")

	tmpl.ExecuteTemplate(w, "langs_fetch.gohtml", mydb.FetchLangs())
}

func langDelete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var lang_to_del string
	lang_to_del = ps.ByName("pc_lang")
	println("Wanting to delete: ", lang_to_del)
	http.Redirect(w, req, "/pc_langs", 301)
}
