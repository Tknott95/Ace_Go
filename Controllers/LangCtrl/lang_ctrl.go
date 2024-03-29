package LangCtrl

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	AdminCtrl "github.com/tknott95/Ace_Go/Controllers/AdminCtrl"
	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"
	globals "github.com/tknott95/Ace_Go/Globals"
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
	if AdminCtrl.IsAdminLoggedIn() == true {
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
	} else {
		fmt.Fprintf(w, "Must be named Trevor Knott yo he is admin!")
	}
}

func LangAdd(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if AdminCtrl.IsAdminLoggedIn() == true {
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
	} else {
		fmt.Fprintf(w, "Must be named Trevor Knott yo he is admin!")
	}
}

func LangUpdate(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if AdminCtrl.IsAdminLoggedIn() == true {

		var updID string
		var newName string

		newName = req.FormValue("lang-title")
		updID = ps.ByName("l-id")

		stmt, err := mydb.Store.DB.Prepare(`UPDATE pc_langs SET lang_name=? WHERE lang_id=?`) // `INSERT INTO customer VALUES ("James");`
		if err != nil {
			println("Unable to insert language into mysql db.")
		}
		result, err := stmt.Exec(newName, updID)
		if err != nil {
			println("Error adding sql lang")
		}

		println("Lang ID: ", updID, " Updated: ", newName, "Mem Address: ", result)

	} else {
		fmt.Fprintf(w, "Must be named Trevor Knott yo he is admin!")
	}
}

func LangSingleFetch(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var langID = ps.ByName("lang-id")

	println("📝 Currently on Edit Lang page.")

	stmt, err := mydb.Store.DB.Prepare("SELECT * FROM pc_langs WHERE lang_id=?")
	if err != nil {
		println("eRROR:", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(langID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	langs := []*Models.Lang{}
	for rows.Next() {
		var lang Models.Lang
		err = rows.Scan(&lang.ID, &lang.LangName)

		langs = append(langs, &lang)
	}

	globals.Tmpl.ExecuteTemplate(w, "langs_edit.gohtml", langs)

}
