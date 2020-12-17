package GormTry

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"log"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}


func quick_start_gorm(){
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil{
		panic("fail to connect database")
	}

	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "D42", Price:100})

	// Read
	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "D42")

	db.Model(&product).Update("Price", 200)
	db.First(&product, 1)
	log.Printf("[quick_start_gorm] (log) first update product:%v\n", product)
	db.Model(&product).Updates(Product{Price:300, Code:"F42"})
	db.First(&product, 1)
	log.Printf("[quick_start_gorm] (log) second update product:%v\n", product)
	db.Model(&product).Updates(map[string]interface{}{"Price": 400, "Code": "G42"})

	db.Delete(&product, 1)

}