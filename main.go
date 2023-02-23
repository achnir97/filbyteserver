package main

import(
	_"fmt"
    _"time" 
	_"context"
	"github.com/gofiber/fiber/v2"
	_"net/http"
	"github.com/achnir97/go_lang_filbytes/api"
	_"os"
	"github.com/gofiber/fiber/v2/middleware/cors"

)

func main() {
	app:=fiber.New()
	app.Use(cors.New(cors.Config{
	AllowOrigins:"*",
	AllowHeaders:"Origin, Content-Type, Accept",
	AllowMethods:"GET. POST, PUT, DELETE",
}))
	app.Get("/apis", api.FIL_Price_n_Block_rewards_for_Each_Node)
	app.Listen(":4000")
}
