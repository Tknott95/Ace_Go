package blogCtrl

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	blogCommentsCtrl "github.com/tknott95/Ace_Go/Controllers/BlogCtrl/CommentsCtrl"
	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"
	globals "github.com/tknott95/Ace_Go/Globals"
	Models "github.com/tknott95/Ace_Go/Models"
)

func BlogPostFetch() []*Models.BlogPost {
	rows, err := mydb.Store.DB.Query("SELECT * FROM blog_ctrl ORDER BY blog_id DESC;")
	if err != nil {
		println("eRROR:", err)
	}
	defer rows.Close()

	posts := []*Models.BlogPost{}
	for rows.Next() {
		var post Models.BlogPost
		err = rows.Scan(&post.ID, &post.Title, &post.Image, &post.Category, &post.Content, &post.Author, &post.Date, &post.Likes, &post.Dislikes, &post.Views, &post.Shares)
		post.Comments = blogCommentsCtrl.FetchComments(post.ID)
		posts = append(posts, &post)
	}
	return posts
}

func SinglePostFetch(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	println("📝 Currently on Edit Blog page.")

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

	globals.Tmpl.ExecuteTemplate(w, "blog_edit.gohtml", posts)

}
