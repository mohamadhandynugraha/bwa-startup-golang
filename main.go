package main

import (
	"bwastartup/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection database okay")

	// tampilin data dari table user
	var users []user.User
	length := len(users)

	fmt.Println(length)

	db.Find(&users)

	length = len(users)
	fmt.Println(length)

	// loop di golang
	for _, user := range users {
		fmt.Println("Hello My name is ", user.Name, " and my job is ", user.Occupation)
	}
}
