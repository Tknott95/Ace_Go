package blogCtrl

import (
	"log"

	mydb "github.com/tknott95/Ace_Go/Controllers/DbCtrl"
	Models "github.com/tknott95/Ace_Go/Models"
)

func fetchComments(post_id int) []*Models.Comment {
	cmnt, err := mydb.Store.DB.Prepare("SELECT * FROM blog_comments WHERE post_id=? ORDER BY last_update DESC;")
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
		err = cmntRows.Scan(&comment.ID, &comment.Author, &comment.Comment, &comment.DatePublished, &comment.LastUpdated)
		comments = append(comments, &comment)
		println(comments)
	}

	return comments

}
