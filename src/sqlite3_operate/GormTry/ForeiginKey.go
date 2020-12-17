package GormTry

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)


type DetectHumanInfo struct{
	//gorm.Model
	GroupId uint `gorm:"primarykey"`
	Sampler string
	SampingLocation string
}

type HistoryData struct {
	gorm.Model
	TotalHydrocarbonVal float64
	GroupId uint
	HumanInfo DetectHumanInfo  `gorm:"foreignKey:GroupId"`
}



func BelongsTo()  {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,         // Disable color
		}, )


	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil{
		panic("fail to connect database")
	}


	db.AutoMigrate(&HistoryData{})


	db.Create(&HistoryData{
		TotalHydrocarbonVal: 88.8,
		GroupId: 2,
		HumanInfo: DetectHumanInfo{
			GroupId: 2,
			Sampler: "victory",
		},
	})


}