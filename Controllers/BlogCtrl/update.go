package blogCtrl

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	AdminCtrl "github.com/tknott95/Ace_Go/Controllers/AdminCtrl"
	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"
)

func BlogUpdate(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if AdminCtrl.IsAdminLoggedIn() == true {
		var blogID string
		var blogTitle string
		var blogAuthor string /* Always Defaults to Trevor Knott on admin */
		var blogCategory string
		var blogContent string
		// var blogTime string

		blogID = req.FormValue("blog-id")

		t := time.Now()

		blogAuthor = req.FormValue("blog-author")
		blogTitle = req.FormValue("blog-title")
		blogTime := t.Format("Mon Jan _2 15:04:05 2006")
		blogCategory = req.FormValue("blog-category")
		blogContent = req.FormValue("blog-content")

		dbInsert, err := mydb.Store.DB.Prepare(`UPDATE blog_ctrl SET blog_title=?, blog_category=?, blog_content=?, blog_author=?, blog_date=? WHERE blog_id=?;`) // `INSERT INTO customer VALUES ("James");`
		if err != nil {
			println("Unable to insert language into mysql db.")
		}

		result, err := dbInsert.Exec(blogTitle, blogCategory, blogContent, blogAuthor, blogTime, blogID)
		if err != nil {
			println("Error adding sql lang")
		}
		fmt.Println(w, "ADD RECORD By NAME:", blogTitle, "RESULT:", result)

		println("Post Title: ", blogTitle, "\n")
		println("Post Author: ", blogAuthor, "\n")
		println("Post Category: ", blogCategory, "\n")
		println("Post Content: ", blogContent, "\n")
		// println("Post Time Published: ", blogTime, "\n")

		http.Redirect(w, req, "/blog_posts", 301)
	} else {
		fmt.Fprintf(w, "Must be named Trevor Knott yo he is admin!")
	}
}
