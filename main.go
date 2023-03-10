package main

import(
	"fmt"
	_"encoding/json"
	"log"
	"net/http"
	_"gorm.io/gorm"
	"github.com/gorilla/mux"
	"github.com/rs/cors"	
	"github.com/achnir97/go_lang_filbytes/endpoints"
	"github.com/achnir97/go_lang_filbytes/api"
	"github.com/robfig/cron"
	"time"
)

var hasRunToday bool

func main() {
	// Run the continous code in a go routine 
	go runContinously()
	
	// Run the daily code at 12.00 am midnight every day. 
	c:=cron.New()
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
	}

}

func  runContinously() {
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
func runOnceADay(){
	err:=api. FIL_Price_n_Block_rewards_for_Each_Node
	if err !=nil{
		fmt.Printf("The Error occured on the execution of the function")
	}
	fmt.Printf("The new data are fetched successfully and stored in the database.")

}

