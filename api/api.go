package api 
 import (
	"fmt"
	"time"
	"net/http"
	_"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_"io/ioutil"
	_"gorm.io/gorm"
	"encoding/json"
	"sync"
	_"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_"github.com/robfig/cron"
	"github.com/achnir97/go_lang_filbytes/config"
	"github.com/joho/godotenv"
	"strconv"
	"os"
)
const ATTOFILL=1e-18  // 1 FIL is equivalent to 1e^18 
const BYTES=1e-12 // 1 terrabytes is equivalent to 1e^12 bytes 

// struct to store each node information from api 
type MinerDetails struct {
   QualityAdjPower  string `json:"qualityAdjPower"`
   NetworkRawBytePower string `json:"networkRawBytePower"`
   BlocksMined string `json:"blocksMined"`
   WeightedBlocksMined string `json:"weightedBlocksMined"`
   TotalRewards  string `json:"totalRewards"`
}
 
var db *gorm.DB // declaring gorm model 

// declaring the global variable to know the total_quality adjustedP_on_daily_basis. 
var  Total_Quality_adjP_on_daily_basis_for_Vogo int
var  Total_Quality_adjP_on_daily_basis_for_Inv  int

type Fetched_Info struct{
	Id string `json:"id"`
	Miner MinerDetails `json:"miner"`
 }

// struct to store the price of the filecoin. 
type Fil_price struct{
	Filecoin Price `json:"filecoin"`
	}
// struct to get the price of the filecoin in KRW from the api call. 
type Price struct{
	Krw float32 `json:"krw"`
}

type  Node_Related_Info struct{
	Id string `json:"id"`
	RobustAddress string `json:"robust"`
	Actor string `json:"actor"`
	CreateHeight int `json:"createHeight"`
	CreateTimestamp int `json:"createTimestamp"`
	LastSeenHeight int  `json:"lastSeenHeight"`
	Balance string `json:"balance"`
	MessageCount int `json:"messageCount"`
	Timestamp int `json:"timestamp"`
	Miner struct {
		Owner struct {
			Address string `json:"address"`
			Balance string `json:"balance"`
		} `json:"owner"`
		Worker struct {
			Address string `json:"address"`
			Balance string `json:"balance"`
		} `json:"worker"`
		ControlAddresses[] struct{
			Address string `json:"address"`
			Balance string `json:"balance"`
		} `json:"controlAddresses"`
		PeerId string `json:"peerId"`
		MultiAddresses []string `json:"multiAddresses"`
		SectorSize int `json:"sectorSize"`
		RawBytePower string `json:"rawBytePower"`
		QualityAdjPower string `json:"qualityAdjPower"`
		NetworkQualityAdjPower string `json:"networkQualityAdjPower"`
		BlocksMined int `json:"blocksMined"`
		WeightedBlocksMined int `json:"weightedBlocksMined"`
		TotalRewards string `json:"totalRewards"`
		Sectors struct{
			Live int `json:"live"`
			Active int `json:"active"`
			Faulty int `json:"faulty"`
			Recovering int `json:"recovering"`
		} `json:"sectors"`
		PreCommitDeposits string `json:"PreCommitDeposits"`
		VestingFunds	string `json:"vestingFunds"`
		IntialPledgeRequirement string `json:"initialPledgeRequirement"`
		AvailableBalance string `json:"availableBalance"`
		SecotorPledgeBalance string `json:"sectorPledgeBalance"`
		PledgeBalance 	string 	`json:"pledgeBalance"`
		RawBytePowerRank int `json:"rawBytePowerRank"`
		QualityAdjPowerRank int `json:"qualityAdjPowerRank"`
	}  `json:"miner"`
	
	OwnerMiners []interface{} `json:"ownedMiners"`

	WorkerMiner []interface{} `json:"workerMiners"`
	Address string `json:"address"`
}



