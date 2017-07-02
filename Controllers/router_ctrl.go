package srvCtrl

import (
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
	mux.GET("/pc_langs", langFetch)
	mux.POST("/pc_langs/delete/:lang_id", mydb.LangDelete) /* Calls both via. url not form val */
	mux.POST("/pc_langs/add", mydb.LangAdd)                /* will use formval in blog portion for sure tho */

	/* UMBRELLA API PORTION */
	/* Will use /api/ always! */
	mux.GET("/api/pc_langs", mydb.ApiLangFetch)

	http.ListenAndServe(globals.PortNumber, mux)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("📝 Currently on Index page.")

	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func langFetch(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("📝 Currently on Language Control page.")

	tmpl.ExecuteTemplate(w, "langs_fetch.gohtml", mydb.FetchLangs())
}
