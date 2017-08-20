package blogCommentsCtrl

import (
	"log"

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
