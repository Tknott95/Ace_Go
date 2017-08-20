package Models

import (
	"time"
)

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
	ID            int
	Author        string
	Comment       string
	DatePublished time.Time
	LastUpdated   time.Time
}
