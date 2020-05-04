package dbConfig

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var (
	appDBName         = "Employee"
	appDBHost         = "localhost"
	appDBUserName     = "postgres"
	appDBUserPassword = "test123"
	DB                *gorm.DB
)

func InitDB() {
	cnnString := fmt.Sprintf("host=%s port=5432 user=%s password='%s' dbname=%s sslmode=disable",
		appDBHost,
		appDBUserName,
		appDBUserPassword,
		appDBName)

	var err error
	DB, err = gorm.Open("postgres", cnnString)

	// DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=user_people password=test123")

	if err != nil {
		log.Println("DB Error %v", err)
	} else {
		log.Println("DB Connected")
	}
}