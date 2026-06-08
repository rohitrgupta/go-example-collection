package main

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})

	// Read
	product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx)       // find product with integer primary key
	products, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx) // find product with code D42
	_ = products

	// Update - update product's price to 200
	n, err := gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
	fmt.Println("Rows affected:", n)
	// Update - update multiple fields
	n, err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, Product{Code: "D42", Price: 100})
	fmt.Println("Rows affected:", n)

	// Delete - delete product
	n, err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
	fmt.Println("Rows affected:", n)
}
