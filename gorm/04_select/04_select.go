package main

import (
	"context"
	"errors"
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
	// create_data()
	// select_1_first_last()
	// select_2()
	// select_3_where()
	// select_4_order_group()
	// select_5_join()
	select_6_sub_query()
}

func select_6_sub_query() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var users []User
	db.Where("id in (?)", db.Table("orders").Select("user_id")).Find(&users)
	fmt.Printf("Users with Orders: %#v\n", len(users))

}

func select_5_join() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	type result struct {
		Name string
		Item string
	}

	var res []result
	db.Model(&User{}).Select("users.name, orders.item").Joins("left join orders on orders.user_id = users.id").Scan(&res)
	fmt.Printf("Join users and orders: %#v\n", res)

}

func select_4_order_group() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// ctx := context.Background()

	// Order by
	var users1 []User
	db.Order("age desc, name").Find(&users1)
	fmt.Printf("Users ordered by age desc and name: %#v\n", len(users1))

	// Limit
	var users2 []User
	db.Limit(2).Find(&users2)
	fmt.Printf("First 2 users: %#v\n", len(users2))

	type result struct {
		Date  time.Time
		Total int
	}

	var res result

	db.Model(&User{}).Select("max(age) as total").
		Where("name LIKE ?", "J%").Group("name").
		Having("max(age) > ?", 10).First(&res)
	fmt.Printf("Group by name and max age: %#v\n", res)
}

func select_3_where() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	ctx := context.Background()

	var user User
	var users []User
	user, err = gorm.G[User](db).Where("id = ?", 2).First(ctx)
	// 	SELECT * FROM users WHERE id = 2;
	fmt.Printf("User with id 2: %#v, error: %v\n", user.Name, err)

	users, err = gorm.G[User](db).Where("id IN ?", []int{1, 2, 3}).Find(ctx)
	if err != nil {
		fmt.Printf("Error finding users: %v\n", err)
	} else {
		fmt.Printf("Users with ids 1, 2, 3: %#v\n", len(users))
	}

	// selct by user struct
	var user2 User
	db.Where(&User{Name: "Jinzhu", Age: 18}).First(&user2)
	fmt.Printf("User with name jinzhu and age 18: %#v\n", user2.Name)

	// seelect by id list
	var users2 []User
	db.Where([]int64{2, 3}).Find(&users2)
	fmt.Printf("Users with ids 2, 3: %#v\n", len(users2))

	// select by where clause
	var user3 User
	db.Where("name = ? AND age >= ?", "Jinzhu", 18).First(&user3)
	fmt.Printf("User with name Jinzhu and age >= 18: %#v\n", user3.Name)

	// select by Inline Condition
	var user4 User
	db.Find(&user4, "name = ?", "Jinzhu")
	fmt.Printf("User with name Jinzhu: %#v\n", user4.Name)

	// OR condition
	var users3 []User
	db.Where("name = 'Jinzhu'").Or(User{Age: 18}).Find(&users3)
	fmt.Printf("Users with name Jinzhu or age 18: %#v\n", len(users3))

}

func select_2() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var user User
	// var users []User

	db.First(&user)
	fmt.Printf("First user: %#v\n", user.Name)

	result := map[string]interface{}{}
	db.Model(&User{}).First(&result)
	for key, value := range result {
		fmt.Printf("First user as map: %s = %#v\n", key, value)
	}

	result = map[string]interface{}{}
	db.Table("users").Take(&result)
	for key, value := range result {
		fmt.Printf("Take user as map: %s = %#v\n", key, value)
	}

}

func select_1_first_last() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// create_data()
	ctx := context.Background()

	// Get the first record ordered by primary key
	user1, err := gorm.G[User](db).First(ctx)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Printf("First user: %#v, error: %v\n", user1.Name, err)

	// Get one record, no specified order
	user2, err := gorm.G[User](db).Take(ctx)
	// SELECT * FROM users LIMIT 1;
	fmt.Printf("Take user: %#v, error: %v\n", user2.Name, err)

	// Get last record, ordered by primary key desc
	user3, err := gorm.G[User](db).Last(ctx)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Printf("Last user: %#v, error: %v\n", user3.Name, err)

	// check error ErrRecordNotFound
	errors.Is(err, gorm.ErrRecordNotFound)
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
		{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
		{Name: "Jackson", Age: 19, Birthday: time.Now()},
		{Name: "Jinze", Age: 20, Birthday: time.Now()},
		{Name: "Jinzhu", Age: 21, Birthday: time.Now()},
		{Name: "Jinzhu", Age: 22, Birthday: time.Now()},
	}

	orders := []Order{
		{UserId: 1, Item: "Item 1"},
		{UserId: 2, Item: "Item 2"},
		{UserId: 3, Item: "Item 3"},
		{UserId: 4, Item: "Item 4"},
		{UserId: 5, Item: "Item 5"},
		{UserId: 1, Item: "Item 6"},
		{UserId: 2, Item: "Item 7"},
	}
	err1 := db.Create(users)
	fmt.Println("Error:", err1)
	err1 = db.Create(orders)
	fmt.Println("Error:", err1)
	for _, user := range users {
		fmt.Println("User ID:", user.ID) // returns inserted data's primary key
	}
}
