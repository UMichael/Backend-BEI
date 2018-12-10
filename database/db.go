package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var db gorm.DB

//Users ... Just a conventional user struct with name and password
type Users struct {
	Name     string
	password string
}

//Question ... Posts struct
type Question struct {
	gorm.Model
	Author    Users `gorm:"foreignkey:UserReferer"`
	Content   string
	UserRefer string
	Comments  []Comment
}

//Comment ... Comment struct
type Comment struct {
	gorm.Model
	Author     Users `gorm:"foreignkey:UserReferer"`
	Content    string
	Acceptable bool
}

//Reply ... Reply struct
type Reply struct {
	gorm.Model
	Author  Users `gorm:"foreignkey:UserReferer"`
	Content string
}

func init() {
	db, err := gorm.Open("postgres", "user=DSC password=DSC dbname=dbname sslmode=disable")
	if err != nil {
		log.Fatalln("error opening database: ", err)
	}
	db.AutoMigrate(&Users{}, &Question{}, &Comment{}, &Reply{})

}

//Create .... CRUD
func (question *Question) Create() {
	db.Create(&question)
	fmt.Println(question)
}

//Delete ... CRUD
func (question *Question) Delete() {
	db.Delete(&question)
}

//Create ... CRUD
func (comment *Comment) Create(question Question) {
	db.Model(&question).Association("Comments").Append(comment)
}

func (comment *Comment) ViewComments(question Question) []Comment {
	//going to get comments and replies to comments
	db.Where(&comment).Association("Question")
}