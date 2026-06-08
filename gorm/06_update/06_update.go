package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      uint
	Birthday time.Time
}

type Order struct {
	UserId int
	Item   string
}

func main() {

}

func update_1() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var user User
	db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)

	ctx := context.Background()

	// Update with conditions
	ctr, err := gorm.G[User](db).Where("age = ?", true).Update(ctx, "age", 18)
	if err != nil {
		panic("failed to update user")
	}
	fmt.Printf("Updated %d records\n", ctr)

	// update multiple fields
	_, err = gorm.G[User](db).Where("age = ?", 18).Updates(ctx, User{Name: "hello", Age: 17})
	if err != nil {
		panic("failed to update user")
	}
}

func delete_1() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	ctx := context.Background()

	// Delete by ID
	ctr, err := gorm.G[User](db).Where("id = ?", 10).Delete(ctx)
	if err != nil {
		panic("failed to delete user")
	}
	fmt.Printf("Deleted %d records\n", ctr)
}
