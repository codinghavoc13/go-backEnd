package models

import "time"

type ResponseDTO struct {
	ResponseID   int64     `json:"response_id" binding:"required"`
	UserID       int64     `json:"user_id" binding:"required"`
	PostID       int64     `json:"post_id" binding:"required"`
	ResponseText string    `json:"response_text" binding:"required"`
	ResponseDate time.Time `json:"response_date" binding:"required"`
	Responder    User      `json:"responder" binding:"required"`
}
