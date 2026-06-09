package main

type User1 struct {
	Name  string `gorm:"index"`
	Name2 string `gorm:"index:idx_name,unique"`
	Name3 string `gorm:"index:,sort:desc,collate:utf8,type:btree,length:10,where:name3 != 'jinzhu'"`
	Name4 string `gorm:"uniqueIndex"`
	Age   int64  `gorm:"index:,class:FULLTEXT,comment:hello \\, world,where:age > 10"`
	Age2  int64  `gorm:"index:,expression:ABS(age)"`
}

// MySQL option
type User2 struct {
	Name string `gorm:"index:,class:FULLTEXT,option:WITH PARSER ngram INVISIBLE"`
}

// PostgreSQL option
type User3 struct {
	Name string `gorm:"index:,option:CONCURRENTLY"`
}

// create composite index `idx_member` with columns `name`, `number`
type User4 struct {
	Name   string `gorm:"index:idx_member"`
	Number string `gorm:"index:idx_member"`
}

// create unique composite index `idx_member` with columns `name`, and `number`
type User5 struct {
	Name   string `gorm:"index:idx_member,unique"`
	Number string `gorm:"index:idx_member,unique"`
}

// order columns in composite index `idx_member` with `number` first, then `name`
type User6 struct {
	Name   string `gorm:"index:idx_member,priority:2"`
	Number string `gorm:"index:idx_member,priority:1"`
}

// multiple indexes on the same column, create index `idx_id` with column `member_number`, and create unique index `idx_oid` with column `oid`
type UserIndex1 struct {
	OID          int64  `gorm:"index:idx_id;index:idx_oid,unique"`
	MemberNumber string `gorm:"index:idx_id"`
}

// create check constraint `name_checker` with condition `name <> 'jinzhu'`
type UserIndex2 struct {
	Name  string `gorm:"check:name_checker,name <> 'jinzhu'"`
	Name2 string `gorm:"check:name <> 'jinzhu'"`
	Name3 string `gorm:"check:,name <> 'jinzhu'"`
}

// Foraign key constraint
type User struct {
	CompanyID  int
	Company    Company    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreditCard CreditCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreditCard struct {
	Number string
	UserID uint
}

type Company struct {
	ID   int
	Name string
}

// create composite primary key with columns `id` and `language_code`
type Product1 struct {
	ID           string `gorm:"primaryKey"`
	LanguageCode string `gorm:"primaryKey"`
	Code         string
	Name         string
}

type Product2 struct {
	CategoryID uint64 `gorm:"primaryKey;autoIncrement:false"`
	TypeID     uint64 `gorm:"primaryKey;autoIncrement:false"`
}
