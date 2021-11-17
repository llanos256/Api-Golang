package commons

import (
	"log"

	"github.com/llanos256/Api-Golang/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@/tick")

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrando....")

	db.AutoMigrate(&models.Tickets{})
}
	

	