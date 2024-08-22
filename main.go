package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"codinghavoc.com/go-back-end/models"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

/*
Massive shoutout to https://github.com/schadokar/go-postgres/tree/master for having a solution to
the DB connection setup in db_conn being lost
*/

func main() {
	server := gin.Default()

	server.Use(cors.Default())

	server.GET("/posts/room/:room_id", getPostsByRoomId)
	server.GET("/posts/all/:id", getAllPostsByUser)
	server.GET("/posts/postDetail/:id", getPostById)
	server.POST("/posts/newPost", createPost)
	server.GET("/rooms/all", getRoomList)
	// server.GET("/")
	server.GET("/test", test)

	server.Run()
}

func test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Test works"})
}

func createPost(context *gin.Context) {
	var postInfo models.PostInfoDTO
	err := context.ShouldBindJSON(&postInfo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Missing fields, try again"})
		return
	}
	postInfo.NumberOfResponses = 0
	postInfo.OrigPostDate = time.Now()
	postInfo.LastResponseDate = postInfo.OrigPostDate
	id, err := postInfo.Save()
	postInfo.PostId = id
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Problem occurred while trying to save, try again"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Post created", "post": postInfo})
}

/*
Will need another version of this and the GetPostsAll in PID that take in a user Id and only get posts matching
that user's ID
*/
func getAllPosts(context *gin.Context) {
	// To send an OK message, the first value can either be
	// 200 or http.StatusOK
	postInfo, err := models.GetAllPosts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving post information"})
		return
	}
	context.JSON(http.StatusOK, postInfo)
}

func getAllPostsByUser(context *gin.Context) {
	context.JSON(http.StatusInternalServerError, gin.H{"message": "This has not been set up yet, come back later"})
}

func getPostsByRoomId(context *gin.Context) {
	var room_id int64
	room_id, err := strconv.ParseInt(context.Param("room_id"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	postInfo, err := models.GetPostsByRoomId(room_id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving post information"})
		return
	}
	context.JSON(http.StatusOK, postInfo)
}

func getPostById(context *gin.Context) {
	var id int64
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("Provided post ID was:", id)
	postInfo, err := models.GetSinglePost(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to find a post with the provided ID"})
		return
	}
	context.JSON(http.StatusOK, postInfo)
}

func getRoomList(context *gin.Context) {
	rooms, err := models.GetAllRooms()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving room information"})
		return
	}
	context.JSON(http.StatusOK, rooms)
}
