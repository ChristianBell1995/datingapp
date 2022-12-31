package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDatabase : intializes and returns mysql db
func newDatabase() *gorm.DB {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, DBNAME)

	var db *gorm.DB
	err := retry(10, 3*time.Second, func() (err error) {
		db, err = gorm.Open(mysql.Open(URL))
		return
	})

	if err != nil {
		log.Fatalf(err.Error())
		return db
	}
	fmt.Println("Database connection established")
	return db
}

func retry(attempts int, sleep time.Duration, f func() error) (err error) {
	for i := 0; ; i++ {
		err = f()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)

		log.Println("retrying after error:", err)
	}
	return fmt.Errorf("After %d attempts, unable to connect to the db, the last error was: %s", attempts, err)
}
