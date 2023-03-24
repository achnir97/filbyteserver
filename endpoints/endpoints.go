package endpoints

import (
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"
	"time"

	"github.com/achnir97/go_lang_filbytes/api"
	_ "github.com/gorilla/mux"
	_ "github.com/rs/cors"
	_ "golang.org/x/crypto/bcrypt"
	_ "gorm.io/gorm"
)

func GetInvFormation(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request parameters
	var fmp_info_for_investors api.Info_For_KSL_FRP_500_and_KSL_100000
	// Connect to the database
	db, err := api.DbConnect()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Execute the raw SQL query
	query := " SELECT * from info_for_ksl_frp_500_and_ksl_100000 ORDER BY DATE DESC LIMIT 1"
	data := db.Raw(query).Scan(&fmp_info_for_investors)
	if data.Error != nil {
		http.Error(w, data.Error.Error(), http.StatusInternalServerError)
		return
	} // Send the response back to the client

	// Execute the raw SQL query

	json.NewEncoder(w).Encode(fmp_info_for_investors)

}

func GetInffrom_25_month(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request parameters
	var fmp_info_for_investors api.Info_For_KSL_FRP_500_and_KSL_100000

	var today_1 string
	// Connect to the database
	db, err := api.DbConnect()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	today := time.Now()
	if today.Day() == 26 {

		today_1 = today.Format("2006-01-02")
	} else {
		prevMonth := today.AddDate(0, -1, 0)
		prevMonthStr := fmt.Sprintf("%d-%02d-26", prevMonth.Year(), prevMonth.Month())
		today_1 = prevMonthStr
	}
	// Execute the raw SQL query
	data := db.Raw("SELECT * from info_for_ksl_frp_500_and_ksl_100000 WHERE DATE =?", today_1).Scan(&fmp_info_for_investors)
	if data.Error != nil {
		http.Error(w, data.Error.Error(), http.StatusInternalServerError)
		return
	} // Send the response back to the client

	json.NewEncoder(w).Encode(fmp_info_for_investors)

}

// end point for the fil_price
func Get_Fil_price(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request parameters
	var fil_price api.Total_Quality_adjP_and_Fil_Reward_for_Vogo_network

	// Connect to the database
	db, err := api.DbConnect()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the raw SQL query
	data := db.Raw("SELECT * FROM  total_quality_adjp_and_fil_reward_for_vogo_networks  order by date desc limit 1 ").Scan(&fil_price)
	if data.Error != nil {
		http.Error(w, data.Error.Error(), http.StatusInternalServerError)
		return
	} // Send the response back to the client

	json.NewEncoder(w).Encode(fil_price)

}

