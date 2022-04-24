package pkg

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbConnect() (*gorm.DB, error) {

	config := InitConfig()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func GetDatabaseConnection(e interface{}) *gorm.DB {
	db, err := dbConnect()

	if err != nil {
		panic(err)
	}

	migrateErr := db.AutoMigrate(&e)
	if migrateErr != nil {
		fmt.Println("    ?? Auto-migration failed - ", err)
		return nil
	}

	return db
}
