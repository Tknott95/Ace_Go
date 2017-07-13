package LangCtrl

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"
	Models "github.com/tknott95/Ace_Go/Models"
)

func FetchLangs() []*Models.Lang {
	rows, err := mydb.Store.DB.Query("SELECT * FROM pc_langs;")
	if err != nil {
		return nil
	}
	defer rows.Close()

	langs := []*Models.Lang{}
	for rows.Next() {
		var l Models.Lang
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

	println("Lang to delete via id:", lang_to_del)

	stmt, err := mydb.Store.DB.Prepare(`DELETE FROM pc_langs WHERE lang_id= ?;`)
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

func LangAdd(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var lang_to_add string
	lang_to_add = req.FormValue("lang_add")

	println("Lang to add:", lang_to_add)

	stmt, err := mydb.Store.DB.Prepare(`INSERT INTO pc_langs(lang_id, lang_name) VALUES(?, ?);`) // `INSERT INTO customer VALUES ("James");`
	if err != nil {
		println("Unable to insert language into mysql db.")
	}

	if lang_to_add != "" {
		/* Adding Record */
		result, err := stmt.Exec(0, lang_to_add)
		if err != nil {
			println("Error adding sql lang")
		}

		fmt.Println(w, "ADD RECORD By NAME:", lang_to_add, "RESULT:", result)
	} else {
		fmt.Println(w, "Unable to add NULL FIELDS!", lang_to_add)
	}

	http.Redirect(w, req, "/pc_langs", 301)
}
