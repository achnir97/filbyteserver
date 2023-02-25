package config 
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	_"context")


type Config struct {
	User string
	Password string 
	Host string 
	Port int 
	DbName string
	SslMode string 
}

func Connect(Config *Config) (*gorm.DB, error){
    dsn := fmt.Sprintf("user=%s, password=%s, host=%s, port=%d, db=%s, sslmode=%s", Config.User,Config.Password, Config.Host, Config.Port, Config.DbName, Config.SslMode)
	db, err:=gorm.Open(postgres.Open(dsn), &gorm.Config{}) 	
	if err!=nil {
		fmt.Printf("The database couldnt be connected, Check your error properly\n")
		return db, err
	}
	fmt.Printf("The data base is connected\n")
	return db, nil 
}