type Node_Info_Daily_and_FIl_Price struct{ 
	gorm.Model
	Date  string 
	Fil_Price float32 
    Current_Sector_Initial_Pledge_32GB float32 
	Fil_Rewards_f01624021_node_1 int64
	Fil_Rewards_f01918123_node_2 int64
	Fil_Rewards_f01987994_node_3 int64
	Cummulative_Fil_Rewards_f01624021_node_1 int64
	Cummulative_Fil_Rewards_f01918123_node_2 int64
	Cummulative_Fil_Rewards_f01987994_node_3 int64
	FRP_f01624021_node_1_adjP int64
	FRP_f01918123_node_2_adjP int64
	FRP_f01987994_node_3_adjP int64
}


//Method to assign values to FMP_integrated_info
func(Miner *Node_Info_Daily_and_FIl_Price)Set_Miner_Info(Fil_Price float32, Fil_Rewards_f01624021_node_1 int64, Fil_Rewards_f01918123_node_2 int64,
	Fil_Rewards_f01987994_node_3 int64, FRP_f01624021_node_1 int64, FRP_f01918123_node_2 int64, 
	FRP_f01987994_node_3 int64){
		Miner.Fil_Price=Fil_Price
		Miner.Fil_Rewards_f01624021_node_1=Fil_Rewards_f01624021_node_1
		Miner.Fil_Rewards_f01918123_node_2=Fil_Rewards_f01918123_node_2
		Miner.Fil_Rewards_f01987994_node_3=Fil_Rewards_f01987994_node_3
		Miner.FRP_f01624021_node_1_adjP=FRP_f01624021_node_1
		Miner.FRP_f01918123_node_2_adjP=FRP_f01918123_node_2
		Miner.FRP_f01987994_node_3_adjP=FRP_f01987994_node_3
}

