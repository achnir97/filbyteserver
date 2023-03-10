package endpoints

	import (
		"fmt"
		"encoding/json"
		_"log"
		"net/http"
	
		_"golang.org/x/crypto/bcrypt"
		_"gorm.io/gorm"
		"github.com/achnir97/go_lang_filbytes/api"
		_"github.com/gorilla/mux"
		_"github.com/rs/cors"	
	)

func GetInvFormation(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request parameters
		var fmp_info_for_investors api.FMP_Info_for_investor
	  // Connect to the database
		  db := api.DbConnect()
	 // Execute the raw SQL query
		 
		  query:= " SELECT * from  fmp_info_for_investors ORDER BY DATE DESC LIMIT 1"
		  data := db.Raw(query).Scan(&fmp_info_for_investors)
		  if data.Error != nil {
			  http.Error(w, data.Error.Error(), http.StatusInternalServerError)
			  return
	  }// Send the response back to the client
	 
	  json.NewEncoder(w).Encode(fmp_info_for_investors)
	  fmt.Println(fmp_info_for_investors)
  }
  