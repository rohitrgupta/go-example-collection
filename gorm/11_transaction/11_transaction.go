package main

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Animal struct {
	Name string
}

func main() {
	// Globally disable
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	// Continuous session mode
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	_ = tx
	// tx.First(&user, 1)
	// tx.Find(&users)
	// tx.Model(&user).Update("Age", 18)

	ctx := context.Background()

	// Basic transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		// Use Generics API inside the transaction
		if err := gorm.G[Animal](tx).Create(ctx, &Animal{Name: "Giraffe"}); err != nil {
			// return any error will rollback
			return err
		}

		if err := gorm.G[Animal](tx).Create(ctx, &Animal{Name: "Lion"}); err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})
	_ = err

	var animal1, animal2 Animal
	// SavePoint and RollbackTo
	tx1 := db.Begin()
	tx1.Create(&animal1)

	tx1.SavePoint("sp1")
	tx1.Create(&animal2)
	tx1.RollbackTo("sp1") // Rollback user2

	tx1.Commit() // Commit user1
}
