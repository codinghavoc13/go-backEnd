package models

import "codinghavoc.com/go-back-end/db_conn"

type Room struct {
	RoomId          int64  `json:"room_id" binding:"required"`
	RoomTitle       string `json:"room_title" binding:"required"`
	RoomDescription string `json:"room_description" binding:"required"`
}

func GetAllRooms() ([]Room, error) {
	var rooms = []Room{}
	db := db_conn.Connect()
	defer db.Close()
	// qryGetAllRooms := `select room_id, room_title, room_description
	// from notification_demo.rooms
	// order by room_title`
	qryGetAllRooms := `select room_id, room_title, room_description 
	from react_forum.rooms
	order by room_title`
	rows, err := db.Query(qryGetAllRooms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var room Room
		err := rows.Scan(&room.RoomId, &room.RoomTitle, &room.RoomDescription)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}
