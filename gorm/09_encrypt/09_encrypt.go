package main

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type EncryptedString string

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// dbValue: current field's value in database
func (es *EncryptedString) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	switch value := dbValue.(type) {
	case []byte:
		*es = EncryptedString(bytes.TrimPrefix(value, []byte("hello")))
	case string:
		*es = EncryptedString(strings.TrimPrefix(value, "hello"))
	default:
		return fmt.Errorf("unsupported data %#v", dbValue)
	}
	return nil
}

// ctx: contains request-scoped values
// field: the field using the serializer, contains GORM settings, struct tags
// dst: current model value, `user` in the below example
// fieldValue: current field's value of the dst
func (es EncryptedString) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return "hello" + string(es), nil
}

type User struct {
	gorm.Model
	Password EncryptedString
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	data := User{
		Password: EncryptedString("pass"),
	}

	db.Create(&data)
	// INSERT INTO `serializer_structs` (`password`) VALUES ("hellopass")

	var result User
	db.First(&result, "id = ?", data.ID)
	// result => User{
	//   Password: EncryptedString("pass"),
	// }

	db.Where(User{Password: EncryptedString("pass")}).Take(&result)
	// SELECT * FROM `users` WHERE `users`.`password` = "hellopass"
	fmt.Printf("User with password 'pass': %#v\n", result)

}