//Get FIL_Rewards and Quality adjusted power of node f01624021 on daily basis 
func FIL_Price_n_Block_rewards_for_Each_Node(context *fiber.Ctx)error{ 
    var wg sync.WaitGroup
    db:=DbConnect()
	c:=make(chan Node_Info_Daily_and_FIl_Price,1)

	//var miner Node_Info_Daily_and_FIl_Price 
	var FIL_PRICE float32
	var FIL_REWARDS_f01624021_node_1  int64 
	var FIL_REWARDS_f01918123_node_2  int64
	var FIL_REWARDS_f01987994_node_3 int64
	/*var QualityAdjPower_f01624021_node_1 int
	var QualityAdjPower_f01918123_node_2 int
	var QualityAdjPower_f01987994_node_3 int */

    // get the price of FILCOIN on daily basis from the coingecko api. 
    wg.Add(1)
	go func() {
	response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=Filecoin&vs_currencies=KRW")
	if err!=nil{
			fmt.Printf("Thee http request failed with erro %s\n", err)
			context.Status(http.StatusBadRequest).JSON(&fiber.Map{
				"Message":"Couldnot fetch FIL_Price"})
	}
	defer response.Body.Close()
	var FilPrice Fil_price
	if err:=json.NewDecoder(response.Body).Decode(&FilPrice);err!=nil{
		return 
	}
	defer wg.Done()
	FIL_PRICE=FilPrice.Filecoin.Krw
	c<-Node_Info_Daily_and_FIl_Price{Fil_Price:FIL_PRICE}
	

	fmt.Printf("The Price of Filecoin is %f\n", FIL_PRICE)
	}()

// calculate Adjusted power and blocks reward for f01624021 from filfox.info/v1/api
	wg.Add(1)
	go func() {
		response, err:=http.Get("https://filfox.info/api/v1/address/f01624021")
		if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch Total_Rewards_For_Each_Node"})
		return 
	}
	defer response.Body.Close()
	var Miner_Info_f01624021   Node_Related_Info
	if err:=json.NewDecoder(response.Body).Decode(&Miner_Info_f01624021);err!=nil{
		return 
	}
	defer wg.Done()



	var Cummulative_Fil_Rewards_f01624021_node_1 int64
	
	query:="SELECT cummulative_fil_rewards_f01624021_node_1 FROM node_info_daily_and_f_il_prices ORDER BY id DESC LIMIT 1"
	
	err = db.Raw(query).Scan(&Cummulative_Fil_Rewards_f01624021_node_1).Error
	if err!=nil {
		fmt.Printf("Fil_rewards cannot be fetched, check your Error, %s\n", err)
	}
	prevCummulative_fil_rewards_for_node_1:=Cummulative_Fil_Rewards_f01624021_node_1
	fmt.Printf("The pervious cumulative_fil_rewards_for_node_1 is %d\n",prevCummulative_fil_rewards_for_node_1)
	
	f, err:=strconv.ParseFloat(Miner_Info_f01624021.Miner.TotalRewards,64) 
	/*Information about the total rewards is string type so we convertinto float */
	if err !=nil {
		fmt.Printf("Total Rewards Cannot be converted into integer, Check your error %s\n", err)
		return
	}
	/*
	total Rewards are then converted into int64 after the conversion into FILL
	Since 1FIL = 1e^18
	*/
	latestCummulative_fil_rewards_for_node_1:=int64(f*1e-18) 


	fmt.Printf("latestCummulative_fil_rewards_for_node_1, %d\n", latestCummulative_fil_rewards_for_node_1)

	if latestCummulative_fil_rewards_for_node_1 > prevCummulative_fil_rewards_for_node_1{
		FIL_REWARDS_f01624021_node_1=latestCummulative_fil_rewards_for_node_1-prevCummulative_fil_rewards_for_node_1
	}else {
		FIL_REWARDS_f01624021_node_1=0 
		fmt.Printf("you have no fil_rewards in last 24 hours")
	}

	 QualityAdjPower_f01624021_node_1, err :=strconv.ParseInt(Miner_Info_f01624021.Miner.QualityAdjPower, 10,64) 
		
	 if err !=nil {
	 	fmt.Printf("Quality adjusted power_cannot be converted into int")
	 }
	fmt.Printf("the sector balance for node are %s\n", Miner_Info_f01624021.Miner.SecotorPledgeBalance)
	//QualityAdjPower_f01624021_node_1=QualityAdjPower_f01624021_node_1
	Node_info:=<-c // some other go routines has send teh data and is assingin to nod_info
	Node_info.Fil_Rewards_f01624021_node_1=FIL_REWARDS_f01624021_node_1
	Node_info.FRP_f01624021_node_1_adjP =QualityAdjPower_f01624021_node_1
    Node_info.Cummulative_Fil_Rewards_f01624021_node_1=latestCummulative_fil_rewards_for_node_1
    
	c<-Node_info

	fmt.Printf("Miner Id : %s\n", Miner_Info_f01624021.Id)
	fmt.Printf("The total_qualityAdj for the node_f01624021 is %d\n",QualityAdjPower_f01624021_node_1)
	fmt.Printf("The total_blocks mined for the node_f01624021 are %d\n",FIL_REWARDS_f01624021_node_1)
	return 
	}()

