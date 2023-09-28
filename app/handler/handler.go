package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/ericfmonteiro/travel-network/app/db"
	"github.com/ericfmonteiro/travel-network/app/model"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	authHeader := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]

	decodedAuthHeader, err := base64.StdEncoding.DecodeString(authHeader)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "Unauthorized: Error to decode header")
		return
	}

	credentials := string(decodedAuthHeader)

	credentialsArray := strings.Split(credentials, ":")

	password := db.AuthenticateUser(credentialsArray[0])

	if credentials != fmt.Sprintf("%s:%s", credentialsArray[0], password) {
		ctx.JSON(http.StatusUnauthorized, "Unauthorized: Invalid credentials")
		return
	}
	ctx.JSON(http.StatusAccepted, "Authentication successful")
}

func ListUsers(ctx *gin.Context) {
	users, err := db.ListUsers()
	if err != nil {
		fmt.Printf("ERROR 5")
		panic(err)
	}
	ctx.JSON(http.StatusOK, users)
}

func GetUserByName(ctx *gin.Context) {
	username := ctx.Param("name")

	user := db.GetUserByName(username)
	ctx.JSON(http.StatusOK, user)
}

func CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("invalid user request payload: %w", err))
		return
	}

	createdUser := db.CreateUser(user)
	ctx.JSON(http.StatusCreated, createdUser)
}

func GetConnectedPosts(ctx *gin.Context) {
	username := ctx.Param("username")

	connectedPosts := db.GetConnectedPosts(username)
	ctx.JSON(http.StatusOK, connectedPosts)
}

func EditUser(ctx *gin.Context) {
	username := ctx.Param("username")

	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("invalid user request payload: %w", err))
		return
	}

	editedUser := db.EditUser(user, username)
	ctx.JSON(http.StatusOK, editedUser)
}

func CreatePost(ctx *gin.Context) {
	var post model.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("invalid post request payload: %w", err))
		return
	}

	createdPost := db.CreatePost(post)
	ctx.JSON(http.StatusCreated, createdPost)
}

func CreateLike(ctx *gin.Context) {
	var like model.Like
	if err := ctx.ShouldBindJSON(&like); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("invalid like request payload: %w", err))
		return
	}

	createdLike := db.CreateLike(like)
	ctx.JSON(http.StatusCreated, createdLike)
}

func CreateComment(ctx *gin.Context) {
	var comment model.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("invalid comment request payload: %w", err))
		return
	}

	createdComment := db.CreateComment(comment)
	ctx.JSON(http.StatusCreated, createdComment)
}

func GetLikesByPost(ctx *gin.Context) {
	postId := ctx.Param("postid")

	likes := db.GetLikesByPost(postId)
	ctx.JSON(http.StatusOK, likes)
}

func GetCommentsByPost(ctx *gin.Context) {
	postId := ctx.Param("postid")

	comments := db.GetCommentsByPost(postId)
	ctx.JSON(http.StatusOK, comments)
}
