package db

import (
	"fmt"
	"os"

	"github.com/example/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// if(os.Getenv())
	// First connect to default postgres database
	// dsn := fmt.Sprintf(
	// 	"postgres://%s:%s@%s:%s/postgres",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// )
	// defaultDB, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN:                  os.Getenv("DATABASE_URL"),
	// 	PreferSimpleProtocol: true,
	// }))
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to connect to default database: %v", err)
	// }

	// // Check if our target database exists
	// dbName := os.Getenv("DB_NAME")
	// var count int64
	// defaultDB.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", dbName).Count(&count)

	// // Create database if it doesn't exist
	// if count == 0 {
	// 	createDB := fmt.Sprintf("CREATE DATABASE %s", dbName)
	// 	if err := defaultDB.Exec(createDB).Error; err != nil {
	// 		return nil, fmt.Errorf("failed to create database: %v", err)
	// 	}
	// }

	// // Close the default database connection
	// sqlDB, err := defaultDB.DB()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get database instance: %v", err)
	// }
	// sqlDB.Close()

	// Connect to our target database
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DATABASE_URL"),
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to target database: %v", err)
	}
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.LeaveBalance{},
		&entity.ClockEntry{},
		&entity.Employee{},
		&entity.LeaveRequest{},
		&entity.Organization{},
		&entity.Payroll{},
		&entity.UserOrg{},
		&entity.Invite{},
	); err != nil {
		return nil, fmt.Errorf("auto migration failed: %v", err)
	}
	return db, nil
}
