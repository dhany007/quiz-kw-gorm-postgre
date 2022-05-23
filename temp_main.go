package main

import (
	"belajar-gorm-postgre/database"
	"belajar-gorm-postgre/models"
	"belajar-gorm-postgre/repositories"
	"fmt"
)

func tempMain() {
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

	// update user
	userID = 2
	fmt.Println("\nUpdate user id", userID)
	userUpdate := models.User{
		Email: "andy.update@gmail.com",
	}
	errUpdate := userRepository.UpdateUserByID(&userUpdate, uint(userID))
	if errUpdate != nil {
		fmt.Println("error:", errUpdate.Error())
	}
	fmt.Println("Updated user success")

	// delete user
	userID = 1
	fmt.Println("\nDelete User ID", userID)
	errDelete := userRepository.DeleteUserByID(uint(userID))
	if errDelete != nil {
		fmt.Println("error:", errDelete.Error())
		return
	}
	fmt.Println("Deleted user success")

	fmt.Println("\nPRODUCT")

	// created product
	productRepository := repositories.NewProductRepository(db)

	product1 := models.Product{
		Name:   "Clear",
		Brand:  "Shampo",
		UserID: 2,
	}
	errProduct1 := productRepository.CreateProduct(&product1)
	if errProduct1 != nil {
		fmt.Println("error:", errProduct1.Error())
		return
	}
	fmt.Println("Created product success")

	// get all product
	fmt.Println("\nGet all products")
	products, errProducts := productRepository.GetAllProducts()
	if errProducts != nil {
		fmt.Println("error: ", errProducts.Error())
		return
	}

	for i, product := range *products {
		fmt.Println("Product ke -", i+1)
		product.Print()
		fmt.Println()
	}

	// get product by id
	productID := 2
	fmt.Println("\nGet product ID:", productID)
	product2, errProduct2 := productRepository.GetProductByID(uint(productID))
	if errProduct2 != nil {
		fmt.Println("error:", errProduct2.Error())
		return
	}
	product2.Print()

	// update product
	productUpdateID := 2
	fmt.Println("\nUpdate Product id", productUpdateID)
	productUpdate := models.Product{
		Name:   "Clear update",
		Brand:  "kebersihan",
		UserID: 2,
	}
	errProductUpdate := productRepository.UpdateProductByID(&productUpdate, uint(productUpdateID))
	if errProductUpdate != nil {
		fmt.Println("error:", errProductUpdate.Error())
		return
	}
	fmt.Println("Updated product success")
}
