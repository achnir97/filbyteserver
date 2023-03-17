package main

import (
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/achnir97/go_lang_filbytes/api"
	"github.com/achnir97/go_lang_filbytes/endpoints"
	"github.com/gorilla/mux"
	_ "github.com/robfig/cron"
	"github.com/rs/cors"
	_ "gorm.io/gorm"
)

//var hasRunToday bool

func main() {
	updated_time := time.Now().String()
	fmt.Println(updated_time)
	// Run the continous code in a go routine
	//runOnceADay()
	// Run the daily code at 12.00 am midnight every day.
	/*	c:=cron.New()
		c.AddFunc("0 0 * * *", func(){
			if !hasRunToday{
				runOnceADay()
				hasRunToday=true
			}

		})
		c.Start()
		// Reset the flag at midnight {
			for {
			now:=time.Now()
			midnight:=time.Date(now.Year(), now.Month(), now.Day()+1, 0,0,0,0, now.Location())
			time.Sleep(midnight.Sub(now))
			hasRunToday=false
		}*/
	/*app := fiber.New()
	app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: "GET. POST, PUT, DELETE",
	}))*/

	// db := api.DbConnect()

	// if !db.Migrator().HasTable(&api.FMP_Info_for_investor_updates{}) {
	// 	if err := db.AutoMigrate(&api.FMP_Info_for_investor_updates{}); err != nil {
	// 		panic("Failed to create table!")
	// 	}
	// 	fmt.Println("Table created!")
	// } else {
	// 	fmt.Println("Table already exists")
	// }
	//db.Create(&api.Node_Info_Daily_and_FIl_Price{})
	//db.Create( &api.FMP_Investment_Integrated_info)

	//db.Find(&api.FMP_Info_for_investor{},"id=?",1)*/
	//app.Get("/apis", api.FIL_Price_n_Block_rewards_for_Each_Node)
	//app.Listen(":4000")

}

func runContinously() {
	for {
		r := mux.NewRouter()
		// Register the endpoints with gorilla/mux router
		r.HandleFunc("/calculate", endpoints.GetInvFormation).Methods(http.MethodGet)
		corsHandler := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowCredentials: true,
		}).Handler(r)
		log.Fatal(http.ListenAndServe(":8080", corsHandler))
		fmt.Printf("Serving is running at port 8080")
	}
}

// will make an api call once a day
func runOnceADay() {

	api.FIL_Price_n_Block_rewards_for_Each_Node()

}
