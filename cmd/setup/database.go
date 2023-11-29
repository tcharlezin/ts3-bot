package setup

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
	"ts3-claimedBot/cmd/models"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: getConnection(),
	}), &gorm.Config{})

	if err != nil {
		log.Println("Erro: ", err)
	}

	_ = db.AutoMigrate(&models.Character{})
	_ = db.AutoMigrate(&models.Guild{})
	_ = db.AutoMigrate(&models.Death{})
	_ = db.AutoMigrate(&models.CharacterGuild{})

	return db
}

func getConnection() *sql.DB {

	dsn := os.Getenv("DSN")
	var counts int
	counts = 0

	var conn *sql.DB
	var err error

	for {
		conn, err = openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready...")
			counts++
		} else {
			log.Println("Connected to Postgres")
			break
		}

		if counts > 10 {
			log.Println(err)
			log.Fatal("Can't connect to database!")
			return nil
		}

		log.Println("Backing off for two seconds")
		time.Sleep(2 * time.Second)
	}

	return conn
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
