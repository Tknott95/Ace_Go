// package srvCtrl

// import (
// 	"database/sql"
// 	"fmt"
// 	"html/template"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// 	globals "github.com/tknott95/MasterGo/Globals"
// )

// var err error

// var db *sql.DB

// var Tmpl = template.Must(template.ParseGlob("./Views/*"))

// func check(err error) {
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func InitServer() {
// 	db, err := sql.Open("mysql", "tknott95:Welcome1!@tcp(admininstance.cfchdss74ohb.us-west-1.rds.amazonaws.com:3306)/adminaws?charset=utf8")
// 	if err != nil {
// 		println("🔒 Connection to AWS database established.\n")
// 	}

// 	check(err)
// 	defer db.Close()

// 	err = db.Ping()
// 	check(err)

// 	mux := httprouter.New()

// 	mux.GET("/", index)
// 	mux.GET("/fetch_langs", fetchLangs)

// 	http.ListenAndServe(globals.PortNumber, mux)
// }

// func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
// 	fmt.Println("📝 Currently on Index page.")

// 	Tmpl.ExecuteTemplate(w, "index.gohtml", nil)
// }

// func fetchLangs(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
// 	rows, err := db.Query(`SELECT lang_name FROM pc_langs;`)
// 	if err != nil {
// 		println("Pc Language Fetch FAILED :(")
// 	}

// 	defer rows.Close()

// 	var name string
// 	var names []string

// 	for rows.Next() {
// 		err = rows.Scan(&name)

// 		names = append(names, name)
// 	}

// 	Tmpl.ExecuteTemplate(w, "langs_fetch.gohtml", names)
// }