// calculate Adjusted power and blocks reward fror f01819003
	wg.Add(1)
	go func() {
		response, err:=http.Get("https://filfox.info/api/v1/address/f01918123")
		if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch Total_Rewards_For_Each_Node"})
		return 
	}
	defer response.Body.Close()
	var Miner_Info_f01918123  Node_Related_Info
	if err:=json.NewDecoder(response.Body).Decode(&Miner_Info_f01918123);err!=nil{
		return 
	}
	defer wg.Done()
	var Cummulative_Fil_Rewards_f01918123_node_2 int64
	
	query:="SELECT Cummulative_Fil_Rewards_f01918123_node_2 FROM node_info_daily_and_f_il_prices ORDER BY id DESC LIMIT 1"
	
	err = db.Raw(query).Scan(&Cummulative_Fil_Rewards_f01918123_node_2).Error
	
	if err!=nil {
		fmt.Printf("Fil_rewards cannot be fetched, check your Error, %s\n", err)
	}
	prevCummulative_fil_rewards_for_node_2:=Cummulative_Fil_Rewards_f01918123_node_2
	fmt.Printf("The pervious cumulative_fil_rewards_for_node_2 is %d\n",prevCummulative_fil_rewards_for_node_2)
	
    
	f, err:=strconv.ParseFloat(Miner_Info_f01918123.Miner.TotalRewards, 64)
	if err !=nil {
		fmt.Printf("Total Rewards Cannot be converted into integer, Check your error\n")
		return
	}
	latestCummulative_fil_rewards_for_node_2:=int64(f*1e-18)
	fmt.Printf("latestCummulative_fil_rewards_for_node_2, %d\n", latestCummulative_fil_rewards_for_node_2)

	if latestCummulative_fil_rewards_for_node_2 > prevCummulative_fil_rewards_for_node_2{

		FIL_REWARDS_f01918123_node_2=latestCummulative_fil_rewards_for_node_2-prevCummulative_fil_rewards_for_node_2
	
	}else {
		FIL_REWARDS_f01918123_node_2=0 
		fmt.Printf("you have no fil_rewards in last 24 hours")
	}

	QualityAdjPower_f01918123_node_2,err:= strconv.ParseInt(Miner_Info_f01918123.Miner.QualityAdjPower,10,64)
	if err !=nil  {
		  fmt.Printf("The Quality adjusted Power of node2 cannot be converted into int")
	}
     Node_info:= <-c
	 Node_info.Fil_Rewards_f01918123_node_2=FIL_REWARDS_f01918123_node_2
	 Node_info.FRP_f01918123_node_2_adjP=QualityAdjPower_f01918123_node_2
	 Node_info.Cummulative_Fil_Rewards_f01918123_node_2=latestCummulative_fil_rewards_for_node_2
    
	
	c<-Node_info
	
	fmt.Printf("Miner Id : %s\n", Miner_Info_f01918123.Id)
	fmt.Printf("The total_qualityAdj for the node_f01918123 is %d\n",QualityAdjPower_f01918123_node_2)
	fmt.Printf("The total_blocks mined for the node__f01918123 are %d\n",FIL_REWARDS_f01918123_node_2)
	return 
	}()                                                   

