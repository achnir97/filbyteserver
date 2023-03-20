package main

import (
	_ "encoding/json"

	"github.com/achnir97/go_lang_filbytes/api"
	_ "github.com/robfig/cron"
	_ "gorm.io/gorm"
)

type Total_Quality_adjP_and_Fil_Reward_for_Vogo_network struct {
	Date                                    string
	Total_Quality_adjP_For_Vogo_Daily_Basis float32
	Total_FIL_Reward_Vogo_daily_Basis       float32
	Current_Sector_Initial_Pledge_32GB      float32
}

//var hasRunToday bool

func main() {

	// prev_info, err := api.Query_Prev_day_info_For_KSL_FRP()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(prev_info)

	api.KSP_FRP_INFO()

	// api.Calculate_total_FIl_reward_and_total_quality_adj_P_and_Fil_Reward_for_Vogo()
	// //api.KSP_FRP_INFO()

	// db, err := api.DbConnect()
	// if err != nil {
	// 	panic(err)
	// }
	// if db.Migrator().HasTable(&api.Total_Quality_adjP_and_Fil_Reward_for_Vogo_network{}) {
	// 	if err := db.Migrator().DropTable(&api.Total_Quality_adjP_and_Fil_Reward_for_Vogo_network{}); err != nil {
	// 		panic("Failed to drop table!")
	// 	}
	// 	fmt.Println("Table dropped!")
	// }

	// if err := db.AutoMigrate(&api.Total_Quality_adjP_and_Fil_Reward_for_Vogo_network{}); err != nil {
	// 	panic("Failed to create table!")
	// }
	// fmt.Println("Table created!")
}

//db.Create(&api.Total_Quality_adjP_and_Fil_Reward_for_Vogo_network{})
// db.Create(&api.FMP_Investment_Integrated_info)

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

// if !db.Migrator().HasTable(&api.Info_For_KSL_FRP_500_and_KSL_100000{}) {
// 	if err := db.AutoMigrate(&api.Info_For_KSL_FRP_500_and_KSL_100000{}); err != nil {
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

/*func runContinously() {
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

}*/
