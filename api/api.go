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

type Node_Info_Daily_and_FIl_Price struct {
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

//Get FIL_Rewards and Quality adjusted power of node f01624021 on daily basis

func FIL_Price_n_Block_rewards_for_Each_Node() {
	var wg sync.WaitGroup
	db := DbConnect()
	c := make(chan Node_Info_Daily_and_FIl_Price, 1)

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
		c <- Node_Info_Daily_and_FIl_Price{Fil_Price: FIL_PRICE}

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
	Node_Info_Daily_and_FIl_Price := Node_Info_Daily_and_FIl_Price{}
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

func FMP_investment_Calculate() {

	// calculated on the day of investement  i.e will remain consant from the day of investment till the 25th of the next month and on 25th at 12.00 am it wil be updated.
	// i.e it is  updated only once a month.
	//Total_Quality_adjP_on_daily_basis:=0

	db := DbConnect()

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
	fmt.Println(FMP_INFO.Total_Quality_adjP_For_Inv_daily_Basis)

	//calculate the time
	now := time.Now()
	//initiliaze the instance of struct for FMP_INFO which is the main struct where all the data will be stored.
	FMP_Info := &FMP_Info_for_investor{}

	// Checks if the date is 25th of the month and time is 0.00 am
	if now.Day() == 26 {
		// Since Node_info is updated once everyday at
		total_Quality_adjP_For_Vogo := Node_info.FRP_f01624021_node_1_adjP + Node_info.FRP_f01918123_node_2_adjP + Node_info.FRP_f01987994_node_3_adjP
		FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis = float32(total_Quality_adjP_For_Vogo)
		FMP_Info.Total_Quality_adjP_For_Inv_daily_Basis = FMP_INFO.Total_Quality_adjP_with_increased_FRP_inv
	} else {

		FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis = FMP_INFO.Total_Quality_adjP_For_Vogo_Daily_Basis
		fmt.Printf("Total_Quaity_adjP_For_Vogo_Daily_Basis\n")
		FMP_Info.Total_Quality_adjP_For_Inv_daily_Basis = FMP_INFO.Total_Quality_adjP_For_Inv_daily_Basis
	}

	// query from thd node_info_daily_f_il_price

	Total_FIL_Reward_Vogo_daily_Basis := Node_info.Fil_Rewards_f01624021_node_1 + Node_info.Fil_Rewards_f01918123_node_2 + Node_info.Fil_Rewards_f01987994_node_3

	FMP_Info.Total_FIL_Reward_Vogo_daily_Basis = Total_FIL_Reward_Vogo_daily_Basis
	fmt.Printf("FMP_Info.Total_FIL_Reward_Vogo_daily_Basis %f\n", FMP_Info.Total_FIL_Reward_Vogo_daily_Basis)

	total_Quality_adjP_For_Inv := FMP_Info.Total_Quality_adjP_For_Inv_daily_Basis

	fmt.Printf("total_Quality_adjP_For_Inv %f\n", total_Quality_adjP_For_Inv)
	/*500 Tib will be on the day of investment and wll be used to calculate the increased FRMo the date of investement
	till the 25 th of the each moent
	*/

	if FMP_Info.Total_FIL_Reward_Vogo_daily_Basis != 0 {

		fmt.Printf("FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis %f\n", FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis)
		fmt.Printf("Total_FIl_Reward_Vogo_daily_basis %f\n", Total_FIL_Reward_Vogo_daily_Basis)
		fmt.Printf("totalQaulity_adjP_For_Inv %f\n", total_Quality_adjP_For_Inv)
		total_quality_adjp_vogo := FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis
		fil_rewards_for_inv := float32(total_Quality_adjP_For_Inv*FMP_Info.Total_FIL_Reward_Vogo_daily_Basis) / total_quality_adjp_vogo

		fmt.Printf("Fil_rewards_for_inv: %f\n", fil_rewards_for_inv)

		FMP_Info.Fil_Rewards_on_daily_basis_for_inv = float32(fil_rewards_for_inv)
		fmt.Printf("FMP_Info.Fil_Rewards_on_daily_basis_for_inv %f\n", FMP_Info.Fil_Rewards_on_daily_basis_for_inv)
	} else {

		FMP_Info.Fil_Rewards_on_daily_basis_for_inv = 0.0
	}

	Staking_on_daily_basis := FMP_Info.Fil_Rewards_on_daily_basis_for_inv

	FMP_Info.Staking_on_daily_basis = Staking_on_daily_basis

	Sector_initial_pledge := Node_info.Current_Sector_Initial_Pledge_32GB

	if Staking_on_daily_basis == 0 {

		FMP_Info.Increased_FRP_on_daily_basis = 0

	} else {

		FMP_Info.Increased_FRP_on_daily_basis = (float32(Staking_on_daily_basis) * (32.0) / float32(Sector_initial_pledge)) / 1000.0
		fmt.Printf("FMP_Info.Increased_FRP_on_daily_basis %f\n", FMP_Info.Increased_FRP_on_daily_basis)
	}

	prevTotalFILRewards := FMP_INFO.Total_FIL_Rewards
	prevTotalFRP := FMP_INFO.Total_FRP
	PrevTotalStaking := FMP_INFO.Total_Staking
	prevTotal_Quality_AdjPow_inv := FMP_INFO.Total_Quality_adjP_with_increased_FRP_inv

	FMP_Info.Total_Staking = PrevTotalStaking + FMP_Info.Staking_on_daily_basis
	FMP_Info.Total_FIL_Rewards = prevTotalFILRewards + FMP_Info.Fil_Rewards_on_daily_basis_for_inv
	FMP_Info.Total_FRP = prevTotalFRP + FMP_Info.Increased_FRP_on_daily_basis
	FMP_Info.Total_Quality_adjP_with_increased_FRP_inv = float32(FMP_Info.Increased_FRP_on_daily_basis + float32(prevTotal_Quality_AdjPow_inv))
	FMP_Info.Date = Node_info.Date
	FMP_Info.Total_Reward_value = float32(float32(FMP_Info.Total_FIL_Rewards) * Node_info.Fil_Price)

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

func Query_Fmp_table(db *gorm.DB) FMP_Info_for_investor_ {
	query := " SELECT * from  fmp_info_for_investors ORDER BY DATE ASC LIMIT 1 OFFSET 28"
	var FMP_Info FMP_Info_for_investor_
	err := db.Raw(query).Scan(&FMP_Info).Error
	if err != nil {
		fmt.Printf("The FMP_information cannot be fetched %s\n", err)
	}
	fmt.Println(FMP_Info)
	return FMP_Info
}

type FMP_Info_for_investor_ struct {
	Date                                      string
	Total_Quality_adjP_For_Vogo_Daily_Basis   float32
	Total_FIL_Reward_Vogo_daily_Basis         float32
	Total_Quality_adjP_For_Inv_daily_Basis    float32
	Total_Quality_adjP_with_increased_FRP_inv float32
	Fil_Rewards_on_daily_basis_for_inv        float32
	Total_Fil_rewards_for_Inv                 float32
	Total_FIL_Rewards                         float32
	Staking_on_daily_basis                    float32
	Total_Staking                             float32
	Total_Reward_value                        float32
	Increased_FRP_on_daily_basis              float32
	Total_FRP                                 float32
	Paid_Reward_to_Investor                   float32
	Total_FIL_Paid_to_Investor                float32
	Value_of_FIL_Paid_to_Investor             float32
	Value_of_Total_FIl_Paid                   float32
}
type FMP_Info_for_investor_updates struct {
	Date                                                                   string
	Total_Quality_adjP_For_Vogo_Daily_Basis                                float32
	Fil_rewards_Daily_basis                                                float32
	Daily_TwentyFive_percent_Reward                                        float32
	Daily_SeventyFive_percent_Locked_Reward                                float32
	Cumulative_fil_Reward                                                  float32
	Daily_Release_1_180_of_SeventyFive_percent_Reward                      float32
	Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward          float32
	Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward float32
	Vogo_25_percent_Reward                                                 float32
	Vogo_75_percent_Reward                                                 float32
	Daily_TwentyFive_percent_Reward_for_inv                                float32
	Cumulative_TwentyFive_percent_Reward_for_Inv                           float32
	Daily_Seventy_five_percent_locked_reward_for_inv                       float32
	Cumulative_Seventy_five_percent_locked_reward_for_inv                  float32
	Daily_Release_of_1__180_of_locked_Reward_for_inv                       float32
	Cumulative_Release_of_1__180_of_locked_Reward_for_inv                  float32
	Daily_Staking_of_inv                                                   float32
	Cumulative_Total_staking_of_inv                                        float32
	FRP_Adj_Power_for_inv                                                  float32
	Total_Fil_rewards_for_Inv_on_daily_basis                               float32
	Frp_Cumulative_Fil_Sum_for_Inv                                         float32
	graduation_messages_for_inv                                            string
}

//FRP 투자계정 (KSL_FRP_500) 현황
func Calculate_KSL_FRP_500() {
	db := DbConnect()
	Info_from_api := Query_Api_table()
	prev_fmp_info := Query_Fmp_table_update()
	Total_FiL_Reward_Vogo := float32(Info_from_api.Total_FIL_Reward_Vogo_daily_Basis)
	Total_Quality_adjP_on_daily_basis_for_Inv := Info_from_api.Total_Quality_adjP_For_Vogo_Daily_Basis
	date_today := time.Now().Day()

	Date := Info_from_api.Date
	var FMP_INFO FMP_Info_for_investor_updates

	Prev_day_Cumulative_fil_Reward := prev_fmp_info.Cumulative_fil_Reward
	Prev_day_seventyFive_percent_Locked_Reward := prev_fmp_info.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward
	Prev_day_twentyFive_percent_Reward := prev_fmp_info.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward
	Prev_day_Cumulative_twentyFive_percent_Reward_for_inv := prev_fmp_info.Cumulative_TwentyFive_percent_Reward_for_Inv
	Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv := prev_fmp_info.Cumulative_Seventy_five_percent_locked_reward_for_inv
	Prev_day_Cumulative_Release_of_1__180_of_locked_Reward_for_inv := prev_fmp_info.Cumulative_Release_of_1__180_of_locked_Reward_for_inv
	Prev_day_Cumulative_Total_staking_of_inv := prev_fmp_info.Cumulative_Total_staking_of_inv
	Prev_day_FRP_Adj_Power_for_inv := prev_fmp_info.FRP_Adj_Power_for_inv
	Prev_day_Frp_Cumulative_Fil_Sum_for_Inv := prev_fmp_info.Frp_Cumulative_Fil_Sum_for_Inv
	Current_Sector_Initial_Pledge_32GB := 0.2015
	FMP_INFO.Date = Date
	// Querry the FMP_Info_for_investor_updates for previous FMP_Info_for_investor_
	if prev_fmp_info.FRP_Adj_Power_for_inv < 1500.0 {
		FMP_INFO.graduation_messages_for_inv = "YOUR STILL YET TO GRADAUTE"
		if Total_FiL_Reward_Vogo == 0 {
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
			FMP_INFO.Vogo_75_percent_Reward = FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward - FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv
			FMP_INFO.Vogo_25_percent_Reward = FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward - (FMP_INFO.Cumulative_Total_staking_of_inv - FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv)
			FMP_INFO.Total_Fil_rewards_for_Inv_on_daily_basis = 0.0
			FMP_INFO.Frp_Cumulative_Fil_Sum_for_Inv = Prev_day_Frp_Cumulative_Fil_Sum_for_Inv
			if date_today == 26 {
				FMP_INFO.FRP_Adj_Power_for_inv = Prev_day_FRP_Adj_Power_for_inv + (FMP_INFO.Cumulative_Total_staking_of_inv / (float32(Current_Sector_Initial_Pledge_32GB) * 32.0) / 1000.0)
			}
			FMP_INFO.FRP_Adj_Power_for_inv = Prev_day_FRP_Adj_Power_for_inv

		}
		FMP_INFO.Fil_rewards_Daily_basis = Total_FiL_Reward_Vogo
		FMP_INFO.Cumulative_fil_Reward = Prev_day_Cumulative_fil_Reward + FMP_INFO.Fil_rewards_Daily_basis
		FMP_INFO.Daily_TwentyFive_percent_Reward = (25.0 / 100.0) * Total_FiL_Reward_Vogo
		FMP_INFO.Daily_SeventyFive_percent_Locked_Reward = (75.0 / 100.0) * Total_FiL_Reward_Vogo
		FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward = FMP_INFO.Daily_SeventyFive_percent_Locked_Reward + (Prev_day_seventyFive_percent_Locked_Reward)*(1.0-1.0/180)
		FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward = FMP_INFO.Daily_TwentyFive_percent_Reward + Prev_day_twentyFive_percent_Reward + (1.0/180.0)*Prev_day_seventyFive_percent_Locked_Reward
		FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv = float32(Total_Quality_adjP_on_daily_basis_for_Inv) / float32(Total_Quality_adjP_on_daily_basis_for_Vogo) * FMP_INFO.Daily_TwentyFive_percent_Reward
		FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv = Prev_day_Cumulative_twentyFive_percent_Reward_for_inv + FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv
		FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv = (float32(Total_Quality_adjP_on_daily_basis_for_Inv) / float32(Total_Quality_adjP_on_daily_basis_for_Vogo)) * FMP_INFO.Daily_SeventyFive_percent_Locked_Reward
		FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv = FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv + (1.0-1.0/180)*Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv
		FMP_INFO.Daily_Release_of_1__180_of_locked_Reward_for_inv = (1.0 / 180.0) * Prev_day_Cumulative_Seventy_five_percent_locked_reward_for_inv
		FMP_INFO.Cumulative_Release_of_1__180_of_locked_Reward_for_inv = Prev_day_Cumulative_Release_of_1__180_of_locked_Reward_for_inv + FMP_INFO.Daily_Release_1_180_of_SeventyFive_percent_Reward
		FMP_INFO.Daily_Staking_of_inv = FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv + FMP_INFO.Daily_Release_1_180_of_SeventyFive_percent_Reward
		FMP_INFO.Cumulative_Total_staking_of_inv = Prev_day_Cumulative_Total_staking_of_inv + FMP_INFO.Daily_Staking_of_inv
		FMP_INFO.Vogo_75_percent_Reward = FMP_INFO.Cumulative_SeventyFive_percent_Locked_Reward_minus_1_180_locked_Reward - FMP_INFO.Cumulative_Seventy_five_percent_locked_reward_for_inv
		FMP_INFO.Vogo_25_percent_Reward = FMP_INFO.Cumulative_TwentyFive_percent_Reward_plus_1_180_locked_Reward - (FMP_INFO.Cumulative_Total_staking_of_inv - FMP_INFO.Cumulative_TwentyFive_percent_Reward_for_Inv)
		FMP_INFO.Total_Fil_rewards_for_Inv_on_daily_basis = FMP_INFO.Daily_TwentyFive_percent_Reward_for_inv + FMP_INFO.Daily_Seventy_five_percent_locked_reward_for_inv

		if date_today == 26 {
			FMP_INFO.FRP_Adj_Power_for_inv = Prev_day_FRP_Adj_Power_for_inv + (FMP_INFO.Cumulative_Total_staking_of_inv / (float32(Current_Sector_Initial_Pledge_32GB) * 32.0) / 1000.0)
		}
		FMP_INFO.FRP_Adj_Power_for_inv = Prev_day_FRP_Adj_Power_for_inv
	}
	FMP_INFO = prev_fmp_info
	FMP_INFO.graduation_messages_for_inv = "YOU ARE GRADUATED FROM INVESTMENT"

	db.Create(&FMP_INFO)
}

// struct to get data base from the database where data were stored from the api call.

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

func DbConnect() *gorm.DB {
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
	return db
}

func Query_Fmp_table_update() FMP_Info_for_investor_updates {
	db := DbConnect()

	var Fmp_table_update FMP_Info_for_investor_updates
	query := "SELECT * FROM FMP_TABLE_update order by date desc limit 1;"
	err := db.Raw(query).Scan(&Fmp_table_update)
	if err != nil {
		fmt.Println(err)
	}
	return Fmp_table_update
}

func Query_Api_table() FMP_Info_for_investor {
	db := DbConnect()

	var Fmp_table FMP_Info_for_investor
	query := "SELECT * FROM FMP_TABLE order by date desc limit 1;"
	err := db.Raw(query).Scan(&Fmp_table)
	if err != nil {
		fmt.Println(err)
	}
	return Fmp_table
}