// calculate Adjusted power and blocks reward fror f01987994
	wg.Add(1)
	go func() {
		response, err:=http.Get("https://filfox.info/api/v1/address/f01987994")
		if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch Total_Rewards_For_Each_Node"})
		return 
	}
	defer response.Body.Close()
	var Miner_Info_f01987994   Node_Related_Info
	if err:=json.NewDecoder(response.Body).Decode(&Miner_Info_f01987994);err!=nil{
		return 
	}
	defer wg.Done()
	FIL_REWARDS_f01987994_node_3=int64(Miner_Info_f01987994.Miner.BlocksMined)
	var Cummulative_Fil_Rewards_f01987994_node_3 int64
	
	query:="SELECT Cummulative_Fil_Rewards_f01987994_node_3 FROM node_info_daily_and_f_il_prices ORDER BY id DESC LIMIT 1"
	
	err = db.Raw(query).Scan(&Cummulative_Fil_Rewards_f01987994_node_3).Error
	if err!=nil {
		fmt.Printf("Fil_rewards cannot be fetched, check your Error in Line 196\n")
	}
	prevCummulative_fil_rewards_for_node_3:=Cummulative_Fil_Rewards_f01987994_node_3

	fmt.Printf("The pervious cumulative_fil_rewards_for_node_3 is %d\n",prevCummulative_fil_rewards_for_node_3)
	
    
	f, err:=strconv.ParseFloat(Miner_Info_f01987994.Miner.TotalRewards, 64)
	if err !=nil {
		fmt.Printf("TotalRewards Cannot be converted into integer, Check your error\n")
		return
	}
	latestCummulative_fil_rewards_for_node_3:=int64(f*1e-18)
	fmt.Printf("latestCummulative_fil_rewards_for_node_3, %d\n", latestCummulative_fil_rewards_for_node_3)

	if latestCummulative_fil_rewards_for_node_3 > prevCummulative_fil_rewards_for_node_3{

		FIL_REWARDS_f01987994_node_3=latestCummulative_fil_rewards_for_node_3-prevCummulative_fil_rewards_for_node_3
	
	}else {
		FIL_REWARDS_f01987994_node_3=0 
		fmt.Printf("you have no fil_rewards in last 24 hours")
	}
	
	QualityAdjPower_f01987994_node_3, err :=strconv.ParseInt(Miner_Info_f01987994.Miner.QualityAdjPower, 10, 64)
	if err!=nil {
		fmt.Printf("The Qaulity adjusted Power of node3 cannot be converted into int.")
	}
	Node_info:= <-c
	Node_info.Fil_Rewards_f01987994_node_3=FIL_REWARDS_f01987994_node_3
	Node_info.FRP_f01987994_node_3_adjP=QualityAdjPower_f01987994_node_3
	Node_info.Cummulative_Fil_Rewards_f01987994_node_3=latestCummulative_fil_rewards_for_node_3
    
	c<-Node_info
	fmt.Printf("Miner Id : %s\n", Miner_Info_f01987994.Id)
	fmt.Printf("The total_qualityAdj for the node_f01987994 is %d\n",QualityAdjPower_f01987994_node_3)
	fmt.Printf("The Fil for the node_f01987994 are %d\n",FIL_REWARDS_f01987994_node_3)
	return 
	}()

	wg.Wait()
	Node_info:=<-c
	fmt.Printf("The node infos are %v\n", Node_info )
    
	Node_Info_Daily_and_FIl_Price:=Node_Info_Daily_and_FIl_Price{}

	Node_Info_Daily_and_FIl_Price.Fil_Price= Node_info.Fil_Price
    Node_Info_Daily_and_FIl_Price.Fil_Rewards_f01624021_node_1=Node_info.Fil_Rewards_f01624021_node_1
    Node_Info_Daily_and_FIl_Price.Fil_Rewards_f01918123_node_2=Node_info.Fil_Rewards_f01918123_node_2
	Node_Info_Daily_and_FIl_Price.Fil_Rewards_f01987994_node_3=Node_info.Fil_Rewards_f01987994_node_3
	Node_Info_Daily_and_FIl_Price.FRP_f01624021_node_1_adjP=Node_info.FRP_f01624021_node_1_adjP
	Node_Info_Daily_and_FIl_Price.FRP_f01918123_node_2_adjP=Node_info.FRP_f01918123_node_2_adjP
	Node_Info_Daily_and_FIl_Price.FRP_f01987994_node_3_adjP=Node_info.FRP_f01987994_node_3_adjP
	Node_Info_Daily_and_FIl_Price.Cummulative_Fil_Rewards_f01624021_node_1=Node_info.Cummulative_Fil_Rewards_f01624021_node_1
	Node_Info_Daily_and_FIl_Price.Cummulative_Fil_Rewards_f01918123_node_2=Node_info.Cummulative_Fil_Rewards_f01918123_node_2
	Node_Info_Daily_and_FIl_Price.Cummulative_Fil_Rewards_f01987994_node_3=Node_info.Cummulative_Fil_Rewards_f01987994_node_3
	
	/*for key,value := range Node_Info_Daily_and_FIl_Price{
		fmt.Printf("%s:  %v\n", key, value)
	}*/

	db.Create(&Node_Info_Daily_and_FIl_Price)
	
	return nil

}


type FMP_Info_for_investor struct {
	gorm.Model
	Date string `json:"date"`
	Total_Quality_adjP_For_Vogo_Daily_Basis int64  `json:"total_Quality_adjP_For_Vogo_Daily_Basis"`
    Total_FIL_Reward_Vogo_daily_Basis  int64		`json:"total_FIL_Reward_Vogo_daily_Basis"`
    Total_Quality_adjP_For_Inv_daily_Basis int64  `json:"total_Quality_adjP_For_Inv_daily_Basis"`
	Total_Quality_adjP_with_increased_FRP_inv int64  `json:"total_Quality_adjP_With_increased_FRP_Inv"`
	Fil_Rewards_on_daily_basis_for_inv    int64   `json:"fil_Rewards_on_daily_basis"`
	Total_Fil_rewards_for_Inv int64		`json:"total_Fil_rewards_for_Inv"`
	Total_FIL_Rewards int64 `json:"total_FIL_Rewards"`
	Staking_on_daily_basis int64  `json:"staking_on_daily_basis"`
	Total_Staking  int64 `json:"total_Staking"`
	Total_Reward_value  int64  `json:"total_Reward_value"`
	Increased_FRP_on_daily_basis float32 `json:"increase_FRP"`
	Total_FRP float32  `json:"total_FRP"`
	Paid_Reward_to_Investor int64`json:"paid_Reward_to_Investor"`
	Total_FIL_Paid_to_Investor int64 `json:"total_FIL_Paid_to_Investor"`
	Value_of_FIL_Paid_to_Investor int64  `json:"value_of_FIL_Paid_to_Investor"`
	Value_of_Total_FIl_Paid  int64  `json:"value_of_Total_FIl_Paid"`
}


