package blogCommentsCtrl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tknott95/Ace_Go/Controllers/AdminCtrl"
	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"
	Models "github.com/tknott95/Ace_Go/Models"
)

type Comment struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	postID := ps.ByName("pid")
	var commentRes Comment
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&commentRes)

	// redirectURL := "http://trevorknott.io/blog/" + postID

	dbInsert, err := mydb.Store.DB.Prepare(`INSERT INTO blog_comments(post_id, comment_id, comment_author, comment_body) VALUES(?, ?, ?, ?);`) // `INSERT INTO customer VALUES ("James");`
	if err != nil {
		println("Unable to insert comment into mysql db via. ", err)
	}

	result, err := dbInsert.Exec(postID, 0, commentRes.Author, commentRes.Body)
	if err != nil {
		println("Error adding(Exec) sql comment via:  ", err)
	} else {
		fmt.Println(w, result)
	}

	fmt.Println(w, "Comment by:", commentRes.Author, "/n For pid: ", postID, " Msg:", commentRes.Body)
	/* @TODO log this for analytics */

	println("\n✅ Comment Added ✅\n")
	println("Post ID :", postID)
	println("Comment Author ", commentRes.Author)
	println("Comment Body:", commentRes.Body, "\n")

	// http.Redirect(w, req, redirectURL, 301)
}

func DelComment(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if AdminCtrl.IsAdminLoggedIn() == true {
		commentID := ps.ByName("cid")
		// strconv.Atoi(postID)

		redirectURL := "http://trevorknott.io/blog" //+ postID

		println("Comment to delete via PID: " + " CID: " + commentID)

		prep, err := mydb.Store.DB.Prepare(`DELETE FROM blog_comments WHERE comment_id= ?;`)
		defer prep.Close()

		rows, err := prep.Query(commentID)
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

		println(w, "Comment to delete via PID: "+" CID: "+commentID)
		http.Redirect(w, req, redirectURL, 301)

	} else {
		println("Failed Comment Delete for CID: " + ps.ByName("cid"))
		fmt.Fprintf(w, "Must be admin...")
	}
}
