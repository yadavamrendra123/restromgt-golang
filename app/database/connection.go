package database

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const AESKeySize = 32

var DB *gorm.DB
var key []byte

func InitDB() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)

	var err error
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Database connection successful")
			break
		}
		log.Printf("Attempt %d: failed to connect to database: %v", i+1, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("All attempts to connect to the database have failed: %v", err)
	}

	key = generateAESKey(AESKeySize)

	RunMigrations(DB)

	return DB
}

func generateAESKey(size int) []byte {
	key := make([]byte, size)
	if _, err := rand.Read(key); err != nil {
		panic(err.Error())
	}
	return key
}
