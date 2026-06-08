package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type Employee struct {
	gorm.Model
	Name      string
	ManagerID *uint
	Team      []Employee `gorm:"foreignkey:ManagerID"`
}

// Author has and belongs to many languages, `author_languages` is the join table
type Author struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:author_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func uintPtr(i uint) *uint {
	return &i
}

func main() {
	create_data()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var user User
	db.Preload("CreditCards").First(&user, "name = ?", "John")
	println("User:", user.Name)
	for _, card := range user.CreditCards {
		println("Credit Card:", card.Number)
	}

	var authors []Author
	db.Preload("Languages").Find(&authors)
	for _, author := range authors {
		println("Author:", author.Name)
		for _, lang := range author.Languages {
			println("Language:", lang.Name)
		}
	}
}

func create_data() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &CreditCard{})
	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&Author{}, &Language{})

	users := []User{
		{Name: "John", CreditCards: []CreditCard{{Number: "1234"}, {Number: "5678"}}},
		{Name: "Jane", CreditCards: []CreditCard{{Number: "4321"}, {Number: "8765"}}},
	}
	for _, user := range users {
		db.Create(&user)
	}

	employees := []Employee{
		{Name: "Alice"},
		{Name: "Bob", ManagerID: uintPtr(1)},
		{Name: "Charlie", ManagerID: uintPtr(1)},
	}
	for _, employee := range employees {
		db.Create(&employee)
	}

	authors := []Author{
		{Name: "Author1", Languages: []Language{{Name: "Go"}, {Name: "Python"}}},
		{Name: "Author2", Languages: []Language{{Name: "Java"}, {Name: "C++"}, {Name: "Go"}}},
	}

	for _, author := range authors {
		db.Create(&author)
	}

}
