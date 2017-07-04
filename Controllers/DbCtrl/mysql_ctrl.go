package mydb

import (
	"crypto/sha1"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	println("Lang to delete via id:", lang_to_del)

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

func LangAdd(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var lang_to_add string
	lang_to_add = req.FormValue("lang_add")

	println("Lang to add:", lang_to_add)

	stmt, err := store.DB.Prepare(`INSERT INTO pc_langs(lang_id, lang_name) VALUES(?, ?);`) // `INSERT INTO customer VALUES ("James");`
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

func ApiLangFetch(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	jsonData, err := json.Marshal(FetchLangs())
	if err != nil {
		fmt.Println("error: ", err)
	}

	w.Write(jsonData)
}

func BlogPostFetch() []*dbModels.BlogPost {
	rows, err := store.DB.Query("SELECT * FROM blog_ctrl;")
	if err != nil {
		println("eRROR:", err)
	}
	defer rows.Close()

	posts := []*dbModels.BlogPost{}
	for rows.Next() {
		var post dbModels.BlogPost
		err = rows.Scan(&post.ID, &post.Title, &post.Image, &post.Category, &post.Content, &post.Author, &post.Date)
		if err != nil {
			println("eRROR:", err)
		}
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
	path := filepath.Join(wd, "public", "pics", fname)
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
	blogCategory = req.FormValue("blog_category")
	blogContent = req.FormValue("blog_content")

	println("\n<> BLOG POST TO ADD <>\n")
	println("Post Title :", blogTitle, "\n")
	println("Post Image Name :", blogImage, "\n")
	println("Post Author :", blogAuthor, "\n")
	println("Post Category :", blogCategory, "\n")
	println("Post Content :", blogContent, "\n")

	dbInsert, err := store.DB.Prepare(`INSERT INTO blog_ctrl(blog_id, blog_title, blog_image, blog_category, blog_content, blog_author) VALUES(?, ?, ?, ?, ?, ?);`) // `INSERT INTO customer VALUES ("James");`
	if err != nil {
		println("Unable to insert language into mysql db.")
	}

	result, err := dbInsert.Exec(0, blogTitle, fname, blogCategory, blogContent, blogAuthor)
	if err != nil {
		println("Error adding sql lang")
	}

	fmt.Println(w, "ADD RECORD By NAME:", blogTitle, "RESULT:", result)

	println("Post Title :", blogTitle, "\n")
	println("Post Image Name :", blogImage, "\n")
	println("Post Author :", blogAuthor, "\n")
	println("Post Category :", blogCategory, "\n")
	println("Post Content :", blogContent, "\n")

	http.Redirect(w, req, "/pc_langs", 301)
}
