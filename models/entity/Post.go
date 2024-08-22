package entity

import "time"

type Post struct {
	PostID            int
	UserID            int
	PostTitle         string
	PostText          string
	PostDate          time.Time
	NumberOfResponses int
	DateLastUpdated   time.Time
}
