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

	productRepository := repositories.NewProductRepository(db)

	product1 := models.Product{
		Name:   "Dove",
		Brand:  "Wings",
		UserID: 2,
	}
	errProduct1 := productRepository.CreateProduct(&product1)
	if errProduct1 != nil {
		fmt.Println("error:", errProduct1.Error())
		return
	}
	fmt.Println("Created product success")

}
