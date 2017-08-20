package blogCtrl

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	AdminCtrl "github.com/tknott95/Ace_Go/Controllers/AdminCtrl"
	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"
)

func BlogPostDel(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if AdminCtrl.IsAdminLoggedIn() == true {
		// GET PICNAME VIA FORM VALUE THEN REMOVE PIC FILE ON DELETE. USE same func as Adding Pics @TODO
		var post_to_del string
		post_to_del = ps.ByName("post_id")

		var pic_to_rmv string
		pic_to_rmv = ps.ByName("pic_rmv")

		println("Blog Post to delete via id:", post_to_del)

		stmt, err := mydb.Store.DB.Prepare(`DELETE FROM blog_ctrl WHERE blog_id= ?;`)
		defer stmt.Close()

		rows, err := stmt.Query(post_to_del)
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

		println(w, "DELETED BLOG POST BY ID:", post_to_del)

		// create new file
		currDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		os.Remove(currDir + "../../Public/pics/" + pic_to_rmv)
	} else {
		fmt.Fprintf(w, "Must be named Trevor Knott yo he is admin!")
	}

	http.Redirect(w, req, "/blog_posts", 301)
}
