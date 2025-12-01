// internal/config/mysql.go
package config

import (
	"fmt"
	"log"
	"user-service/internal/model"

	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", user, password, host, port, dbname)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        return fmt.Errorf("failed to connect to database: %v", err)
    }

    // Ping qilish
    sqlDB, err := DB.DB()
    if err != nil {
        return fmt.Errorf("failed to get sql.DB from gorm: %v", err)
    }

    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)

    if err := sqlDB.Ping(); err != nil {
        return fmt.Errorf("ping to DB failed: %v", err)
    }

   fmt.Println("GORM MySQL connected successfully")

    // AutoMigrate
    if err := DB.AutoMigrate(
       &model.User{},
    ); err != nil {
        return fmt.Errorf("auto migrate failed: %v", err)
    }

    log.Println(" AutoMigrate completed successfully")
    return nil
}