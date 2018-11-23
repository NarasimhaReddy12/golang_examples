package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type Model struct {
	Id uint `gorm:primary_key`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	gorm.Model
	Name string
}

type UserModel struct {
	Id int `gorm:primary_key`
	UserName string
	Address string
}

func main() {
	db, err := gorm.Open("mysql", "root:password@/sample?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		log.Println("Connection failed to open")
	}
	log.Println("Connection established...")

	db.CreateTable(&UserModel{})
	db.DropTableIfExists(&UserModel{})
	db.Debug().AutoMigrate(&UserModel{})

	user := &UserModel{
		UserName: "John",
		Address: "New York",
	}

	db.Create(user)

	var users []UserModel = []UserModel{
		UserModel{UserName: "name1", Address: "address1"},
		UserModel{UserName: "name2", Address: "address2"},
		UserModel{UserName: "name3", Address: "address3"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	user = &UserModel{UserName: "John", Address: "New York"}
	db.Find(&user)

	user.Address = "Bangalore"
	db.Save(&user)

	db.Model(&user).Update("UserName", "NewUser Name")


}