package models

import "time"

type PostDetailDTO struct {
	PostId            int64         `json:"postId"`
	PosterId          int           `json:"posterId" binding:"required"`
	PostTitle         string        `json:"postTitle" binding:"required"`
	PostText          string        `json:"postText" binding:"required"`
	NumberOfResponses int           `json:"number_of_response"`
	OrigPostDate      time.Time     `json:"orig_post_date"`
	LastResponseDate  time.Time     `json:"last_response_date"`
	Responses         []ResponseDTO `json:"responses"`
}
