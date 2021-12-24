package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type Token struct {
	Token string `gorm:"primaryKey"`
}

//Connects to db
func InitDb() {

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Print(err)
	}

	db = conn
}

//Creates a token table if it does not exist and deletes all contents
func Migrate() {
	db.AutoMigrate(&Token{})
	db.Debug().Where("1 = 1").Delete(&Token{})
}

//Continously receives tokens from a deduplication process
//Buffers these tokens and inserts them in batches to the db
//Buffer size of 100000 is picked to avoid storing every token in buffer
//Batch size of 1000 allows us to parallelize the insertions
func SaveTokens(token_chan chan string) {
	const batch_size = 1000

	models := make([]Token, 0, 100000)
	for token := range token_chan {
		models = append(models, Token{Token: token})
		if len(models) == 100000 {
			db.CreateInBatches(models, batch_size)
			models = models[:0]
		}
	}
	if len(models) > 0 {
		db.CreateInBatches(&models, batch_size)
	}
}
