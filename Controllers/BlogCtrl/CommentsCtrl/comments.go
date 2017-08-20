package blogCommentsCtrl

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"
	Models "github.com/tknott95/Ace_Go/Models"
)

func FetchComments(post_id int) []*Models.Comment {
	cmnt, err := mydb.Store.DB.Prepare("SELECT * FROM blog_comments WHERE post_id=?;")
	if err != nil {
		println("eRROR:", err)
	}
	defer cmnt.Close()

	comments := []*Models.Comment{}

	cmntRows, err := cmnt.Query(post_id)
	if err != nil {
		log.Fatal(err)
	}
	defer cmntRows.Close()

	for cmntRows.Next() {
		var comment Models.Comment
		err = cmntRows.Scan(&comment.Post_ID, &comment.ID, &comment.Author, &comment.Comment, &comment.DatePublished, &comment.LastUpdated)
		comments = append(comments, &comment)
	}

	return comments
}

func AddComment(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	postID := ps.ByName("pid")
	// strconv.Atoi(postID)
	author := req.FormValue("author")
	body := req.FormValue("body")

	dbInsert, err := mydb.Store.DB.Prepare(`INSERT INTO blog_comments(post_id, comment_id, comment_author, comment_body) VALUES(?, ?, ?, ?);`) // `INSERT INTO customer VALUES ("James");`
	if err != nil {
		println("Unable to insert language into mysql db.")
	}

	result, err := dbInsert.Exec(postID, 0, author, body)
	if err != nil {
		println("Error adding sql lang")
	}

	fmt.Println(w, "Comment by:", author, "/n For pid: ", postID, " RESULT:", result)
	/* @TODO log this for analytics */

	println("\n✅ Comment Added ✅\n")
	println("Post ID :", postID)
	println("Comment Author ", author)
	println("Comment Body:", body, "\n")
	fmt.Println(w, "Comment by:", author, "/n For pid: ", postID, " RESULT:", result)

	//http.Redirect(w, req, "/blog_posts", 301)
}
