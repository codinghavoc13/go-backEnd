package models

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"codinghavoc.com/go-back-end/db_conn"
)

type PostInfoDTO struct {
	RoomId            int64         `json:"room_id" binding:"required"`
	PostId            int64         `json:"post_id"`
	PosterId          int           `json:"poster_id" binding:"required"`
	PostTitle         string        `json:"post_title" binding:"required"`
	PostText          string        `json:"post_text" binding:"required"`
	NumberOfResponses int           `json:"number_of_responses"`
	OrigPostDate      time.Time     `json:"orig_post_date"`
	LastResponseDate  time.Time     `json:"last_response_date"`
	Poster            User          `json:"poster"`
	Responses         []ResponseDTO `json:"responses"`
}

func (p PostInfoDTO) Save() (int64, error) {
	db := db_conn.Connect()
	defer db.Close()

	// insertQry := `INSERT INTO notification_demo.posts(
	// date_last_updated, number_responses, post_date, post_text, post_title, user_id)
	// VALUES ($1, $2, $3, $4, $5, $6) returning post_id`
	insertQry := `INSERT INTO react_forum.posts(
	date_last_updated, number_responses, post_date, post_text, post_title, user_id)
	VALUES ($1, $2, $3, $4, $5, $6) returning post_id`

	var id int64

	stmt, err := db.Prepare(insertQry)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	err = db.QueryRow(insertQry, p.LastResponseDate, p.NumberOfResponses, p.OrigPostDate, p.PostText, p.PostTitle, p.PosterId).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, err
}

/*Will need another version of this that takes in a user ID and only returns that user's posts*/
func GetAllPosts() ([]PostInfoDTO, error) {
	var postInfo = []PostInfoDTO{}
	db := db_conn.Connect()
	defer db.Close()
	//need to generate a new scheme
	// findAllPosts := `SELECT post_id, date_last_updated, number_responses, post_date, post_text, post_title, user_id, room_id
	// FROM notification_demo.posts;`
	findAllPosts := `SELECT post_id, date_last_updated, number_responses, post_date, post_text, post_title, user_id, room_id
	FROM react_forum.posts;`
	rows, err := db.Query(findAllPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post PostInfoDTO
		err := rows.Scan(&post.PostId, &post.LastResponseDate, &post.NumberOfResponses,
			&post.OrigPostDate, &post.PostText, &post.PostTitle, &post.PosterId, &post.RoomId)
		if err != nil {
			return nil, err
		}
		fn, ln := GetUserInfo(int64(post.PosterId))
		post.Poster.ID = post.PosterId
		post.Poster.FirstName = fn
		post.Poster.LastName = ln

		postInfo = append(postInfo, post)
	}
	return postInfo, nil
}
func GetPostsByRoomId(roomId int64) ([]PostInfoDTO, error) {
	var postInfo = []PostInfoDTO{}
	db := db_conn.Connect()
	defer db.Close()
	//need to generate a new scheme
	// findAllPosts := `SELECT post_id, date_last_updated, number_responses, post_date, post_text, post_title, user_id, room_id
	// FROM notification_demo.posts
	// where room_id = $1;`
	findAllPosts := `SELECT post_id, date_last_updated, number_responses, post_date, post_text, post_title, user_id, room_id
	FROM react_forum.posts
	where room_id = $1
	order by date_last_updated desc;`
	rows, err := db.Query(findAllPosts, roomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post PostInfoDTO
		err := rows.Scan(&post.PostId, &post.LastResponseDate, &post.NumberOfResponses,
			&post.OrigPostDate, &post.PostText, &post.PostTitle, &post.PosterId, &post.RoomId)
		if err != nil {
			return nil, err
		}
		fn, ln := GetUserInfo(int64(post.PosterId))
		post.Poster.ID = post.PosterId
		post.Poster.FirstName = fn
		post.Poster.LastName = ln

		postInfo = append(postInfo, post)
	}
	return postInfo, nil
}

func GetSinglePost(postId int64) (PostInfoDTO, error) {
	var postDetail PostInfoDTO
	db := db_conn.Connect()
	defer db.Close()
	//need to get the post info
	// findPost := `SELECT post_id, date_last_updated, number_responses, post_date, post_text, post_title, user_id, room_id
	// FROM notification_demo.posts where post_id = $1`
	findPost := `SELECT post_id, date_last_updated, number_responses, post_date, post_text, post_title, user_id, room_id
	FROM react_forum.posts where post_id = $1`
	row := db.QueryRow(findPost, postId)
	err := row.Scan(&postDetail.PostId, &postDetail.LastResponseDate, &postDetail.NumberOfResponses,
		&postDetail.OrigPostDate, &postDetail.PostText, &postDetail.PostTitle, &postDetail.PosterId, &postDetail.RoomId)
	if err != nil {
		return postDetail, err
	}
	fn, ln := GetUserInfo(int64(postDetail.PosterId))
	postDetail.Poster.ID = postDetail.PosterId
	postDetail.Poster.FirstName = fn
	postDetail.Poster.LastName = ln
	//then get all of the responses and drop that into the responses array
	// getResponsesQry := `select post_id, response_id, user_id, response_text, response_date
	// from notification_demo.responses
	// where post_id = $1
	// order by response_date asc`
	getResponsesQry := `select post_id, response_id, user_id, response_text, response_date
	from react_forum.responses
	where post_id = $1 
	order by response_date asc`
	rows, err := db.Query(getResponsesQry, postId)
	if err != nil {
		fmt.Println("1")
		return postDetail, err
	}
	defer rows.Close()
	for rows.Next() {
		var response ResponseDTO
		err := rows.Scan(&response.PostID, &response.ResponseID, &response.UserID,
			&response.ResponseText, &response.ResponseDate)
		if err != nil {
			fmt.Println("2")
			return postDetail, err
		}
		fn, ln := GetUserInfo(int64(response.UserID))
		response.Responder.ID = int(response.UserID)
		response.Responder.FirstName = fn
		response.Responder.LastName = ln

		postDetail.Responses = append(postDetail.Responses, response)
	}

	return postDetail, nil
}

// func getUserInfo(userId int64)(string, string){
// 	fn, ln := GetUserInfo(int64(post.PosterId))
// 	post.Poster.ID = post.PosterId
// 	post.Poster.FirstName = fn
// 	post.Poster.LastName = ln
// }
