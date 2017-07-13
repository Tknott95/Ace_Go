package blogCtrl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ApiBlogFetch(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	jsonData, err := json.Marshal(BlogPostFetch())
	if err != nil {
		fmt.Println("error: ", err)
	}

	w.Write(jsonData)
}
