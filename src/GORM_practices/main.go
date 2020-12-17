package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Product struct {
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open("mysql", "test.db")
	if err != nil {
		panic("[main] (error) Fail to connect database")
	}
	defer db.Close()

	//Miagrate the schema
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "l1212", Price: 1000})

	//Read
	var product = Product{}
	db.First(&product, 1)
	log.Printf("[main] (log) id = 1 , product = %v\n", product)
	db.First(&product, "code = ?", "l1212")
	log.Printf("[main] (log) Code = l1212 , product = %v\n", product)

}
