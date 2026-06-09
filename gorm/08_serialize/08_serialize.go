package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Data struct {
	Name string
	Tags Tags                   `gorm:"serializer:json"`
	Info map[string]interface{} `gorm:"serializer:json"`
	Bits Bits                   `gorm:"type:bytes;serializer:gob"`
}

type Tags []string
type Bits []byte

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Data{})

	data := Data{Name: "GORM", Tags: []string{"go", "gorm"}, Info: map[string]interface{}{"V1": 1, "V2": 2}, Bits: []byte{1, 2, 3, 4}}
	db.Create(&data)

	var result Data
	db.First(&result)
	fmt.Println(result)

}
