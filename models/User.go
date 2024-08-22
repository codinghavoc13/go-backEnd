package models

import (
	"codinghavoc.com/go-back-end/db_conn"
)

type User struct {
	ID        int    `json:"user_id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

func GetUserInfo(userId int64) (string, string) {
	firstName := ""
	lastName := ""
	db := db_conn.Connect()
	defer db.Close()
	// getUserInfoQry := `select first_name, last_name from notification_demo.users
	// where user_id = $1`
	getUserInfoQry := `select first_name, last_name from react_forum.users
	where user_id = $1`
	row := db.QueryRow(getUserInfoQry, userId)
	err := row.Scan(&firstName, &lastName)
	if err != nil {
		return "", ""
	}

	return firstName, lastName
}
