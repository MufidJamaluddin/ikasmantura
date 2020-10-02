package main

import (
	error2 "backend/error"
	"backend/utils"
	"database/sql"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	var (
		err   error
		sqlDb *sql.DB
		db    *gorm.DB
	)

	if err = godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	if sqlDb, err = sql.Open("mysql", os.Getenv("DB_URI")); err != nil {
		log.Fatal(err)
	}

	wkhtmltopdf.SetPath(os.Getenv("WKHTMLTOPDF_PATH"))

	sqlSetup(sqlDb)

	defer func() {
		_ = sqlDb.Close()
	}()

	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	//history.Register(db)

	Migrate(db)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		IdleTimeout:   30 * time.Second,
		CaseSensitive: false,
		ErrorHandler:  error2.CustomErrorHandler,
	})

	Route(app, db)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	err = app.Listen(port)
	log.Fatal(err)
}

func sqlSetup(sqlDb *sql.DB) {
	sqlDb.SetConnMaxLifetime(
		time.Minute * utils.TryParseDuration(os.Getenv("MAX_LIFETIME"), 5))
	sqlDb.SetMaxIdleConns(utils.TryParseInt(os.Getenv("MAX_IDLE_CONNECTIONS"), 2))
	sqlDb.SetMaxOpenConns(utils.TryParseInt(os.Getenv("MAX_OPEN_CONNECTIONS"), 10))
}