func FMP_investment_Calculate() *FMP_Info_for_investor {
    
	// calculated on the day of investement  i.e will remain consant from the day of investment till the 25th of the next month and on 25th at 12.00 am it wil be updated. 
	// i.e it is  updated only once a month. 
	//Total_Quality_adjP_on_daily_basis:=0 
	db:=DbConnect()
	Node_info:=QueryNodeinfo(db)
	FMP_INFO:=Query_Fmp_table(db)
	//calculate the time 
	now:=time.Now()
	//initiliaze the instance of struct for FMP_INFO which is the main struct where all the data will be stored. 
	FMP_Info := &FMP_Info_for_investor{}

	// Checks if the date is 25th of the month and time is 0.00 am 
	if now.Day()==25 && now.Hour()==0 && now.Minute()==0{
	   // Since Node_info is updated once everyday at  
		total_Quality_adjP_For_Vogo:=Node_info.FRP_f01624021_node_1_adjP+ Node_info.FRP_f01918123_node_2_adjP + Node_info.FRP_f01987994_node_3_adjP
		FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis=total_Quality_adjP_For_Vogo
		FMP_Info.Total_Quality_adjP_For_Inv_daily_Basis=FMP_INFO.Total_Quality_adjP_For_Inv_daily_Basis
	} else {
		
		FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis=FMP_INFO.Total_Quality_adjP_For_Vogo_Daily_Basis
	}
		
// query from thd node_info_daily_f_il_price

	Total_FIL_Reward_Vogo_daily_Basis := Node_info.Fil_Rewards_f01624021_node_1 +Node_info.Fil_Rewards_f01918123_node_2 + Node_info.Fil_Rewards_f01987994_node_3
	FMP_Info.Total_FIL_Reward_Vogo_daily_Basis= Total_FIL_Reward_Vogo_daily_Basis

	total_Quality_adjP_For_Inv := 500 
	/*500 Tib will be on the day of investment and wll be used to calculate the increased FRMo the date of investement 
	till the 25 th of the each moent
	*/
	
	if Total_FIL_Reward_Vogo_daily_Basis == 0 {
	
		FMP_Info.Fil_Rewards_on_daily_basis_for_inv= 0
	
		} else {
	
			FMP_Info.Fil_Rewards_on_daily_basis_for_inv = (int64(total_Quality_adjP_For_Inv)/FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis ) * Total_FIL_Reward_Vogo_daily_Basis
	}

	Staking_on_daily_basis := FMP_Info.Fil_Rewards_on_daily_basis_for_inv

	FMP_Info.Staking_on_daily_basis = Staking_on_daily_basis

	Sector_initial_pledge := Node_info.Current_Sector_Initial_Pledge_32GB
	

	if Staking_on_daily_basis == 0 {
	
		FMP_Info.Increased_FRP_on_daily_basis = 0
	
		} else {
	
			FMP_Info.Increased_FRP_on_daily_basis = float32(Staking_on_daily_basis)/float32(Sector_initial_pledge * 32)
	}
	
	
	prevTotalFILRewards := FMP_INFO.Total_FIL_Rewards
	prevTotalFRP:= FMP_INFO.Total_FRP
	PrevTotalStaking:=FMP_INFO.Total_Staking
	prevTotal_Quality_AdjPow_inv:= FMP_INFO.Total_Quality_adjP_with_increased_FRP_inv

	FMP_Info.Total_Staking=PrevTotalStaking+FMP_Info.Staking_on_daily_basis
	FMP_Info.Total_FIL_Rewards=prevTotalFILRewards+FMP_Info.Fil_Rewards_on_daily_basis_for_inv
	FMP_Info.Total_FRP=prevTotalFRP+FMP_Info.Increased_FRP_on_daily_basis
	FMP_Info.Total_Quality_adjP_with_increased_FRP_inv= int64(FMP_Info.Increased_FRP_on_daily_basis+float32(prevTotal_Quality_AdjPow_inv))	
	db.Create(&FMP_Info)
	return FMP_Info
}


