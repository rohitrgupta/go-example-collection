package main

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//  NOTE: Not working for sqlite

type User struct {
	gorm.Model
	Username string
	Orders   []Order
}

type Order struct {
	gorm.Model
	UserID uint
	Price  float64
}

func main() {
	// create_data()
	select_1()
}

func select_1() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	ctx := context.Background()
	user, err := gorm.G[User](db).Preload("Order", nil).Find(ctx)
	// SELECT * FROM users;
	// SELECT * FROM orders WHERE user_id IN (1,2,3,4);
	fmt.Printf("Preload user: %#v\n", len(user))

	// Custom Preloading SQL
	user, err = gorm.G[User](db).Preload("Orders", func(db gorm.PreloadBuilder) error {
		db.Order("orders.price DESC")
		return nil
	}).Find(ctx)

}

func create_data() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// ctx := context.Background()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Order{})

	users := []User{
		{Username: "Jinzhu"},
		{Username: "Jackson"},
		{Username: "Jinze"},
		{Username: "Jinzhu"},
		{Username: "Jinzhu"},
	}

	orders := []Order{
		{UserID: 1, Price: 100.0},
		{UserID: 2, Price: 200.0},
		{UserID: 3, Price: 300.0},
		{UserID: 4, Price: 400.0},
		{UserID: 5, Price: 500.0},
		{UserID: 1, Price: 600.0},
		{UserID: 2, Price: 700.0},
	}
	err1 := db.Create(users)
	fmt.Println("Error:", err1)
	err1 = db.Create(orders)
	fmt.Println("Error:", err1)
	for _, user := range users {
		fmt.Println("User ID:", user.ID) // returns inserted data's primary key
	}
}
