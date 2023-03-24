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
	Fil_Price                               float32
}

//var hasRunToday bool

func main() {

	// r := mux.NewRouter()
	// // Register the endpoints with gorilla/mux router
	// r.HandleFunc("/pledge_inv", endpoints.GetInvFormation).Methods("GET")
	// r.HandleFunc("/info_for_25th_month", endpoints.GetInffrom_25_month).Methods("GET")
	// r.HandleFunc("/daily_price", endpoints.Get_Fil_price).Methods("GET")

	// corsHandler := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://222.112.183.197:3001"},
	// 	AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowCredentials: true,
	// }).Handler(r)
	// fmt.Printf("Serving is running at port 8080")
	// log.Fatal(http.ListenAndServe(":8080", corsHandler))
	api.FIL_Price_n_Block_rewards_for_Each_Node_from_API()

	/*t := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 19, 0, 0, 0, time.Local)
	if t.Before(time.Now()) {
		t = t.AddDate(0, 0, 1)
	}

	// Create a ticker with a 1-minute delay
	ticker := time.NewTicker(1 * time.Minute)

	// Run the functions on schedule
	for {
		select {
		case <-ticker.C:
			// Get the current time
			now := time.Now()

			// Check if it's time to run the functions
			if now.Equal(t) {
				// Run the functions
				api.FIL_Price_n_Block_rewards_for_Each_Node_from_API()
				time.Sleep(1 * time.Minute)
				api.FRP_Calculate_total_Fil_Reward_After_fetched_From_API()
				time.Sleep(1 * time.Minute)
				api.Calculate_KSL_FRP_500()

				// Set the schedule for the next day
				t = t.AddDate(0, 0, 1)
			}
		}
	}*/

}
