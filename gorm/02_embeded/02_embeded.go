package main

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded;embeddedPrefix:author_"`
	Upvotes int32
}

// equals
// type Blog struct {
//   ID          int64
//   AuthorName  string
//   AuthorEmail string
//   Upvotes     int32
// }

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// Migrate the schema
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Blog{})

	// Create Author
	err = gorm.G[Author](db).Create(ctx, &Author{Name: "John Doe", Email: "john.doe@example.com"})

}
