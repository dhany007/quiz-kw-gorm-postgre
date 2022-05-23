package main

import (
	"belajar-gorm-postgre/database"
	"belajar-gorm-postgre/repositories"
	"fmt"
)

func main() {
	fmt.Println("gorm-postgree")
	db := database.StartDB()

	productRepository := repositories.NewProductRepository(db)

	// delete product
	productDeleteID := 2
	fmt.Println("\nDelete Product id", productDeleteID)
	errDeleteUpdate := productRepository.DeleteProductByID(uint(productDeleteID))
	if errDeleteUpdate != nil {
		fmt.Println("error:", errDeleteUpdate.Error())
		return
	}
	fmt.Println("delete product success")

}
