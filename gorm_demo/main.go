// Author: magician
// File:   main.go
// Date:   2021/2/25
package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type HelloWorld struct {
	gorm.Model
	Name string
	Sex  bool
	Age  int
}

type User struct {
	gorm.Model
	Name string `gorm:"primary_key;column:user_name;type:varchar(100);"`
}

func (u User) TableName() string {
	return "qm_users"
}

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint
	// one
	IDCard IDCard
	// many2many
	Teachers []Teacher `gorm:"many2many:student_teachers;"`
}

type IDCard struct {
	gorm.Model
	StudentID uint
	Num       int
}

type Teacher struct {
	gorm.Model
	TeacherName string
	// many2many
	Students []Student `gorm:"many2many:student_teachers;"`
}

func main() {
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
	//_ = db.AutoMigrate(&HelloWorld{})
	//_ = db.AutoMigrate(&User{})
	_ = db.AutoMigrate(&Class{}, &Student{}, &IDCard{}, &Teacher{})

	// Create
	//db.Create(&HelloWorld{
	//	Name: "rui",
	//	Sex:  false,
	//	Age:  80,
	//})

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

	i := IDCard{
		Num: 123456,
	}
	t := Teacher{
		TeacherName: "老师傅",
	}
	s := Student{
		StudentName: "qm",
		IDCard: i,
		Teachers: []Teacher{t},
	}
	c := Class{
		ClassName: "奇妙的班级",
		Students: []Student{s},
	}
	//classes -> students -> id_cards -> teachers -> student_teachers
	_ = db.Create(&c).Error

	r := gin.Default()
	r.POST("/student", func(c *gin.Context) {
		var student Student
		_ = c.BindJSON(&student)
		db.Create(&student)
	})
	r.GET("/student/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var student Student
		_ = c.BindJSON(&student)
		db.Preload("Teachers").Preload("IDCard").Where("id = ?", id).First(&student)
		c.JSON(200, gin.H{
			"s": student,
		})
	})
	r.GET("/class/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var class Class
		_ = c.BindJSON(&class)
		db.Preload("Students").Preload("Students.Teachers").Preload("Students.IDCard").Where("id = ?", id).First(&class)
		c.JSON(200, gin.H{
			"c": class,
		})
	})
	r.Run(":8080")
}
