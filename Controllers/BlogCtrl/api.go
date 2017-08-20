package blogCtrl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"
	Models "github.com/tknott95/Ace_Go/Models"
	blogCommentsCtrl "github.com/tknott95/Ace_Go/Controllers/BlogCtrl/CommentsCtrl"
)

func ApiBlogFetch(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	jsonData, err := json.Marshal(BlogPostFetch())
	if err != nil {
		fmt.Println("error: ", err)
	}

	w.Write(jsonData)
}

func ApiSingleFetch(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	println("📝 Currently on Single Fetch Blog API.")

	post_to_edit := ps.ByName("post-id")
	stmt, err := mydb.Store.DB.Prepare("SELECT * FROM blog_ctrl WHERE blog_id=? ORDER BY blog_id DESC;")
	if err != nil {
		println("eRROR:", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(post_to_edit)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	posts := []*Models.BlogPost{}
	for rows.Next() {
		var post Models.BlogPost
		err = rows.Scan(&post.ID, &post.Title, &post.Image, &post.Category, &post.Content, &post.Author, &post.Date, &post.Likes, &post.Dislikes, &post.Views, &post.Shares)
		post.Comments = blogCommentsCtrl.FetchComments(post.ID)
		posts = append(posts, &post)
	}

	jsonData, err := json.Marshal(posts)
	if err != nil {
		fmt.Println("error: ", err)
	}

	w.Write(jsonData)
}
