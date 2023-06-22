package mysql

import (
	"fmt"
	"log"
	"os"

	"github.com/batthanhvan/src/db/class"
	"github.com/batthanhvan/src/db/student"
	"github.com/batthanhvan/src/db/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormConfig = &gorm.Config{
	SkipDefaultTransaction:                   true,
	DisableAutomaticPing:                     true,
	DisableForeignKeyConstraintWhenMigrating: true,
}

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	DB, err = gorm.Open(mysql.Open(DBURL), gormConfig)

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database", Dbdriver)
	}
	MigrateTable()
}

func MigrateTable() {

	DB.AutoMigrate(&student.Student{})
	DB.AutoMigrate(&class.Class{})
	DB.AutoMigrate(&user.User{})

}
