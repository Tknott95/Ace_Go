package Models

type BlogPost struct {
	ID       int
	Title    string
	Image    string
	Category string
	Content  string
	Author   string
	Date     string
	Likes    int
	Dislikes int
	Views    int
	Shares   int
	Comments []*Comment
}

type Comment struct {
	Post_ID       int
	ID            int
	Author        string
	Comment       string
	DatePublished string
	LastUpdated   string
}
