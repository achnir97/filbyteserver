package config 
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"context")


type Config struct {
	User string
	Password string 
	Host string 
	Port int 
	DbName string
	SslMOde string 
}

func Connect(Config *Config) (db *gorm.DB, error){
    dsn := fmt.Sprinf("user=%s, password=%s, host=%d, port=%d, db=%s, sslmode=%s", Config.User,Config.Password, Config.Host, Config.Port, Config.DbName, Config.SslMode )
	db, err:=gorm.Open(postgres.Open(dsn, postgres.Config{}))
	
	if err!=nil {
		fmt.Printf("The database couldnt be connected, Check your error properly")
		return db, err
	}
	fmt.Printf("The data base is connected")
	return db, nil 
}
