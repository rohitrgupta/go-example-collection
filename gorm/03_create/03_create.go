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

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// Migrate the schema
	db.AutoMigrate(&User{})

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	// Create a single record

	err = gorm.G[User](db).Create(ctx, &user) // pass pointer of data to Create

	// Create with result
	result := gorm.WithResult()
	err = gorm.G[User](db, result).Create(ctx, &user)

	fmt.Println("User ID:", user.ID)                   // returns inserted data's primary key
	fmt.Println("Result:", result)                     // returns
	fmt.Println("Rows Affected:", result.RowsAffected) // returns inserted records count

	users := []User{
		{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		{Name: "Jackson", Age: 19, Birthday: time.Now()},
	}
	err1 := db.Create(users)
	fmt.Println("Error:", err1)
	for _, user := range users {
		fmt.Println("User ID:", user.ID) // returns inserted data's primary key
	}

	fmt.Println("Batch Insert")
	var users2 = []User{{Name: "jinzhu_1"}, {Name: "jinzhu_10000"}}
	// batch size 100
	db.CreateInBatches(users2, 100)
	for _, user := range users2 {
		fmt.Println("User ID:", user.ID) // returns inserted data's primary key
	}

}