/*

FRP_TABLE_AND_KSL_TABLE for filbytesdev

Insert into info_for_ksl_frp_500_and_ksl_100000 (
    date,
    total_quality_adjp_for_vogo_daily_basis,
    fil_rewards_daily_basis ,
    daily_twenty_five_percent_reward,
    daily_seventy_five_percent_locked_reward ,
    cumulative_fil_reward,
    daily_release_1_180_of_seventy_five_percent_reward ,
    cumulative_twenty_five_percent_reward_plus_1_180_locked_reward ,
    cumulative_seventy_five_percent_locked_reward_minus_1_180_locke,
    vogo_25_percent_reward ,
    vogo_75_percent_reward ,
    daily_twenty_five_percent_reward_for_inv ,
    cumulative_twenty_five_percent_reward_for_inv ,
    daily_seventy_five_percent_locked_reward_for_inv ,
    cumulative_seventy_five_percent_locked_reward_for_inv ,
    daily_release_of_1__180_of_locked_reward_for_inv ,
    cumulative_release_of_1__180_of_locked_reward_for_inv ,
    daily_staking_of_inv ,www
    cumulative_total_staking_of_inv ,
    frp_adj_power_for_inv ,
    total_fil_rewards_for_inv_on_daily_basis ,
    frp_cumulative_fil_sum_for_inv ,
    daily_twenty_five_percent_reward_for_ksl_p1 ,
    daily_seventy_five_percent_reward_for_ksl_p1 ,
    cumulative_twenty_five_percent_reward_for_ksl_p1 ,
    cumulative_seventy_five_percent_reward_for_ksl_p1 ,
    daily_one_eighty_release_for_ksl_p1 ,
    cumulative_of_daily_one_eighty_release_for_ksl_p1 ,
    value_of_seventy_five_percent__locked_reward_for_ksl_p1 ,
    total_f_il_reward_ksl_p1_on_daily_basis ,
    cumulative_total_fil_reward_ksl_p1,
    pledge_investement_present_value_ksl_p1,
    fil_when_deal_is_over_ksl_p1,
    graduation_messages_for_inv) values(
    '2023/2/26',2295,18.38,4.59, 13.78, 2591.48,   0,         1025.56,   1565.92,      779.07,  1194.60,      1.13,   153.77, 3.38,  369.26,
	   2.04,   92.04,     3.17,        245.80,    563, 4.51 ,615.06, 0.69,2.06,0.69,2.06,0,0,0,0,0,0,0,'YOUR STILL YET TO GRADAUTE' );
   {2022-11-27 1922   0     0      0     19.51    0.0812778  4.961278 14.548722  3.6701114      10.75989    0       1.27    0    3.7888331 0.021166667 0.021166667 0.021166667 1.2911667 500  0     5.07  0 0 0 0 0 0 0 0 0 0 0 YOU  ARE  STILL YET TO GRADAUTE}


insert into  info_for_ksl_frp_500_and_ksl_100000 (date, total_quality_adjp_for_vogo_daily_basis, fil_rewards_daily_basis, daily_twenty_five_percent_reward,daily_seventy_five_percent_locked_reward, cumulative_fil_reward, daily_release_1_180_of_seventy_five_percent_reward, cumulative_twenty_five_percent_reward_plus_1_180_locked_reward, cumulative_seventy_five_percent_locked_reward_minus_1_180_locke, vogo_25_percent_reward, vogo_75_percent_reward, daily_twenty_five_percent_reward_for_inv, cumulative_twenty_five_percent_reward_for_inv, daily_seventy_five_percent_locked_reward_for_inv, cumulative_seventy_five_percent_locked_reward_for_inv, daily_release_of_1__180_of_locked_reward_for_inv, cumulative_release_of_1__180_of_locked_reward_for_inv, daily_staking_of_inv, cumulative_total_staking_of_inv, frp_adj_power_for_inv,
total_fil_rewards_for_inv_on_daily_basis, frp_cumulative_fil_sum_for_inv,
daily_twenty_five_percent_reward_for_ksl_p1,
daily_seventy_five_percent_reward_for_ksl_p1,
cumulative_twenty_five_percent_reward_for_ksl_p1,
cumulative_seventy_five_percent_reward_for_ksl_p1,
daily_one_eighty_release_for_ksl_p1,
cumulative_of_daily_one_eighty_release_for_ksl_p1,
value_of_seventy_five_percent__locked_reward_for_ksl_p1,
total_f_il_reward_ksl_p1_on_daily_basis,
cumulative_total_fil_reward_ksl_p1,
pledge_investement_present_value_ksl_p1,
fil_when_deal_is_over_ksl_p1,
graduation_messages_for_inv,
twenty_five_percent_reward_paid_for_ksl_p1,
cumulative_twenty_five_percent_reward_paid_for_ksl_p1,
one_eighty_reward_paid_for_ksl_p1,
cumulative_one_eighty_reward_paid_for_ksl_p1,
cumulative_reward_paid_ksl_p1,
value_of_fil_paid_to_ksl_p1,
cumulative_value_of_fil_paid_to_ksl_p1,
daily_fil_paid_to_inv,
 cumulative_daily_fil_paid_to_inv,
value_of_fil_paid_to_inv,
value_of_total_fil_paid_to_inv)
values (‘2023-02-26’, 2295, 18.38, 4.59, 13.78, 2591.48, 5.62, 1025.56, 1565.92,779.07, 1194.60, 1.13,153.77,3.38,369.26, 2.04, 92.04, 3.17, 92.04, 563, 4.51,
615.06,0.69, 2.06, 0.69, 2.06, 0, 0, 18190, 2.74, 2.74, 24253, 8574, ‘you are not yet graduated’,0, 0, 0, 0, 0, 0,0, 0, 0, 0, 0);
{2023-02-27 2295 18.1 4.525 13.575001 2609.58 0 1038.7847 1570.7954 789.8132 1200.2567 1.1100545 154.88007 3.3301637 370.53873 2.0514445 94.091446 3.161499 95.2015 0 0 0 0 563 0 615.06 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 YOU  ARE  STILL YET TO GRADAUTE}
{2023-02-27 2295 18.1 4.525 13.575001 2609.58 0 1038.7847 1570.7954 -Inf -Inf 1.1100545 154.88007 3.3301637 370.53873 2.0514445 94.091446 3.161499 95.2015 0 0 0 0 563 0 615.06
	+Inf +Inf +Inf +Inf 0.011444445 0.011444445 0 0 0 0 0 0 0 +Inf +Inf +Inf +Inf +Inf YOU  ARE  STILL YET TO GRADAUTE}

	{2023-02-28 2295 0 0 0 2609.58 8.726642 1047.5114 1562.0687 795.11523 1189.5344 0 154.88007 0 368.48016 2.0585485 96.149994 2.0585485 97.26005 0 0 0 0 563 0 615.06 0 0
		1.3660883 4.054171 0.022649003 0.034093447 0 0 0 0 0 0 0
		36694.3 0 0 0 8576.883 YOU  ARE  STILL YET TO GRADAUTE}

		{2023-02-28 2295 0 0 0 2609.58 8.726642 1047.5114 1562.0687 795.11523 1189.5344 0 154.88007 0 368.48016 2.0585485 96.149994 2.0585485 97.26005 0 0 0 0 563 0 615.06 0 0
			1.3660883 4.054171 0.022649003 0.034093447 0 0 0 0 0 0 0 36694.3 0 0 0 8576.883 YOU  ARE  STILL YET TO GRADAUTE}
*/
