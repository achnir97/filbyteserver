package api

import (
	"encoding/json"
	"fmt"
	_ "io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/achnir97/go_lang_filbytes/config"
	_ "github.com/go-playground/validator/v10"
	_ "github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/robfig/cron"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const ATTOFILL = 1e-18 // 1 FIL is equivalent to 1e^18
const BYTES = 1e-12    // 1 terrabytes is equivalent to 1e^12 bytes

// struct to store each node information from api
type MinerDetails struct {
	QualityAdjPower     string `json:"qualityAdjPower"`
	NetworkRawBytePower string `json:"networkRawBytePower"`
	BlocksMined         string `json:"blocksMined"`
	WeightedBlocksMined string `json:"weightedBlocksMined"`
	TotalRewards        string `json:"totalRewards"`
}

// declaring the global variable to know the total_quality adjustedP_on_daily_basis.
var Total_Quality_adjP_on_daily_basis_for_Vogo int
var Total_Quality_adjP_on_daily_basis_for_Inv int

type Fetched_Info struct {
	Id    string       `json:"id"`
	Miner MinerDetails `json:"miner"`
}

// struct to store the price of the filecoin.
type Fil_price struct {
	Filecoin Price `json:"filecoin"`
}

// struct to get the price of the filecoin in KRW from the api call.
type Price struct {
	Krw float32 `json:"krw"`
}

type Node_Related_Info struct {
	Id              string `json:"id"`
	RobustAddress   string `json:"robust"`
	Actor           string `json:"actor"`
	CreateHeight    int    `json:"createHeight"`
	CreateTimestamp int    `json:"createTimestamp"`
	LastSeenHeight  int    `json:"lastSeenHeight"`
	Balance         string `json:"balance"`
	MessageCount    int    `json:"messageCount"`
	Timestamp       int    `json:"timestamp"`
	Miner           struct {
		Owner struct {
			Address string `json:"address"`
			Balance string `json:"balance"`
		} `json:"owner"`
		Worker struct {
			Address string `json:"address"`
			Balance string `json:"balance"`
		} `json:"worker"`
		ControlAddresses []struct {
			Address string `json:"address"`
			Balance string `json:"balance"`
		} `json:"controlAddresses"`
		PeerId                 string   `json:"peerId"`
		MultiAddresses         []string `json:"multiAddresses"`
		SectorSize             int      `json:"sectorSize"`
		RawBytePower           string   `json:"rawBytePower"`
		QualityAdjPower        string   `json:"qualityAdjPower"`
		NetworkQualityAdjPower string   `json:"networkQualityAdjPower"`
		BlocksMined            int      `json:"blocksMined"`
		WeightedBlocksMined    int      `json:"weightedBlocksMined"`
		TotalRewards           string   `json:"totalRewards"`
		Sectors                struct {
			Live       int `json:"live"`
			Active     int `json:"active"`
			Faulty     int `json:"faulty"`
			Recovering int `json:"recovering"`
		} `json:"sectors"`
		PreCommitDeposits       string `json:"PreCommitDeposits"`
		VestingFunds            string `json:"vestingFunds"`
		IntialPledgeRequirement string `json:"initialPledgeRequirement"`
		AvailableBalance        string `json:"availableBalance"`
		SecotorPledgeBalance    string `json:"sectorPledgeBalance"`
		PledgeBalance           string `json:"pledgeBalance"`
		RawBytePowerRank        int    `json:"rawBytePowerRank"`
		QualityAdjPowerRank     int    `json:"qualityAdjPowerRank"`
	} `json:"miner"`

	OwnerMiners []interface{} `json:"ownedMiners"`

	WorkerMiner []interface{} `json:"workerMiners"`
	Address     string        `json:"address"`
}

type Node_Info_Daily_AdjP_and_FIl_Price struct {
	gorm.Model
	Date                                     string
	Fil_Price                                float32
	Current_Sector_Initial_Pledge_32GB       float32
	Fil_Rewards_f01624021_node_1             float32
	Fil_Rewards_f01918123_node_2             float32
	Fil_Rewards_f01987994_node_3             float32
	Cummulative_Fil_Rewards_f01624021_node_1 float32
	Cummulative_Fil_Rewards_f01918123_node_2 float32
	Cummulative_Fil_Rewards_f01987994_node_3 float32
	FRP_f01624021_node_1_adjP                float32
	FRP_f01918123_node_2_adjP                float32
	FRP_f01987994_node_3_adjP                float32
}
type Node_Info_Daily_and_FIl_Price_ struct {
	Date                                     string
	Fil_Price                                float32
	Current_Sector_Initial_Pledge_32GB       float32
	Fil_Rewards_f01624021_node_1             float32
	Fil_Rewards_f01918123_node_2             float32
	Fil_Rewards_f01987994_node_3             float32
	Cummulative_Fil_Rewards_f01624021_node_1 float32
	Cummulative_Fil_Rewards_f01918123_node_2 float32
	Cummulative_Fil_Rewards_f01987994_node_3 float32
	FRP_f01624021_node_1_adjP                float32
	FRP_f01918123_node_2_adjP                float32
	FRP_f01987994_node_3_adjP                float32
}

type Info_For_KSL_FRP_500_and_KSL_100000 struct {
	Date                                                                   string  `json:"column:Date"`
	Total_Quality_adjP_For_Vogo_Daily_Basis                                float32 `json:"total_Quality_adjP_For_Vogo_Daily_Basis"`
	Fil_rewards_Daily_basis                                                float32 `json:"fil_rewards_Daily_basis"`
	Daily_TwentyFive_percent_Reward                                        float32 `json:"daily_TwentyFive_percent_Reward"`
	Daily_SeventyFive_percent_Locked_Reward                                float32 `json:"daily_SeventyFive_percent_Locked_Reward"`
	Cumulative_fil_Reward                                                  float32 `json:"cumulative_fil_Reward"`
	Daily_Release_1_180_of_SeventyFive_percent_Reward                      float32 `json:"daily_Release_1_180_of_SeventyFive_percent_Reward"`
	Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward          float32 `json:"cumulative_TwentyFive_percent_Reward_"`
	Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward float32 `json:"cumulative_SeventyFive_percent_Locked_Reward_"`
	Vogo_25_percent_Reward                                                 float32 `json:"vogo_25_percent_Reward"`
	Vogo_75_percent_Reward                                                 float32 `json:"vogo_75_percent_Reward"`
	Daily_TwentyFive_percent_Reward_for_inv                                float32 `json:"daily_TwentyFive_percent_Reward_for_inv"`
	Cumulative_TwentyFive_percent_Reward_for_Inv                           float32 `json:"cumulative_TwentyFive_percent_Reward_for_inv"`
	Daily_Seventy_five_percent_locked_reward_for_inv                       float32 `json:"daily_Seventy_five_percent_Locked_Reward_for_inv"`
	Cumulative_Seventy_five_percent_locked_reward_for_inv                  float32 `json:"cumulative_Seventy_five_percent_Locked_Reward_for_inv"`
	Daily_Release_of_1__180_of_locked_Reward_for_inv                       float32 `json:"daily_Release_of_1__180_of_Locked_Reward_for_inv"`
	Cumulative_Release_of_1__180_of_locked_Reward_for_inv                  float32 `json:"cumulative_Release_of_1__180_of_Locked_Reward_for_inv"`
	Daily_Staking_of_inv                                                   float32 `json:"daily_Staking_of_inv"`
	Cumulative_Total_staking_of_inv                                        float32 `json:"cumulative_Total_staking_of_inv"`
	FRP_Adj_Power_for_inv                                                  float32 `json:"FRP_Adj_Power_for_inv"`
	Total_Fil_rewards_for_Inv_on_daily_basis                               float32 `json:"total_Fil_rewards_for_Inv_on_daily_basis"`
	Frp_Cumulative_Fil_Sum_for_Inv                                         float32 `json:"frp_Cumulative_Fil_Sum_for_Inv"`
	Daily_twenty_five_percent_Reward_for_KSL_P1                            float32 `json:"daily_twenty_five_percent_Reward_for_KSL_P1"`
	Daily_seventy_five_percent_Reward_for_KSL_P1                           float32 `json:"daily_seventy_five_percent_Reward_for_KSL_P1"`
	Cumulative_TwentyFive_percent_Reward_for_KSL_P1                        float32 `json:"cumulative_TwentyFive_percent_Reward_for_KSL_P1"`
	Cumulative_Seventy_five_percent_Reward_for_KSL_P1                      float32 `json:"cumulative_Seventy_five_percent_Reward_for_KSL_"`
	Daily_One_Eighty_Release_for_KSL_P1                                    float32 `json:"daily_One_Eighty_Release_for_KSL_P1"`
	Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1                      float32 `json:"cumulative_of_Daily_One_Eighty_Release_for_KSL_P1"`
	Value_of_Seventy_Five_percent__locked_Reward_for_KSL_P1                float32 `json:"value_of_Seventy_Five_percent__locked_Reward_KSL_p1"`
	Total_FIl_Reward_KSL_P1_on_Daily_basis                                 float32 `json:"total_FIl_Reward_KSL_P1_on_Daily_basis"`
	Cumulative_Total_Fil_Reward_KSL_P1                                     float32 `json:"cumulative_Total_Fil_Reward_for_KSL_P1"`
	Pledge_investement_Present_value_KSL_p1                                float32 `json:"pledge_investement_Present_value_KSL_p1"`
	Fil_when_deal_is_over_KSL_p1                                           float32 `json:"fil_when_deal_is_over_KSL_p1"`
	Graduation_messages_for_inv                                            string  `json:"graduation_messages_for_inv"`
}

type FMP_Info_for_investor struct {
	gorm.Model
	Date                                      string  `json:"date"`
	Total_Quality_adjP_For_Vogo_Daily_Basis   float32 `json:"total_Quality_adjP_For_Vogo_Daily_Basis"`
	Total_FIL_Reward_Vogo_daily_Basis         float32 `json:"total_FIL_Reward_Vogo_daily_Basis"`
	Total_Quality_adjP_For_Inv_daily_Basis    float32 `json:"total_Quality_adjP_For_Inv_daily_Basis"`
	Total_Quality_adjP_with_increased_FRP_inv float32 `json:"total_Quality_adjP_With_increased_FRP_Inv"`
	Fil_Rewards_on_daily_basis_for_inv        float32 `json:"fil_Rewards_on_daily_basis"`
	Total_Fil_rewards_for_Inv                 float32 `json:"total_Fil_rewards_for_Inv"`
	Total_FIL_Rewards                         float32 `json:"total_FIL_Rewards"`
	Staking_on_daily_basis                    float32 `json:"staking_on_daily_basis"`
	Total_Staking                             float32 `json:"total_Staking"`
	Total_Reward_value                        float32 `json:"total_Reward_value"`
	Increased_FRP_on_daily_basis              float32 `json:"increase_FRP"`
	Total_FRP                                 float32 `json:"total_FRP"`
	Paid_Reward_to_Investor                   float32 `json:"paid_Reward_to_Investor"`
	Total_FIL_Paid_to_Investor                float32 `json:"total_FIL_Paid_to_Investor"`
	Value_of_FIL_Paid_to_Investor             float32 `json:"value_of_FIL_Paid_to_Investor"`
	Value_of_Total_FIl_Paid                   float32 `json:"value_of_Total_FIl_Paid"`
}
type Total_Quality_adjP_and_Fil_Reward_for_Vogo_network struct {
	Date                                    string
	Total_Quality_adjP_For_Vogo_Daily_Basis float32
	Total_FIL_Reward_Vogo_daily_Basis       float32
	Current_Sector_Initial_Pledge_32GB      float32
}

// func (Node_Info_Daily_and_FIl_Price *Node_Info_Daily_and_FIl_Price) BeforeCreate(*gorm.DB) error {
// 	db.SetColumn("Date", time.Now().Format("2006-01-02"))
// 	return
// }

//Get FIL_Rewards and Quality adjusted power of node f01624021 on daily basis
func FIL_Price_n_Block_rewards_for_Each_Node_from_API() {
	var wg sync.WaitGroup
	db, err := DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	c := make(chan Node_Info_Daily_AdjP_and_FIl_Price, 1)

	//var miner Node_Info_Daily_and_FIl_Price
	var FIL_PRICE float32
	var FIL_REWARDS_f01624021_node_1 float32
	var FIL_REWARDS_f01918123_node_2 float32
	var FIL_REWARDS_f01987994_node_3 float32

	// get the price of FILCOIN on daily basis from the coingecko api.
	wg.Add(1)
	go func() {
		response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=Filecoin&vs_currencies=KRW")
		if err != nil {
			fmt.Printf("Thee http request failed with erro %s\n", err)
			return
		}

		defer response.Body.Close()
		var FilPrice Fil_price
		if err := json.NewDecoder(response.Body).Decode(&FilPrice); err != nil {
			return
		}
		defer wg.Done()
		FIL_PRICE = FilPrice.Filecoin.Krw
		c <- Node_Info_Daily_AdjP_and_FIl_Price{Fil_Price: FIL_PRICE}

		fmt.Printf("The Price of Filecoin is %f\n", FIL_PRICE)
	}()

	// calculate Adjusted power and blocks reward for f01624021 from filfox.info/v1/api
	wg.Add(1)
	go func() {
		response, err := http.Get("https://filfox.info/api/v1/address/f01624021")
		if err != nil {
			fmt.Printf("The http Request failed with error %s\n", err)
			return
		}
		defer response.Body.Close()
		var Miner_Info_f01624021 Node_Related_Info
		if err := json.NewDecoder(response.Body).Decode(&Miner_Info_f01624021); err != nil {
			return
		}
		defer wg.Done()

		var Cummulative_Fil_Rewards_f01624021_node_1 float32

		query := "SELECT cummulative_fil_rewards_f01624021_node_1 FROM node_info_daily_and_f_il_prices ORDER BY id DESC LIMIT 1"

		err = db.Raw(query).Scan(&Cummulative_Fil_Rewards_f01624021_node_1).Error
		if err != nil {
			fmt.Printf("Fil_rewards cannot be fetched, check your Error, %s\n", err)
		}
		prevCummulative_fil_rewards_for_node_1 := Cummulative_Fil_Rewards_f01624021_node_1
		fmt.Printf("The pervious cumulative_fil_rewards_for_node_1 is %f\n", prevCummulative_fil_rewards_for_node_1)

		f, err := strconv.ParseFloat(Miner_Info_f01624021.Miner.TotalRewards, 64)
		/*Information about the total rewards is string type so we convertinto float */
		if err != nil {
			fmt.Printf("Total Rewards Cannot be converted into integer, Check your error %s\n", err)
			return
		}
		/*
			total Rewards are then converted into float32 after the conversion into FILL
			Since 1FIL = 1e^18
		*/
		latestCummulative_fil_rewards_for_node_1 := float32(f * 1e-18)

		fmt.Printf("latestCummulative_fil_rewards_for_node_1, %f\n", latestCummulative_fil_rewards_for_node_1)

		if latestCummulative_fil_rewards_for_node_1 > prevCummulative_fil_rewards_for_node_1 {
			FIL_REWARDS_f01624021_node_1 = float32(latestCummulative_fil_rewards_for_node_1 - prevCummulative_fil_rewards_for_node_1)
		} else {
			FIL_REWARDS_f01624021_node_1 = 0
			fmt.Printf("you have no fil_rewards in last 24 hours")
		}

		QualityAdjPower_f01624021_node_1, err := strconv.ParseInt(Miner_Info_f01624021.Miner.QualityAdjPower, 10, 64)

		if err != nil {
			fmt.Printf("Quality adjusted power_cannot be converted into int")
		}
		fmt.Printf("the sector balance for node are %s\n", Miner_Info_f01624021.Miner.SecotorPledgeBalance)
		//QualityAdjPower_f01624021_node_1=QualityAdjPower_f01624021_node_1
		Node_info := <-c // some other go routines has send teh data and is assingin to nod_info
		Node_info.Fil_Rewards_f01624021_node_1 = FIL_REWARDS_f01624021_node_1
		Node_info.FRP_f01624021_node_1_adjP = float32(QualityAdjPower_f01624021_node_1)
		Node_info.Cummulative_Fil_Rewards_f01624021_node_1 = float32(latestCummulative_fil_rewards_for_node_1)

		c <- Node_info

		fmt.Printf("Miner Id : %s\n", Miner_Info_f01624021.Id)
		fmt.Printf("The total_qualityAdj for the node_f01624021 is %d\n", QualityAdjPower_f01624021_node_1)
		fmt.Printf("The total_blocks mined for the node_f01624021 are %f\n", FIL_REWARDS_f01624021_node_1)

	}()

	// calculate Adjusted power and blocks reward fror f01819003
	wg.Add(1)
	go func() {
		response, err := http.Get("https://filfox.info/api/v1/address/f01918123")
		if err != nil {
			fmt.Printf("The http Request failed with error %s\n", err)
			return
		}
		defer response.Body.Close()
		var Miner_Info_f01918123 Node_Related_Info
		if err := json.NewDecoder(response.Body).Decode(&Miner_Info_f01918123); err != nil {
			return
		}
		defer wg.Done()
		var Cummulative_Fil_Rewards_f01918123_node_2 float32

		query := "SELECT Cummulative_Fil_Rewards_f01918123_node_2 FROM node_info_daily_and_f_il_prices ORDER BY id DESC LIMIT 1"

		err = db.Raw(query).Scan(&Cummulative_Fil_Rewards_f01918123_node_2).Error

		if err != nil {
			fmt.Printf("Fil_rewards cannot be fetched, check your Error, %s\n", err)
		}
		prevCummulative_fil_rewards_for_node_2 := Cummulative_Fil_Rewards_f01918123_node_2
		fmt.Printf("The pervious cumulative_fil_rewards_for_node_2 is %f\n", prevCummulative_fil_rewards_for_node_2)

		f, err := strconv.ParseFloat(Miner_Info_f01918123.Miner.TotalRewards, 64)
		if err != nil {
			fmt.Printf("Total Rewards Cannot be converted into integer, Check your error\n")
			return
		}
		latestCummulative_fil_rewards_for_node_2 := float32(f * 1e-18)
		fmt.Printf("latestCummulative_fil_rewards_for_node_2, %f\n", latestCummulative_fil_rewards_for_node_2)

		if latestCummulative_fil_rewards_for_node_2 > prevCummulative_fil_rewards_for_node_2 {

			FIL_REWARDS_f01918123_node_2 = float32(latestCummulative_fil_rewards_for_node_2 - prevCummulative_fil_rewards_for_node_2)

		} else {
			FIL_REWARDS_f01918123_node_2 = 0
			fmt.Printf("you have no fil_rewards in last 24 hours")
		}

		QualityAdjPower_f01918123_node_2, err := strconv.ParseInt(Miner_Info_f01918123.Miner.QualityAdjPower, 10, 64)
		if err != nil {
			fmt.Printf("The Quality adjusted Power of node2 cannot be converted into int")
		}
		Node_info := <-c
		Node_info.Fil_Rewards_f01918123_node_2 = float32(FIL_REWARDS_f01918123_node_2)
		Node_info.FRP_f01918123_node_2_adjP = float32(QualityAdjPower_f01918123_node_2)
		Node_info.Cummulative_Fil_Rewards_f01918123_node_2 = float32(latestCummulative_fil_rewards_for_node_2)

		c <- Node_info

		fmt.Printf("Miner Id : %s\n", Miner_Info_f01918123.Id)
		fmt.Printf("The total_qualityAdj for the node_f01918123 is %d\n", QualityAdjPower_f01918123_node_2)
		fmt.Printf("The total_blocks mined for the node__f01918123 are %f\n", FIL_REWARDS_f01918123_node_2)

	}()

	// calculate Adjusted power and blocks reward fror f01987994
	wg.Add(1)
	go func() {
		response, err := http.Get("https://filfox.info/api/v1/address/f01987994")
		if err != nil {
			fmt.Printf("The http Request failed with error %s\n", err)
			return
		}
		defer response.Body.Close()
		var Miner_Info_f01987994 Node_Related_Info
		if err := json.NewDecoder(response.Body).Decode(&Miner_Info_f01987994); err != nil {
			return
		}
		defer wg.Done()
		FIL_REWARDS_f01987994_node_3 = float32(Miner_Info_f01987994.Miner.BlocksMined)
		var Cummulative_Fil_Rewards_f01987994_node_3 float32

		query := "SELECT Cummulative_Fil_Rewards_f01987994_node_3 FROM node_info_daily_and_f_il_prices ORDER BY id DESC LIMIT 1"

		err = db.Raw(query).Scan(&Cummulative_Fil_Rewards_f01987994_node_3).Error
		if err != nil {
			fmt.Printf("Fil_rewards cannot be fetched, check your Error in Line 196\n")
		}
		prevCummulative_fil_rewards_for_node_3 := Cummulative_Fil_Rewards_f01987994_node_3

		fmt.Printf("The pervious cumulative_fil_rewards_for_node_3 is %f\n", prevCummulative_fil_rewards_for_node_3)

		f, err := strconv.ParseFloat(Miner_Info_f01987994.Miner.TotalRewards, 64)
		if err != nil {
			fmt.Printf("TotalRewards Cannot be converted into integer, Check your error\n")
			return
		}
		latestCummulative_fil_rewards_for_node_3 := float32(f * 1e-18)
		fmt.Printf("latestCummulative_fil_rewards_for_node_3, %f\n", latestCummulative_fil_rewards_for_node_3)

		if latestCummulative_fil_rewards_for_node_3 > prevCummulative_fil_rewards_for_node_3 {

			FIL_REWARDS_f01987994_node_3 = float32(latestCummulative_fil_rewards_for_node_3 - prevCummulative_fil_rewards_for_node_3)

		} else {
			FIL_REWARDS_f01987994_node_3 = 0
			fmt.Printf("you have no fil_rewards in last 24 hours")
		}

		QualityAdjPower_f01987994_node_3, err := strconv.ParseInt(Miner_Info_f01987994.Miner.QualityAdjPower, 10, 64)
		if err != nil {
			fmt.Printf("The Qaulity adjusted Power of node3 cannot be converted into int.")
		}
		Node_info := <-c
		Node_info.Fil_Rewards_f01987994_node_3 = FIL_REWARDS_f01987994_node_3
		Node_info.FRP_f01987994_node_3_adjP = float32(QualityAdjPower_f01987994_node_3)
		Node_info.Cummulative_Fil_Rewards_f01987994_node_3 = float32(latestCummulative_fil_rewards_for_node_3)

		c <- Node_info
		fmt.Printf("Miner Id : %s\n", Miner_Info_f01987994.Id)
		fmt.Printf("The total_qualityAdj for the node_f01987994 is %d\n", QualityAdjPower_f01987994_node_3)
		fmt.Printf("The Fil for the node_f01987994 are %f\n", FIL_REWARDS_f01987994_node_3)

	}()

	wg.Wait()
	Node_info := <-c
	fmt.Printf("The node infos are %v\n", Node_info)
	Node_Info_Daily_and_FIl_Price := Node_Info_Daily_AdjP_and_FIl_Price{}
	Node_Info_Daily_and_FIl_Price.Fil_Price = Node_info.Fil_Price
	Node_Info_Daily_and_FIl_Price.Fil_Rewards_f01624021_node_1 = Node_info.Fil_Rewards_f01624021_node_1
	Node_Info_Daily_and_FIl_Price.Fil_Rewards_f01918123_node_2 = Node_info.Fil_Rewards_f01918123_node_2
	Node_Info_Daily_and_FIl_Price.Fil_Rewards_f01987994_node_3 = Node_info.Fil_Rewards_f01987994_node_3
	Node_Info_Daily_and_FIl_Price.FRP_f01624021_node_1_adjP = Node_info.FRP_f01624021_node_1_adjP
	Node_Info_Daily_and_FIl_Price.FRP_f01918123_node_2_adjP = Node_info.FRP_f01918123_node_2_adjP
	Node_Info_Daily_and_FIl_Price.FRP_f01987994_node_3_adjP = Node_info.FRP_f01987994_node_3_adjP
	Node_Info_Daily_and_FIl_Price.Cummulative_Fil_Rewards_f01624021_node_1 = Node_info.Cummulative_Fil_Rewards_f01624021_node_1
	Node_Info_Daily_and_FIl_Price.Cummulative_Fil_Rewards_f01918123_node_2 = Node_info.Cummulative_Fil_Rewards_f01918123_node_2
	Node_Info_Daily_and_FIl_Price.Cummulative_Fil_Rewards_f01987994_node_3 = Node_info.Cummulative_Fil_Rewards_f01987994_node_3
	db.Create(&Node_Info_Daily_and_FIl_Price)

}

func FRP_Calculate_total_Fil_Reward_After_fetched_From_API() {
	db, err := DbConnect()
	if err != nil {
		fmt.Println(err)
	}

	Node_info := QueryNodeinfo(db)
	fmt.Println(Node_info.FRP_f01624021_node_1_adjP)
	fmt.Println(Node_info.FRP_f01918123_node_2_adjP)
	fmt.Println(Node_info.FRP_f01987994_node_3_adjP)
	fmt.Println(Node_info.Fil_Price)
	fmt.Printf(Node_info.Date)
	fmt.Println(Node_info.Current_Sector_Initial_Pledge_32GB)
	fmt.Println(Node_info.Fil_Rewards_f01624021_node_1)
	fmt.Println(Node_info.Fil_Rewards_f01918123_node_2)
	fmt.Println(Node_info.Fil_Rewards_f01987994_node_3)

	FMP_INFO := Query_Fmp_table(db)
	fmt.Println(FMP_INFO.Date)
	fmt.Println(FMP_INFO.Total_Quality_adjP_For_Vogo_Daily_Basis)
	fmt.Println(FMP_INFO.Total_FIL_Reward_Vogo_daily_Basis)

	//calculate the time
	now := time.Now()
	//initiliaze the instance of struct for FMP_INFO which is the main struct where all the data will be stored.
	FMP_Info := &Total_Quality_adjP_and_Fil_Reward_for_Vogo_network{}

	// Checks if the date is 25th of the month and time is 0.00 am
	if now.Day() == 26 {
		// Since Node_info is updated once everyday at
		total_Quality_adjP_For_Vogo := Node_info.FRP_f01624021_node_1_adjP + Node_info.FRP_f01918123_node_2_adjP + Node_info.FRP_f01987994_node_3_adjP
		FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis = float32(total_Quality_adjP_For_Vogo)

	} else {

		FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis = FMP_INFO.Total_Quality_adjP_For_Vogo_Daily_Basis
		fmt.Printf("Total_Quaity_adjP_For_Vogo_Daily_Basis\n")
	}

	Total_FIL_Reward_Vogo_daily_Basis := Node_info.Fil_Rewards_f01624021_node_1 + Node_info.Fil_Rewards_f01918123_node_2 + Node_info.Fil_Rewards_f01987994_node_3
	FMP_Info.Total_FIL_Reward_Vogo_daily_Basis = Total_FIL_Reward_Vogo_daily_Basis
	fmt.Printf("FMP_Info.Total_FIL_Reward_Vogo_daily_Basis %f\n", FMP_Info.Total_FIL_Reward_Vogo_daily_Basis)

	/*500 Tib will be on the day of investment and wll be used to calculate the increased FRMo the date of investement
	till the 25 th of the each moent
	*/
	db.Create(&FMP_Info)
	//return FMP_Info
}

func QueryNodeinfo(db *gorm.DB) Node_Info_Daily_and_FIl_Price_ {
	query := "SELECT * from node_info_daily_and_f_il_prices ORDER BY DATE ASC LIMIT 1 OFFSET 29"
	var node_info Node_Info_Daily_and_FIl_Price_
	err := db.Raw(query).Scan(&node_info).Error
	if err != nil {
		fmt.Printf("The node information cannot be fetched %s\n", err)
	}
	fmt.Println(node_info)
	return node_info // printing function for flot type.
}

func Query_Fmp_table(db *gorm.DB) Total_Quality_adjP_and_Fil_Reward_for_Vogo_network {
	query := " SELECT * from  fmp_info_for_investors ORDER BY DATE ASC LIMIT 1 OFFSET 28"
	var FMP_Info Total_Quality_adjP_and_Fil_Reward_for_Vogo_network
	err := db.Raw(query).Scan(&FMP_Info).Error
	if err != nil {
		fmt.Printf("The FMP_information cannot be fetched %s\n", err)
	}
	fmt.Println(FMP_Info)
	return FMP_Info
}

//FRP 투자계정 (KSL_FRP_500) 현황
func Calculate_KSL_FRP_500() {
	// db, err := DbConnect()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	Info_from_api, err := Query_total_Quality_and_Fil_reward()
	if err != nil {
		fmt.Println(err.Error())
	}

	prev_fmp_info, err := Query_Prev_day_info_For_KSL_FRP()
	if err != nil {
		fmt.Println(err.Error())
	}
	Total_FiL_Reward_Vogo := Info_from_api.Total_FIL_Reward_Vogo_daily_Basis
	Total_Quality_adjP_on_daily_basis_for_Inv := Info_from_api.Total_Quality_adjP_For_Vogo_Daily_Basis
	date_today := time.Now().Day()

	//date is fetched from the info_from_api table - the latest date is updated
	Date := Info_from_api.Date
	var FMP_INFO Info_For_KSL_FRP_500_and_KSL_100000

	Prev_day_Cumulative_fil_Reward := prev_fmp_info.Cumulative_fil_Reward
	Prev_day_seventyFive_percent_Locked_Reward := prev_fmp_info.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward
	Prev_day_twentyFive_percent_Reward := prev_fmp_info.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward
	Prev_day_Cumulative_twentyFive_percent_Reward_for_inv := prev_fmp_info.Cumulative_TwentyFive_percent_Reward_for_Inv
	Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv := prev_fmp_info.Cumulative_Seventy_five_percent_locked_reward_for_inv
	Prev_day_Cumulative_Release_of_1__180_of_locked_Reward_for_inv := prev_fmp_info.Cumulative_Release_of_1__180_of_locked_Reward_for_inv
	Prev_day_Cumulative_Total_staking_of_inv := prev_fmp_info.Cumulative_Total_staking_of_inv
	Prev_day_FRP_Adj_Power_for_inv := prev_fmp_info.FRP_Adj_Power_for_inv
	Prev_day_Frp_Cumulative_Fil_Sum_for_Inv := prev_fmp_info.Frp_Cumulative_Fil_Sum_for_Inv
	Prev_day_Cumulative_TwentyFive_percent_Reward_for_KSL_P1 := prev_fmp_info.Cumulative_TwentyFive_percent_Reward_for_KSL_P1
	Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1 := prev_fmp_info.Cumulative_Seventy_five_percent_Reward_for_KSL_P1
	Prev_day_Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1 := prev_fmp_info.Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1
	Prev_day_FIl_When_deal_is_over_KSL_p1 := prev_fmp_info.Fil_when_deal_is_over_KSL_p1

	Current_Sector_Initial_Pledge_32GB := Info_from_api.Current_Sector_Initial_Pledge_32GB
	FMP_INFO.Date = Date
	// Querry the FMP_Info_for_investor_updates for previous FMP_Info_for_investor_
	if prev_fmp_info.FRP_Adj_Power_for_inv < 1500.0 {
		FMP_INFO.Graduation_messages_for_inv = "YOUR STILL YET TO GRADAUTE"
		if Total_FiL_Reward_Vogo == 0 {
			FMP_INFO.Fil_rewards_Daily_basis = 0
			FMP_INFO.Cumulative_fil_Reward = Prev_day_Cumulative_fil_Reward
			FMP_INFO.Daily_TwentyFive_percent_Reward = 0.0
			FMP_INFO.Daily_SeventyFive_percent_Locked_Reward = 0.0
			FMP_INFO.Daily_Release_of_1__180_of_locked_Reward = float32(1/180) * Prev_day_seventyFive_percent_Locked_Reward
			FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward = Prev_day_seventyFive_percent_Locked_Reward * (1.0 - 1.0/180.0)
			FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward = Prev_day_twentyFive_percent_Reward + (1.0/180.0)*Prev_day_seventyFive_percent_Locked_Reward
			FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv = 0.0
			FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv = Prev_day_Cumulative_twentyFive_percent_Reward_for_inv
			FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv = 0.0
			FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv = Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv * (1.0 - 1.0/180.0)
			FMP_INFO.Daily_Release_of_1__180_of_locked_Reward_for_inv = float32(1.0/180.0) * Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv
			FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv = Prev_day_Cumulative_Release_of_1__180_of_locked_Reward_for_inv + FMP_INFO.Daily_Release_1_180_of_SeventyFive_percent_Reward
			FMP_INFO.Daily_Staking_of_inv = FMP_INFO.Daily_Release_of_1__180_of_locked_Reward_for_inv
			FMP_INFO.Cumulative_Total_staking_of_inv = Prev_day_Cumulative_Total_staking_of_inv + FMP_INFO.Daily_Staking_of_inv
			FMP_INFO.Total_Fil_rewards_for_Inv_on_daily_basis = 0.0
			FMP_INFO.Daily_twenty_five_percent_Reward_for_KSL_P1 = 0.0
			FMP_INFO.Daily_seventy_five_percent_Reward_for_KSL_P1 = 0.0
			FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_KSL_P1 = Prev_day_Cumulative_TwentyFive_percent_Reward_for_KSL_P1
			FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1 = (1.0 - 1.0/180.0) * Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1
			FMP_INFO.Daily_One_Eighty_Release_for_KSL_P1 = (1.0 / 180.0) * Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1
			FMP_INFO.Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1 = Prev_day_Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1 + FMP_INFO.Daily_One_Eighty_Release_for_KSL_P1
			FMP_INFO.Fil_when_deal_is_over_KSL_p1 = Prev_day_FIl_When_deal_is_over_KSL_p1 + FMP_INFO.Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1
			FMP_INFO.Vogo_75_percent_Reward = FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward - FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv - FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1
			FMP_INFO.Vogo_25_percent_Reward = FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward - FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv - FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv - FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_KSL_P1

			FMP_INFO.Frp_Cumulative_Fil_Sum_for_Inv = Prev_day_Frp_Cumulative_Fil_Sum_for_Inv
			if date_today == 26 {
				FMP_INFO.FRP_Adj_Power_for_inv = Prev_day_FRP_Adj_Power_for_inv + (FMP_INFO.Cumulative_Total_staking_of_inv / (float32(Current_Sector_Initial_Pledge_32GB) * float32(32.0)) / float32(1000.0))
			}
			FMP_INFO.FRP_Adj_Power_for_inv = Prev_day_FRP_Adj_Power_for_inv

		}
		FMP_INFO.Fil_rewards_Daily_basis = Total_FiL_Reward_Vogo
		FMP_INFO.Cumulative_fil_Reward = Prev_day_Cumulative_fil_Reward + FMP_INFO.Fil_rewards_Daily_basis
		FMP_INFO.Daily_TwentyFive_percent_Reward = float32(25.0/100.0) * Total_FiL_Reward_Vogo
		FMP_INFO.Daily_SeventyFive_percent_Locked_Reward = float32(75.0/100.0) * Total_FiL_Reward_Vogo
		FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward = FMP_INFO.Daily_SeventyFive_percent_Locked_Reward + (Prev_day_seventyFive_percent_Locked_Reward)*float32(1.0-1.0/180)
		FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward = FMP_INFO.Daily_TwentyFive_percent_Reward + Prev_day_twentyFive_percent_Reward + float32(1.0/180.0)*Prev_day_seventyFive_percent_Locked_Reward
		FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv = float32(Total_Quality_adjP_on_daily_basis_for_Inv) / float32(Total_Quality_adjP_on_daily_basis_for_Vogo) * FMP_INFO.Daily_TwentyFive_percent_Reward
		FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv = Prev_day_Cumulative_twentyFive_percent_Reward_for_inv + FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv
		FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv = (float32(Total_Quality_adjP_on_daily_basis_for_Inv) / float32(Total_Quality_adjP_on_daily_basis_for_Vogo)) * FMP_INFO.Daily_SeventyFive_percent_Locked_Reward
		FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv = FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv + float32(1.0-1.0/180)*Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv
		FMP_INFO.Daily_Release_of_1__180_of_locked_Reward_for_inv = float32(1.0/180.0) * Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv
		FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv = Prev_day_Cumulative_Release_of_1__180_of_locked_Reward_for_inv + FMP_INFO.Daily_Release_1_180_of_SeventyFive_percent_Reward
		FMP_INFO.Daily_Staking_of_inv = FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv + FMP_INFO.Daily_Release_1_180_of_SeventyFive_percent_Reward
		FMP_INFO.Cumulative_Total_staking_of_inv = Prev_day_Cumulative_Total_staking_of_inv + FMP_INFO.Daily_Staking_of_inv
		FMP_INFO.Total_Fil_rewards_for_Inv_on_daily_basis = FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv + FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv
		FMP_INFO.Daily_twenty_five_percent_Reward_for_KSL_P1 = float32(0.3) * FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv * float32(1143.0) / float32(Total_Quality_adjP_on_daily_basis_for_Vogo)
		FMP_INFO.Daily_seventy_five_percent_Reward_for_KSL_P1 = float32(0.3) * FMP_INFO.Daily_SeventyFive_percent_Locked_Reward * float32(1143.0) / float32(Total_Quality_adjP_on_daily_basis_for_Vogo)
		FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_KSL_P1 = Prev_day_Cumulative_TwentyFive_percent_Reward_for_KSL_P1 + FMP_INFO.Daily_twenty_five_percent_Reward_for_KSL_P1
		FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1 = FMP_INFO.Daily_seventy_five_percent_Reward_for_KSL_P1 + float32(1.0-1.0/180.0)*Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1
		FMP_INFO.Daily_One_Eighty_Release_for_KSL_P1 = float32(1.0/180.0) * Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1
		FMP_INFO.Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1 = Prev_day_Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1 + FMP_INFO.Daily_One_Eighty_Release_for_KSL_P1
		FMP_INFO.Fil_when_deal_is_over_KSL_p1 = float32(10000.0-(10000.0/350.0)*50.0) + FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_KSL_P1 + FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1 +
			FMP_INFO.Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1
		FMP_INFO.Vogo_75_percent_Reward = FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward - FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv - FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1
		FMP_INFO.Vogo_25_percent_Reward = FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward - FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv - FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv - FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_KSL_P1

		if date_today == 26 {
			FMP_INFO.FRP_Adj_Power_for_inv = Prev_day_FRP_Adj_Power_for_inv + (FMP_INFO.Cumulative_Total_staking_of_inv / (float32(Current_Sector_Initial_Pledge_32GB) * float32(32.0)) / float32(1000.0))
		}
		FMP_INFO.FRP_Adj_Power_for_inv = Prev_day_FRP_Adj_Power_for_inv
	}
	FMP_INFO = prev_fmp_info
	FMP_INFO.Graduation_messages_for_inv = "YOU ARE GRADUATED FROM INVESTMENT"

	fmt.Println(FMP_INFO)

	//db.Create(&FMP_INFO)
}

// struct to get data base from the database where data were stored from the api call.

func DbConnect() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("")
	}
	dbUser := os.Getenv("USERNAME")
	dbPassword := os.Getenv("PASSWORD")
	dbIP := os.Getenv("DBIP")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")
	dbSslMode := os.Getenv("DBSSLMODE")

	dbConfig := config.Config{
		User:     dbUser,
		Password: dbPassword,
		Host:     dbIP,
		Port:     dbPort,
		DbName:   dbName,
		SslMode:  dbSslMode,
	}
	db, err := config.Connect(&dbConfig)

	if err != nil {
		fmt.Printf("There is error in connecting to data\n")
	}
	return db, nil
}

