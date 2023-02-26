package main

import(
	"fmt"
    _"time" 
	_"context"
	"github.com/gofiber/fiber/v2"
	_"net/http"
	"github.com/achnir97/go_lang_filbytes/api"
	"github.com/gofiber/fiber/v2/middleware/cors"

	
)

func main() {
	app:=fiber.New()
	app.Use(cors.New(cors.Config{
	AllowOrigins:"*",
	AllowHeaders:"Origin, Content-Type, Accept",
	AllowMethods:"GET. POST, PUT, DELETE",
}))

    db:=api.DbConnect()
	
	if !db.Migrator().HasTable(&api.FMP_Info_for_investor{}) {
		if err := db.AutoMigrate(&api.FMP_Info_for_investor{});err!=nil {
			panic ("Failed to create table!")
		}
		fmt.Println("Table created!")
	} else {
		fmt.Println("Table already exists")
	}
	//db.Create(&api.FMP_Info_for_investor{})
	
	db.Find(&api.FMP_Info_for_investor{},"id=?",1)

	//app.Get("/apis", api.FIL_Price_n_Block_rewards_for_Each_Node)
	app.Listen(":4000")
}