func QueryNodeinfo(db *gorm.DB) Node_Info_Daily_and_FIl_Price_ {
	query:="SELECT * from node_info_daily_and_f_il_prices ORDER BY DATE ASC LIMIT 1"
	var node_info Node_Info_Daily_and_FIl_Price_
	err:=db.Raw(query).Scan(&node_info).Error
	if err!=nil {
		fmt.Printf("The node information cannot be fetched %s\n",err)
	}
	fmt.Println(node_info)
	return node_info // printing function for flot type. 
	}


func Query_Fmp_table(db *gorm.DB) FMP_Info_for_investor_{
	query:= " SELECT * from  fmp_info_for_investors ORDER BY DATE ASC LIMIT 1"
	var FMP_Info FMP_Info_for_investor_
	err:=db.Raw(query).Scan(&FMP_Info).Error
	if err !=nil {
		fmt.Printf("The FMP_information cannot be fetched %s\n", err)
	}
	fmt.Println(FMP_Info)
	return FMP_Info
}



type FMP_Info_for_investor_ struct {
	Date string 
   Total_Quality_adjP_For_Vogo_Daily_Basis int64  
   Total_FIL_Reward_Vogo_daily_Basis  int64      
   Total_Quality_adjP_For_Inv_daily_Basis int64  
   Total_Quality_adjP_with_increased_FRP_inv int64
   Fil_Rewards_on_daily_basis_for_inv    int64  
   Total_Fil_rewards_for_Inv int64    
   Total_FIL_Rewards int64  
   Staking_on_daily_basis int64  
   Total_Staking  int64 
   Total_Reward_value  int64  
   Increased_FRP_on_daily_basis float32 
   Total_FRP float32  
   Paid_Reward_to_Investor int64 
   Total_FIL_Paid_to_Investor int64 
   Value_of_FIL_Paid_to_Investor int64
   Value_of_Total_FIl_Paid  int64  }
// struct to get data base from the database where data were stored from the api call. 
type Node_Info_Daily_and_FIl_Price_ struct{ 
		Date  string 
		Fil_Price float32 
		Current_Sector_Initial_Pledge_32GB float32 
		Fil_Rewards_f01624021_node_1 int64
		Fil_Rewards_f01918123_node_2 int64
		Fil_Rewards_f01987994_node_3 int64
		Cummulative_Fil_Rewards_f01624021_node_1 int64
		Cummulative_Fil_Rewards_f01918123_node_2 int64
		Cummulative_Fil_Rewards_f01987994_node_3 int64
		FRP_f01624021_node_1_adjP int64
		FRP_f01918123_node_2_adjP int64
		FRP_f01987994_node_3_adjP int64
}


// function  to connect to database 
func DbConnect() *gorm.DB {
	err:=godotenv.Load(".env")
   if err!=nil{
	fmt.Printf("")
   }
    dbUser:=os.Getenv("USERNAME")
	dbPassword:=os.Getenv("PASSWORD")
	dbIP:=os.Getenv("DBIP")
    dbPort :=os.Getenv("DBPORT")
	dbName:=os.Getenv("DBNAME")
	dbSslMode:=os.Getenv("DBSSLMODE")
	
	dbConfig:= config.Config {
       User:dbUser,
	   Password:dbPassword,
	   Host:dbIP,
	   Port:dbPort,
	   DbName:dbName,
	   SslMode:dbSslMode,
	}
    db, err:=config.Connect(&dbConfig)
   
	if err != nil {
		fmt.Printf("There is error in connecting to data\n")
	}
	return db
}