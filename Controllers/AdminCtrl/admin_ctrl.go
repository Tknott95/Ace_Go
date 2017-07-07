package AdminCtrl

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	globals "github.com/tknott95/MasterGo/Globals"
)

func AdminLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("Admin Login Called")

	globals.Tmpl.ExecuteTemplate(w, "admin_signin.gohtml", nil)
}
