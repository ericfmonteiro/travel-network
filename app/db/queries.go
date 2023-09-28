package db

import (
	"fmt"
	"time"

	"github.com/ericfmonteiro/travel-network/app/model"
)

func GetByID(id string) model.User {
	return model.User{}
}

func ListUsers() ([]model.User, error) {
	query := "SELECT ID, NAME FROM users"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByName(name string) model.User {
	query := "SELECT * FROM users where name = '%s'"
	formatQuery := fmt.Sprintf(query, name)
	rows, err := database.Query(formatQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var user model.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CPF, &user.Password, &user.NumberOfTrips, &user.Bio); err != nil {
			panic(err)
		}
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return user
}

func CreateUser(user model.User) model.User {
	sqlStatement := `
	INSERT INTO users (name, email, cpf, password, numtrips, bio)
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := database.Exec(sqlStatement, user.Name, user.Email, user.CPF, user.Password, user.NumberOfTrips, user.Bio)
	if err != nil {
		panic(err)
	}
	return user
}

func EditUser(user model.User, username string) model.User {
	sqlStatement := `
	update users
	set name = '%s', email = '%s', cpf = '%s', password = '%s', numtrips = %s, bio = '%s'
	where name = '%s'`

	formatQuery := fmt.Sprintf(sqlStatement, user.Name, user.Email, user.CPF, user.Password, user.NumberOfTrips, user.Bio, username)

	_, err := database.Exec(formatQuery)
	if err != nil {
		panic(err)
	}
	return user
}

func GetConnectedPosts(username string) []model.Post {
	query := `select p.id, p.title, p.content, p.postdate, p.userid, u.name from posts p
	left join users u
	on p.userid = u.id 
	where u.name != '%s'`
	formatQuery := fmt.Sprintf(query, username)

	rows, err := database.Query(formatQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.PostDate, &post.UserID, &post.UserName); err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return posts
}

func CreatePost(post model.Post) model.Post {
	sqlStatement := `
	INSERT INTO posts (title, content, postdate, userid)
	VALUES ($1, $2, $3, $4)`
	_, err := database.Exec(sqlStatement, post.Title, post.Content, time.Now(), post.UserID)
	if err != nil {
		panic(err)
	}
	return post
}

func CreateLike(like model.Like) model.Like {
	sqlStatement := `
	INSERT INTO likes (userid, postid)
	VALUES ($1, $2)`
	_, err := database.Exec(sqlStatement, like.UserID, like.PostID)
	if err != nil {
		panic(err)
	}
	return like
}

func CreateComment(comment model.Comment) model.Comment {
	sqlStatement := `
	INSERT INTO comments (content, userid, postid)
	VALUES ($1, $2, $3)`
	_, err := database.Exec(sqlStatement, comment.Content, comment.UserID, comment.PostID)
	if err != nil {
		panic(err)
	}
	return comment
}

func GetLikesByPost(postId string) []model.Like {
	query := "SELECT * FROM likes where postid = '%s'"
	formatQuery := fmt.Sprintf(query, postId)
	rows, err := database.Query(formatQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var likes []model.Like
	for rows.Next() {
		var like model.Like
		if err := rows.Scan(&like.ID, &like.UserID, &like.PostID); err != nil {
			panic(err)
		}
		likes = append(likes, like)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return likes
}

func GetCommentsByPost(postId string) []model.Comment {
	query := "SELECT * FROM comments where postid = '%s'"
	formatQuery := fmt.Sprintf(query, postId)
	rows, err := database.Query(formatQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.PostID); err != nil {
			panic(err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return comments
}
