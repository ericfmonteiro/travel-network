package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint   `json:"id"`
	Name          string `gorm:"column:name" json:"name"`
	Email         string `gorm:"column:email"`
	CPF           string `gorm:"column:cpf"`
	Password      string `gorm:"column:password"`
	NumberOfTrips string `gorm:"column:numtrips" json:"numtrips"`
	Bio           string `gorm:"column:bio"`
}

type Post struct {
	gorm.Model
	ID       uint      `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	PostDate time.Time `json:"postDate"`
	UserID   uint      `json:"userId"`
	UserName string    `json:"userName"`
}

type Comment struct {
	gorm.Model
	ID      uint   `json:"id"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
	PostID  uint   `json:"postId"`
}

type Like struct {
	gorm.Model
	ID     uint `json:"id"`
	UserID uint `json:"userId"`
	PostID uint `json:"postId"`
}
