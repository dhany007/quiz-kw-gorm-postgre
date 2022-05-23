package main

import (
	"belajar-gorm-postgre/database"
	"fmt"
)

func main() {
	fmt.Println("gorm-postgree")
	database.StartDB()
}
