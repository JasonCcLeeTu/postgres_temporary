package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   string `gorm:"primaryKey;column:id"`
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

func main() {
	dsn := "user=jason password=97325291 dbname=test host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	defer func() {
		sql, err := db.DB()
		if err != nil {
			log.Println("sql DB error:", err.Error())
		}
		sql.Close()
	}()
	if err != nil {
		log.Fatalf("gorm open error:%s", err.Error())
	}

	tableName := "USER"

	// id, err := uuid.NewUUID()
	// if err != nil {
	// 	log.Fatalf("new UUID error:%s", err.Error())
	// }
	// newUsr := User{
	// 	ID:   id.String(),
	// 	Name: "Miko",
	// 	Age:  20,
	// }

	QueryData(db, tableName)
	//	UpdateData(db, tableName)
	//DeleteData(db, tableName)
	//CreateTable(db, tableName)
	//InertData(db, tableName, newUsr)
	//DropTable(db, tableName)

}

func QueryData(db *gorm.DB, tableName string) { //查詢資料
	var list []User
	if err := db.Table(tableName).Where("age < ?", 30).Select("name,age").Find(&list).Error; err != nil {
		log.Fatalf("query data error:%s", err.Error())
	}

	for _, value := range list {
		log.Printf("%+v \n", value)
	}

}

func InertData(db *gorm.DB, tableName string, data User) { //新增資料
	if err := db.Table(tableName).Create(&data).Error; err != nil {
		log.Fatalf("insert data error:%s", err.Error())
	}
	log.Println("insert data successfully")
}

func CreateTable(db *gorm.DB, tableName string) { //建立table
	if err := db.Table(tableName).AutoMigrate(&User{}); err != nil {
		log.Fatalf("create table error:%s", err.Error())
	}
	log.Println("create table successfully")
}

func UpdateData(db *gorm.DB, tableName string) { //更新資料
	if err := db.Table(tableName).Where("age < ?", 30).Update("name", "Mike").Error; err != nil {
		log.Fatalf("update data error:%s", err.Error())
	}
	log.Println("update data successfully")
}

func DeleteData(db *gorm.DB, tableName string) { //刪除資料
	if err := db.Table(tableName).Where("age <= ?", 30).Delete(&User{}).Error; err != nil {
		log.Fatalf("delete data error:%s", err.Error())
	}

	log.Println("delete data successfully")
}

func DropTable(db *gorm.DB, tableName string) { //刪除table
	if err := db.Migrator().DropTable(tableName); err != nil {
		log.Fatalf("DropTable error:%s", err.Error())
	}
	log.Println("drop table successfully")
}
