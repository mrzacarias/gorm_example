package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Product is this example base struct
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	log.Println("=> Opening DB")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	log.Println("=> Migrate the schema")
	db.AutoMigrate(&Product{})

	log.Println("=> Create Item")
	db.Create(&Product{Code: "L1212", Price: 1000})

	var product Product
	log.Println("=> Reading Items")
	db.First(&product, 1) // find product with id 1
	product.PrettyPrint()

	db.First(&product, "code = ?", "L1212") // find product with code l1212
	product.PrettyPrint()

	log.Println("=> Updating Items")
	db.Model(&product).Update("Price", 2000)
	product.PrettyPrint()

	var products []Product
	log.Println("=> Listing Items")
	db.Find(&products)
	for _, p := range products {
		p.PrettyPrint()
	}
}

func (p Product) PrettyPrint() {
	fmt.Printf(
		"\tID: %v\n\tCode: %v\n\ttPrice: %v\n\tCreatedAt: %v\n\tUpdatedAt: %v\n\n",
		p.ID, p.Code, p.Price, p.CreatedAt, p.UpdatedAt,
	)
}