func Query_Prev_day_info_For_KSL_FRP() (Info_For_KSL_FRP_500_and_KSL_100000, error) {
	db, err := DbConnect()
	if err != nil {
		fmt.Println(err.Error())
	}

	var info_for_inv Info_For_KSL_FRP_500_and_KSL_100000
	query := "SELECT * FROM info_for_ksl_frp_500_and_ksl_100000  order by date desc limit 1;"
	db.Raw(query).Scan(&info_for_inv)
	if err != nil {
		fmt.Println(err)
	}

	return info_for_inv, nil
}

func Query_total_Quality_and_Fil_reward() (Total_Quality_adjP_and_Fil_Reward_for_Vogo_network, error) {
	db, err := DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	var Quality_adj_fil_reward Total_Quality_adjP_and_Fil_Reward_for_Vogo_network
	query := "SELECT * FROM node_info_daily_adjp_and_f_il_prices order by date desc limit 1;"
	db.Raw(query).Scan(&Quality_adj_fil_reward)
	fmt.Println(Quality_adj_fil_reward)
	return Quality_adj_fil_reward, nil
}

func Calculate_total_FIl_reward_and_total_quality_adj_P_and_Fil_Reward_for_Vogo() {
	db, err := DbConnect()
	if err != nil {
		fmt.Println(err)
	}

	query := "SELECT * FROM node_info_daily_adjp_and_f_il_prices order by date asc"
	FMP_Info := Node_Info_Daily_and_FIl_Price_{}
	rows, err := db.Raw(query).Rows()
	if err != nil {
		fmt.Printf("The error occured %v\n", err)
	}

	for rows.Next() {
		if err := db.ScanRows(rows, &FMP_Info); err != nil {
			fmt.Printf("The error occurred %v\n", err)
			return
		}
		fmt.Println(FMP_Info.Date)
		var Info_for_vogo Total_Quality_adjP_and_Fil_Reward_for_Vogo_network
		Info_for_vogo.Date = FMP_Info.Date
		Info_for_vogo.Total_FIL_Reward_Vogo_daily_Basis = FMP_Info.Fil_Rewards_f01624021_node_1 + FMP_Info.Fil_Rewards_f01918123_node_2 + FMP_Info.Fil_Rewards_f01987994_node_3
		Info_for_vogo.Total_Quality_adjP_For_Vogo_Daily_Basis = FMP_Info.FRP_f01624021_node_1_adjP + FMP_Info.FRP_f01918123_node_2_adjP + FMP_Info.FRP_f01987994_node_3_adjP
		Info_for_vogo.Current_Sector_Initial_Pledge_32GB = FMP_Info.Current_Sector_Initial_Pledge_32GB

		db.Create(&Info_for_vogo)
		defer rows.Close()
	}
}

