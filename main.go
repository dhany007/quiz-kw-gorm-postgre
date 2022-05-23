package main

import (
	"belajar-gorm-postgre/database"
	"belajar-gorm-postgre/models"
	"belajar-gorm-postgre/repositories"
	"fmt"
)

func main() {
	fmt.Println("gorm-postgree")
	db := database.StartDB()

	userRepository := repositories.NewUserRepository(db)

	user1 := models.User{
		Email: "dhany@koinworks.com",
	}
	errUser1 := userRepository.CreateUser(&user1)
	if errUser1 != nil {
		fmt.Println("error:", errUser1.Error())
	}
	fmt.Println("Created success")
}
