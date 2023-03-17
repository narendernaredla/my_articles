package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql pkg
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	// username := os.Getenv("db_user")
	// password := os.Getenv("db_pass")
	// dbName := os.Getenv("db_name")
	// dbHost := os.Getenv("db_host")
	dbType := os.Getenv("db_type")
	// dbPort := os.Getenv("db_port")

	// dbURI := fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8", username, password, dbHost, dbPort, dbName)

	conn, err := gorm.Open(dbType, "root:root@123@/my_articles_db?charset=utf8")
	if err != nil {
		fmt.Print(err)
	}

	db = conn
}

//GetDB ... get db connection
func GetDB() *gorm.DB {
	return db
}
