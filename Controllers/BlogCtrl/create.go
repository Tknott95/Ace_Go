package blogCtrl

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	AdminCtrl "github.com/tknott95/Ace_Go/Controllers/AdminCtrl"
	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"
)

func BlogPostAdd(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if AdminCtrl.IsAdminLoggedIn() == true {
		var blogTitle string
		var blogAuthor string /* Always Defaults to Trevor Knott on admin */
		var blogCategory string
		var blogContent string
		// var blogDate string @TODO Needs to be time.Now() go way

		blogTitle = req.FormValue("blog_title")
		blogImage, header, err := req.FormFile("blog_image")
		if err != nil {
			panic(err)
		}

		defer blogImage.Close()

		ext := strings.Split(header.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, blogImage)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "Public", "pics", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		// copy
		blogImage.Seek(0, 0)
		io.Copy(nf, blogImage)

		println("File name:", fname)

		blogAuthor = "Trevor Knott"
		blogTime := time.Now().Format("dd-mm-yyyy")
		blogCategory = req.FormValue("blog_category")
		blogContent = req.FormValue("blog_content")

		println("\n<> BLOG POST TO ADD <>\n")
		println("Post Title :", blogTitle, "\n")
		println("Post Image Name :", blogImage, "\n")
		println("Post Author :", blogAuthor, "\n")
		println("Post Category :", blogCategory, "\n")
		println("Post Content :", blogContent, "\n")

		dbInsert, err := mydb.Store.DB.Prepare(`INSERT INTO blog_ctrl(blog_id, blog_title, blog_image, blog_category, blog_content, blog_author, blog_date) VALUES(?, ?, ?, ?, ?, ?, ?);`) // `INSERT INTO customer VALUES ("James");`
		if err != nil {
			println("Unable to insert language into mysql db.")
		}

		result, err := dbInsert.Exec(0, blogTitle, fname, blogCategory, blogContent, blogAuthor, blogTime)
		if err != nil {
			println("Error adding sql lang")
		}

		fmt.Println(w, "ADD RECORD By NAME:", blogTitle, "RESULT:", result)
		/* @TODO log this for analytics */
		println("Post Title :", blogTitle, "\n")
		println("Post Image Name :", blogImage, "\n")
		println("Post Author :", blogAuthor, "\n")
		println("Post Category :", blogCategory, "\n")
		println("Post Content :", blogContent, "\n")

		http.Redirect(w, req, "/blog_posts", 301)
	} else {
		fmt.Fprintf(w, "Must be named Trevor Knott yo he is admin!")
	}
}