func KSP_FRP_INFO() {
	db, err := DbConnect()
	if err != nil {
		fmt.Println(err.Error())
	}
	query := "SELECT * FROM  total_quality_adjp_and_fil_reward_for_vogo_networks  order by date asc offset 3"
	total_Quality_adjP_on_daily_basis_for_Vogo := Total_Quality_adjP_and_Fil_Reward_for_Vogo_network{}
	rows, err := db.Raw(query).Rows()
	if err != nil {
		fmt.Printf("The error occured %v\n", err)
	}
	for rows.Next() {
		defer rows.Close()
		if err := db.ScanRows(rows, &total_Quality_adjP_on_daily_basis_for_Vogo); err != nil {
			fmt.Printf("The error occured %v\n", err)
			return
		}

		fmt.Println(total_Quality_adjP_on_daily_basis_for_Vogo)
		prev_fmp_info, err := Query_Prev_day_info_For_KSL_FRP()
		if err != nil {
			fmt.Println(err.Error())
		}

		Total_FiL_Reward_Vogo := total_Quality_adjP_on_daily_basis_for_Vogo.Total_FIL_Reward_Vogo_daily_Basis
		Total_Quality_adjP_For_Vogo_Daily_Basis := total_Quality_adjP_on_daily_basis_for_Vogo.Total_Quality_adjP_For_Vogo_Daily_Basis
		fmt.Printf("total_Quality_adjP_on_daily_basis_for_Vogo %f\n", Total_Quality_adjP_For_Vogo_Daily_Basis)
		fmt.Printf("Total_fil_rewards are %f\n", Total_FiL_Reward_Vogo)
		Total_Quality_adjP_on_daily_basis_for_Inv := float32(500)
		fmt.Printf("Total_Quality_adjP_on_daily_basis_for_Inv  %f\n", Total_Quality_adjP_on_daily_basis_for_Inv)
		date_today := time.Now().Day()

		//date is fetched from the info_from_api table - the latest date is updated
		Date := total_Quality_adjP_on_daily_basis_for_Vogo.Date
		fmt.Printf("Date is %s\n", Date)

		var FMP_INFO Info_For_KSL_FRP_500_and_KSL_100000

		Prev_day_Cumulative_fil_Reward := prev_fmp_info.Cumulative_fil_Reward
		Prev_day_seventyFive_percent_Locked_Reward := prev_fmp_info.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward
		Prev_day_twentyFive_percent_Reward := prev_fmp_info.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward
		Prev_day_Cumulative_twentyFive_percent_Reward_for_inv := prev_fmp_info.Cumulative_TwentyFive_percent_Reward_for_Inv
		Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv := prev_fmp_info.Cumulative_Seventy_five_percent_locked_reward_for_inv
		Prev_day_Cumulative_Release_of_1__180_of_locked_Reward_for_inv := prev_fmp_info.Cumulative_Release_of_1__180_of_locked_Reward_for_inv
		Prev_day_Cumulative_Total_staking_of_inv := prev_fmp_info.Cumulative_Total_staking_of_inv
		Prev_day_FRP_Adj_Power_for_inv := prev_fmp_info.FRP_Adj_Power_for_inv
		Prev_day_Frp_Cumulative_Fil_Sum_for_Inv := prev_fmp_info.Frp_Cumulative_Fil_Sum_for_Inv
		Prev_day_Cumulative_TwentyFive_percent_Reward_for_KSL_P1 := prev_fmp_info.Cumulative_TwentyFive_percent_Reward_for_KSL_P1
		Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1 := prev_fmp_info.Cumulative_Seventy_five_percent_Reward_for_KSL_P1
		Prev_day_Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1 := prev_fmp_info.Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1
		Prev_day_FIl_When_deal_is_over_KSL_p1 := prev_fmp_info.Fil_when_deal_is_over_KSL_p1

		// query the Current Sector initail Pledge information
		Current_Sector_Initial_Pledge_32GB := total_Quality_adjP_on_daily_basis_for_Vogo.Current_Sector_Initial_Pledge_32GB

		FMP_INFO.Date = Date
		// Querry the FMP_Info_for_investor_updates for previous FMP_Info_for_investor_
		if prev_fmp_info.FRP_Adj_Power_for_inv < 1500.0 {

			FMP_INFO.Graduation_messages_for_inv = "YOUR STILL YET TO GRADAUTE"

			if Total_FiL_Reward_Vogo == 0 {
				FMP_INFO.Total_Quality_adjP_For_Vogo_Daily_Basis = Total_Quality_adjP_For_Vogo_Daily_Basis
				fmt.Println(FMP_INFO.Total_Quality_adjP_For_Vogo_Daily_Basis)
				FMP_INFO.Fil_rewards_Daily_basis = 0
				FMP_INFO.Cumulative_fil_Reward = Prev_day_Cumulative_fil_Reward
				FMP_INFO.Daily_TwentyFive_percent_Reward = 0.0
				FMP_INFO.Daily_SeventyFive_percent_Locked_Reward = 0.0
				FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward = Prev_day_seventyFive_percent_Locked_Reward * (1.0 - 1.0/180.0)
				FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward = Prev_day_twentyFive_percent_Reward + (1.0/180.0)*Prev_day_seventyFive_percent_Locked_Reward
				FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv = 0.0
				FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv = Prev_day_Cumulative_twentyFive_percent_Reward_for_inv

				FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv = 0.0

				FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv = Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv * (1.0 - 1.0/180.0)

				FMP_INFO.Daily_Release_of_1__180_of_locked_Reward_for_inv = (1.0 / 180.0) * Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv

				FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv = Prev_day_Cumulative_Release_of_1__180_of_locked_Reward_for_inv + FMP_INFO.Daily_Release_1_180_of_SeventyFive_percent_Reward

				FMP_INFO.Daily_Staking_of_inv = FMP_INFO.Daily_Release_of_1__180_of_locked_Reward_for_inv

				FMP_INFO.Cumulative_Total_staking_of_inv = Prev_day_Cumulative_Total_staking_of_inv + FMP_INFO.Daily_Staking_of_inv

				FMP_INFO.Total_Fil_rewards_for_Inv_on_daily_basis = 0.0

				FMP_INFO.Daily_twenty_five_percent_Reward_for_KSL_P1 = 0.0

				FMP_INFO.Daily_seventy_five_percent_Reward_for_KSL_P1 = 0.0

				FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_KSL_P1 = Prev_day_Cumulative_TwentyFive_percent_Reward_for_KSL_P1

				FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1 = float32(1.0-1.0/180.0) * Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1

				if Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1 == 0 {
					FMP_INFO.Daily_One_Eighty_Release_for_KSL_P1 = 0.0
				} else {
					FMP_INFO.Daily_One_Eighty_Release_for_KSL_P1 = (1.0 / 180.0) * Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1
				}
				FMP_INFO.Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1 = Prev_day_Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1 + FMP_INFO.Daily_One_Eighty_Release_for_KSL_P1
				FMP_INFO.Fil_when_deal_is_over_KSL_p1 = Prev_day_FIl_When_deal_is_over_KSL_p1 + FMP_INFO.Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1
				FMP_INFO.Vogo_75_percent_Reward = FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward - FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv - FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1
				FMP_INFO.Vogo_25_percent_Reward = FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward - FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv - FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv - FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_KSL_P1

				FMP_INFO.Frp_Cumulative_Fil_Sum_for_Inv = Prev_day_Frp_Cumulative_Fil_Sum_for_Inv
				if date_today == 26 {
					FMP_INFO.FRP_Adj_Power_for_inv = Prev_day_FRP_Adj_Power_for_inv + (FMP_INFO.Cumulative_Total_staking_of_inv / (float32(Current_Sector_Initial_Pledge_32GB) * 32.0) / 1000.0)
				}
				FMP_INFO.FRP_Adj_Power_for_inv = Prev_day_FRP_Adj_Power_for_inv

			} else {
				FMP_INFO.Fil_rewards_Daily_basis = Total_FiL_Reward_Vogo
				FMP_INFO.Total_Quality_adjP_For_Vogo_Daily_Basis = Total_Quality_adjP_For_Vogo_Daily_Basis
				FMP_INFO.Cumulative_fil_Reward = Prev_day_Cumulative_fil_Reward + FMP_INFO.Fil_rewards_Daily_basis
				FMP_INFO.Daily_TwentyFive_percent_Reward = float32(25.0/100.0) * Total_FiL_Reward_Vogo
				FMP_INFO.Daily_SeventyFive_percent_Locked_Reward = float32(75.0/100.0) * Total_FiL_Reward_Vogo

				if Prev_day_seventyFive_percent_Locked_Reward == 0 {
					FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward = FMP_INFO.Daily_TwentyFive_percent_Reward
					FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward = FMP_INFO.Daily_SeventyFive_percent_Locked_Reward
				} else {
					FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward = FMP_INFO.Daily_SeventyFive_percent_Locked_Reward + (Prev_day_seventyFive_percent_Locked_Reward)*float32(1.0-1.0/180)
					if Prev_day_twentyFive_percent_Reward == 0 {
						FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward = FMP_INFO.Daily_TwentyFive_percent_Reward
					} else {
						FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward = FMP_INFO.Daily_TwentyFive_percent_Reward + (Prev_day_twentyFive_percent_Reward)*float32(1.0-1.0/180)
					}
				}
				FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv = Total_Quality_adjP_on_daily_basis_for_Inv / Total_Quality_adjP_For_Vogo_Daily_Basis * FMP_INFO.Daily_TwentyFive_percent_Reward
				FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv = Prev_day_Cumulative_twentyFive_percent_Reward_for_inv + FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv
				FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv = (Total_Quality_adjP_on_daily_basis_for_Inv / Total_Quality_adjP_For_Vogo_Daily_Basis) * FMP_INFO.Daily_SeventyFive_percent_Locked_Reward

				if Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv == 0 {
					FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv = FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv
					FMP_INFO.Daily_Release_of_1__180_of_locked_Reward_for_inv = 0
					FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv = FMP_INFO.Daily_Release_1_180_of_SeventyFive_percent_Reward
				} else {
					FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv = FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv + float32(1.0-1.0/180)*Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv
					FMP_INFO.Daily_Release_of_1__180_of_locked_Reward_for_inv = float32(1.0/180.0) * Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv
					FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv = Prev_day_Cumulative_Release_of_1__180_of_locked_Reward_for_inv + FMP_INFO.Daily_Release_1_180_of_SeventyFive_percent_Reward
				}

				//FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv = Prev_day_Cumulative_Release_of_1__180_of_locked_Reward_for_inv + FMP_INFO.Daily_Release_1_180_of_SeventyFive_percent_Reward
				FMP_INFO.Daily_Staking_of_inv = FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv + FMP_INFO.Daily_Release_1_180_of_SeventyFive_percent_Reward
				FMP_INFO.Cumulative_Total_staking_of_inv = Prev_day_Cumulative_Total_staking_of_inv + FMP_INFO.Daily_Staking_of_inv
				FMP_INFO.Total_Fil_rewards_for_Inv_on_daily_basis = FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv + FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv
				FMP_INFO.Frp_Cumulative_Fil_Sum_for_Inv = Prev_day_Frp_Cumulative_Fil_Sum_for_Inv + FMP_INFO.Total_Fil_rewards_for_Inv_on_daily_basis

				if Date <= "2023-2-26" {
					FMP_INFO.Daily_twenty_five_percent_Reward_for_KSL_P1 = 0
					FMP_INFO.Daily_seventy_five_percent_Reward_for_KSL_P1 = 0
					FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_KSL_P1 = 0
					FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1 = 0
					FMP_INFO.Daily_One_Eighty_Release_for_KSL_P1 = 0
					FMP_INFO.Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1 = 0
					FMP_INFO.Fil_when_deal_is_over_KSL_p1 = 0
					FMP_INFO.Vogo_75_percent_Reward = FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward - FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv
					FMP_INFO.Vogo_25_percent_Reward = FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward - FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv - FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv
				} else {
					FMP_INFO.Daily_twenty_five_percent_Reward_for_KSL_P1 = float32(0.3) * FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv * (float32(1143) / float32(Total_Quality_adjP_on_daily_basis_for_Vogo))
					FMP_INFO.Daily_seventy_five_percent_Reward_for_KSL_P1 = float32(0.3) * FMP_INFO.Daily_SeventyFive_percent_Locked_Reward * (float32(1143) / float32(Total_Quality_adjP_on_daily_basis_for_Vogo))
					FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_KSL_P1 = Prev_day_Cumulative_TwentyFive_percent_Reward_for_KSL_P1 + FMP_INFO.Daily_twenty_five_percent_Reward_for_KSL_P1
					if Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1 == 0 {
						FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1 = FMP_INFO.Daily_seventy_five_percent_Reward_for_KSL_P1
						FMP_INFO.Daily_One_Eighty_Release_for_KSL_P1 = float32(0)
					} else {
						FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1 = FMP_INFO.Daily_seventy_five_percent_Reward_for_KSL_P1 + float32(1.0-1.0/180.0)*Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1
						FMP_INFO.Daily_One_Eighty_Release_for_KSL_P1 = float32(1.0/180.0) * Prev_day_Cumulative_Seventy_five_percent_Reward_for_KSL_P1
					}
					FMP_INFO.Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1 = Prev_day_Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1 + FMP_INFO.Daily_One_Eighty_Release_for_KSL_P1
					FMP_INFO.Fil_when_deal_is_over_KSL_p1 = float32(10000.0-(10000.0/350.0)*50.0) + FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_KSL_P1 + FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1 +
						FMP_INFO.Cumulative_of_Daily_One_Eighty_Release_for_KSL_P1
					FMP_INFO.Vogo_75_percent_Reward = FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward - FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv - FMP_INFO.Cumulative_Seventy_five_percent_Reward_for_KSL_P1
					FMP_INFO.Vogo_25_percent_Reward = FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward - FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv - FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv - FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_KSL_P1
				}
				if date_today == 26 {
					FMP_INFO.FRP_Adj_Power_for_inv = Prev_day_FRP_Adj_Power_for_inv + (FMP_INFO.Cumulative_Total_staking_of_inv / (float32(Current_Sector_Initial_Pledge_32GB) * 32.0) / float32(1000.0))
				} else {
					FMP_INFO.FRP_Adj_Power_for_inv = float32(500.0)
				}

			}

		}
		db.Create(&FMP_INFO)

	}
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
    daily_staking_of_inv ,
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
    '2022/11/26',1922,19.51,4.88, 14.63, 19.51, 0, 4.88,14.63, 3.61,10.83,1.27,1.27,3.81, 3.81, 0.0,0.0, 1.27,1.27, 500, 5.07, 5.07,0,0, 0,0,0,0,0,0,0,0,0,'YOUR STILL YET TO GRADAUTE' );

*/
