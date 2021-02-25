// Author: magician
// File:   main.go
// Date:   2021/2/25
package main

import (
	"gorm.io/driver/mysql"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type HelloWorld struct {
	gorm.Model
	Name string
	Sex bool
	Age int
}

func main()  {
	// MySQL
	dsn := "magician:FTmagic123$@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// PostgreSQL
	//dsn := "host=localhost user=postgres password=123456 dbname=gorm port=5472 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate
	_ = db.AutoMigrate(&HelloWorld{})

	// Create
	db.Create(&HelloWorld{
		Name:  "rui",
		Sex:   false,
		Age:   80,
	})

	// First, Find
	//var hello []HelloWorld
	//db.First(&hello)
	//db.Where("age < ?", 21).Find(&hello)
	//fmt.Println(hello)

	// Update
	//db.Where("id = ?", 1).First(&HelloWorld{}).Update("name", "qimiao")
	//db.Where("id = ?", 1).First(&HelloWorld{}).Updates(HelloWorld{
	//	Name: "qimiaoshuai",
	//	Age: 26,
	//})
	//db.Where("id in (?)", []int{1, 2}).Find(&[]HelloWorld{}).Updates(map[string]interface{}{
	//	"Name": "qimiaoshuai",
	//	"Sex": true,
	//	"Age": 26,
	//})

	// Delete
	//db.Where("id in (?)", []int{1, 2}).Delete(&HelloWorld{})
	//db.Where("id in (?)", []int{1, 2}).Unscoped().Delete(&HelloWorld{})
}
