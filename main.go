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

	// created user
	user1 := models.User{
		Email: "andi@koinworks.com",
	}
	errUser1 := userRepository.CreateUser(&user1)
	if errUser1 != nil {
		fmt.Println("error:", errUser1.Error())
	}
	fmt.Println("Created success")

	// get all user
	users, err := userRepository.GetAllUsers()
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	for i, user := range *users {
		fmt.Println("User: ", i+1)
		user.Print()
		fmt.Println()
	}

	// get user by id
	userID := 2
	fmt.Println("\nGet user ID:", userID)
	user2, errUser2 := userRepository.GetUserByID(uint(userID))
	if errUser2 != nil {
		fmt.Println("error:", errUser2.Error())
		return
	}
	user2.Print()

}
