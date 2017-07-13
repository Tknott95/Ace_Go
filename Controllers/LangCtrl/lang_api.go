package LangCtrl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ApiLangFetch(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	jsonData, err := json.Marshal(FetchLangs())
	if err != nil {
		fmt.Println("error: ", err)
	}

	w.Write(jsonData)
}
