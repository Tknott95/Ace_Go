package blogCtrl

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	AdminCtrl "github.com/tknott95/Ace_Go/Controllers/AdminCtrl"
	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"
	Models "github.com/tknott95/Ace_Go/Models"
)

func BlogPostFetch() []*Models.BlogPost {
	rows, err := mydb.Store.DB.Query("SELECT * FROM blog_ctrl;")
	if err != nil {
		println("eRROR:", err)
	}
	defer rows.Close()

	posts := []*Models.BlogPost{}
	for rows.Next() {
		var post Models.BlogPost
		err = rows.Scan(&post.ID, &post.Title, &post.Image, &post.Category, &post.Content, &post.Author, &post.Date)

		posts = append(posts, &post)
	}
	return posts
}

func ApiBlogFetch(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	jsonData, err := json.Marshal(BlogPostFetch())
	if err != nil {
		fmt.Println("error: ", err)
	}

	w.Write(jsonData)
}

func BlogPostAdd(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var blogTitle string
	// var blogImage string  /* String 4 testing run 1 */
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
	blogTime := time.Now()
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

	println("Post Title :", blogTitle, "\n")
	println("Post Image Name :", blogImage, "\n")
	println("Post Author :", blogAuthor, "\n")
	println("Post Category :", blogCategory, "\n")
	println("Post Content :", blogContent, "\n")

	http.Redirect(w, req, "/blog_posts", 301)
}

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
		fmt.Fprintf(w, "Must be admin!")
	}

	http.Redirect(w, req, "/blog_posts", 301)
}
