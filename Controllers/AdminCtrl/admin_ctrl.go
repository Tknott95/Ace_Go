package AdminCtrl

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func adminLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	println("Admin Login Called")
}
