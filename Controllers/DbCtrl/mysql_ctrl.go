package mydb

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	dbModels "github.com/tknott95/MasterGo/Models"
)

var store = newDB()

func newDB() *dbModels.SQLStore {
	db, err := sql.Open("mysql", "tknott95:Welcome1!@tcp(admininstance.cfchdss74ohb.us-west-1.rds.amazonaws.com:3306)/adminaws?charset=utf8")
	if err != nil {
		println("🔒 Connection to AWS database established.\n")
	}

	return &dbModels.SQLStore{
		DB: db,
	}
}

func FetchLangs() []*dbModels.Lang {
	rows, err := store.DB.Query("SELECT * FROM pc_langs;")
	if err != nil {
		return nil
	}
	defer rows.Close()

	langs := []*dbModels.Lang{}
	for rows.Next() {
		var l dbModels.Lang
		err = rows.Scan(&l.ID, &l.LangName)
		if err != nil {
			return nil
		}
		langs = append(langs, &l)
	}
	return langs
}

func LangDelete(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var lang_to_del string
	lang_to_del = ps.ByName("lang_id")

	println("Lang to delete via id: :", lang_to_del)

	stmt, err := store.DB.Prepare(`DELETE FROM pc_langs WHERE lang_id= ?;`)
	defer stmt.Close()

	rows, err := stmt.Query(lang_to_del)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// ...
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	println(w, "DELETED LANG BY ID:", lang_to_del)

	http.Redirect(w, req, "/pc_langs", 301)
}
