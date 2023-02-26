package config 
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	_"context")


	type Config struct {
		User     string
		Password string
		Host     string
		Port     string
		DbName   string
		SslMode  string
	}
	
	func Connect(config *Config) (*gorm.DB, error) {
		dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", config.User, config.Password, config.Host, config.Port, config.DbName, config.SslMode)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf("The database couldn't be connected, Check your error properly\n")
			return nil, err
		}
		fmt.Printf("The database is connected\n")
		return db, nil
	}
