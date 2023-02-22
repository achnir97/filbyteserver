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
	"strconv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/robfig/cron"
)

type MinerDetails struct {
   QualityAdjPower  string `json:"qualityAdjPower"`
   NetworkRawBytePower string `json:"networkRawBytePower"`
   BlocksMined string `json:"blocksMined"`
   WeightedBlocksMined string `json:"weightedBlocksMined"`
   TotalRewards  string `json:"totalRewards"`
}
 

var  Total_Quality_adjP_on_daily_basis_for_Vogo int
var  Total_Quality_adjP_on_daily_basis_for_Inv  int

type Fetched_Info struct{
	Id string `json:"id"`
	Miner MinerDetails `json:"miner"`
 }


type Fil_price struct{
	Filecoin Price `json:"filecoin"`
	}

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
		SectorSize int64 `json:"sectorSize"`
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



type FMP_Investment_Integrated_info struct{ 
    gorm.Model 
	Fil_Price float32 
    Current_Sector_Initial_Pledge_32GB float32 
	Fil_Rewards_f01624021_node_1 int 
	Fil_Rewards_f01918123_node_2 int
	Fil_Rewards_f01987994_node_3 int
	FRP_f01624021_node_1_adjP int
	FRP_f01918123_node_2_adjP int
	FRP_f01987994_node_3_adjP int
}


//Method to assign values to FMP_integrated_info
func(Miner *FMP_Investment_Integrated_info)Set_Miner_Info(Fil_Price float32, Fil_Rewards_f01624021_node_1 int, Fil_Rewards_f01918123_node_2 int,
	Fil_Rewards_f01987994_node_3 int, FRP_f01624021_node_1 int, FRP_f01918123_node_2 int, 
	FRP_f01987994_node_3 int){
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

	 c:=make(chan FMP_Investment_Integrated_info,1)

	//var miner FMP_Investment_Integrated_info 
	var FIL_PRICE float32
	var FIL_REWARDS_f01624021_node_1  int 
	var FIL_REWARDS_f01918123_node_2  int
	var FIL_REWARDS_f01987994_node_3 int
	/*var QualityAdjPower_f01624021_node_1 int
	var QualityAdjPower_f01918123_node_2 int
	var QualityAdjPower_f01987994_node_3 int */

    // get the price of FILCOIN on daily basis.  
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
	c<-FMP_Investment_Integrated_info{Fil_Price:FIL_PRICE}
	

	fmt.Printf("The Price of Filecoin is %f\n", FIL_PRICE)
	}()

// calculate Adjusted power and blocks reward fror f01624021
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
	FIL_REWARDS_f01624021_node_1=Miner_Info_f01624021.Miner.BlocksMined
	QualityAdjPower_f01624021_node_1, err :=strconv.Atoi(Miner_Info_f01624021.Miner.QualityAdjPower) 
	if err !=nil {
		fmt.Printf("Quality adjusted power_cannot be converted into int")
	}
	QualityAdjPower_f01624021_node_1=QualityAdjPower_f01624021_node_1
	Node_info:=<-c
	Node_info.Fil_Rewards_f01624021_node_1=FIL_REWARDS_f01624021_node_1
	Node_info.FRP_f01624021_node_1_adjP =QualityAdjPower_f01624021_node_1
    c<-Node_info

	fmt.Printf("Miner Id : %s\n", Miner_Info_f01624021.Id)
	fmt.Printf("The total_qualityAdj for the node_f01624021 is %s\n",QualityAdjPower_f01624021_node_1)
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
	FIL_REWARDS_f01918123_node_2=Miner_Info_f01918123.Miner.BlocksMined
	QualityAdjPower_f01918123_node_2,err:= strconv.Atoi(Miner_Info_f01918123.Miner.QualityAdjPower)
	if err !=nil  {
		  fmt.Printf("The Quality adjusted Power of node2 cannot be converted into int")
	}
     Node_info:= <-c
	 Node_info.Fil_Rewards_f01918123_node_2=FIL_REWARDS_f01918123_node_2
	 Node_info.FRP_f01918123_node_2_adjP=QualityAdjPower_f01918123_node_2
	
	c<-Node_info
	
	fmt.Printf("Miner Id : %s\n", Miner_Info_f01918123.Id)
	fmt.Printf("The total_qualityAdj for the node_f01819003 is %s\n",QualityAdjPower_f01918123_node_2)
	fmt.Printf("The total_blocks mined for the node_f01819003 are %d\n",FIL_REWARDS_f01918123_node_2)
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
	FIL_REWARDS_f01987994_node_3=Miner_Info_f01987994.Miner.BlocksMined
	QualityAdjPower_f01987994_node_3, err :=strconv.Atoi(Miner_Info_f01987994.Miner.QualityAdjPower)
	if err!=nil {
		fmt.Printf("The Qaulity adjusted Power of node3 cannot be converted into int.")
	}
	Node_info:= <-c
	Node_info.Fil_Rewards_f01987994_node_3=FIL_REWARDS_f01987994_node_3
	Node_info.FRP_f01987994_node_3_adjP=QualityAdjPower_f01987994_node_3
   c<-Node_info
	fmt.Printf("Miner Id : %s\n", Miner_Info_f01987994.Id)
	fmt.Printf("The total_qualityAdj for the node_f01987994 is %s\n",QualityAdjPower_f01987994_node_3)
	fmt.Printf("The total_blocks mined for the node_f01987994 are %d\n",FIL_REWARDS_f01987994_node_3)

	return 
	}()

		
	wg.Wait()
	Node_info:=<-c
	fmt.Printf("The node infos are %v\n", Node_info )	
	return nil
}


type FMP_Info_for_investor struct {
	gorm.Model 
	ID uint   `gorm:"primaryKey"`
	Total_Quality_adjP_For_Vogo_Daily_Basis int  `json:"total_Quality_adjP_For_Vogo_Daily_Basis"`
    Total_FIL_Reward_Vogo_daily_Basis  int 		`json:"total_FIL_Reward_Vogo_daily_Basis"`
    Total_Quality_adjP_For_Inv_daily_Basis int  `json:"total_Quality_adjP_For_Inv_daily_Basis"`
	Total_Quality_adjP_with_increased_FRP_inv int  `json:"total_Quality_adjP_With_increased_FRP_Inv"`
	Fil_Rewards_on_daily_basis     int   `json:"fil_Rewards_on_daily_basis"`
	Total_Fil_rewards_for_Inv int		`json:"total_Fil_rewards_for_Inv`
	Total_FIL_Rewards int `json:"total_FIL_Rewards`
	Staking_on_daily_basis int  `json:"staking_on_daily_basis"`
	Total_Staking  int `json:"total_Staking"`
	Total_Reward_value  int  `json:'total_Reward_value`
	Increased_FRP_on_daily_basis float32 `json:"increase_FRP"`
	Total_FRP float32  `json:"total_FRP"`
	Paid_Reward_to_Investor int `json:"paid_Reward_to_Investor"`
	Total_FIL_Paid_to_Investor int `json:"total_FIL_Paid_to_Investor"`
	Value_of_FIL_Paid_to_Investor int  `json:"value_of_FIL_Paid_to_Investor"`
	Value_of_Total_FIl_Paid  int  `json:"value_of_Total_FIl_Paid"`
}


func FMP_investment_Calculate(Node_info *FMP_Investment_Integrated_info) *FMP_Info_for_investor {
     
	Total_Quality_adjP_on_daily_basis:=0 
	now:=time.Now()
	FMP_Info := &FMP_Info_for_investor{}
	
	if now.Day()==25 && now.Hour()==0 && now.Minute()==0{
	    
		var Total_Quality_adjP_For_Inv_daily_Basis int 
		
		total_Quality_adjP_For_Vogo := Node_info.FRP_f01624021_node_1_adjP + Node_info.FRP_f01918123_node_2_adjP + Node_info.FRP_f01987994_node_3_adjP
		
		Total_Quality_adjP_on_daily_basis_for_Vogo=total_Quality_adjP_For_Vogo
		
		FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis=total_Quality_adjP_For_Vogo
		
		query := "SELECT Total_Quality_adjP_with_increased_FRP_inv  FROM table_name ORDER BY id DESC LIMIT 1"
		
		err:=db.Raw(query).Scan(&Total_Quality_adjP_For_Inv_daily_Basis).Error
		
		if err!=nil {
			fmt.Printf("You cannot query Total_Quality_AdjP_with_increassed_FRP_Inv")}

		FMP_Info.Total_Quality_adjP_For_Inv_daily_Basis=Total_Quality_adjP_For_Inv_daily_Basis
	
	} else {
		var Total_Quality_adjP_For_Vogo_Daily_Basis int
		
		query := "SELECT Total_Quality_adjP_For_Vogo_Daily_Basis FROM table_name ORDER BY id DESC LIMIT 1"
		
		err:=db.Raw(query).Scan(&Total_Quality_adjP_For_Vogo_Daily_Basis).Error
		
		if err!=nil {
		
			fmt.Printf("Your database cannot be queries")
		}
		FMP_Info.Total_Quality_adjP_For_Vogo_Daily_Basis=Total_Quality_adjP_For_Vogo_Daily_Basis
		
	}
		

	
	Total_FIL_Reward_Vogo_daily_Basis := Node_info.Fil_Rewards_f01624021_node_1 + Node_info.Fil_Rewards_f01918123_node_2 + Node_info.Fil_Rewards_f01987994_node_3
	
	FMP_Info.Total_Fil_rewards_For_Vogo = Total_FIL_Reward_Vogo

	total_Quality_adjP_For_Inv := 500.0

	if Total_FIL_Reward_Vogo == 0 {
	
		FMP_Info.Fil_Rewards_on_daily_basis = 0
	
		} else {
	
			FMP_Info.Fil_Rewards_on_daily_basis = (total_Quality_adjP_For_Inv /total_Quality_adjP_For_Vogo) * Total_FIL_Reward_Vogo
	}

	Staking_on_daily_basis := FMP_Info.Fil_Rewards_on_daily_basis

	FMP_Info.Staking_on_daily_basis = Staking_on_daily_basis

	Sector_initial_pledge := Node_info.Initial_Collateral_Sector_pledge

	if Staking_on_daily_basis == 0 {
	
		FMP_Info.Increased_FRP = 0
	
		} else {
	
			FMP_Info.Increased_FRP_on_daily_basis = Staking_on_daily_basis / (Sector_initial_pledge * 32)
	}
	var (
	
		prevTotalFILRewards int64
	
		prevTotalFRP int64
	
		prevTotalStaking int64
	
		prevTotal_Quality_AdjPow int64 
	)
	
	query := "SELECT Total_FIL_Reward, Total_FRP, Total_Staking,Total_Quality_adjP_with_increased_FRP_inv  FROM table_name ORDER BY id DESC LIMIT 1"
	
	err := db.Raw(query).Scan(&prevTotalFILRewards, &prevTotalFRP, &prevTotalStaking, &prevTotal_Quality_AdjPow).Error
	
	if err != nil {	
		fmt.Printf("The database cannot be fetched from the database ")
	}

    
	FMP_Info.Total_Staking=prevTotalFRP+FMP_Info.Staking_on_daily_basis
	
	FMP_Info.Total_FIL_Reward=prevTotalFILRewards+FMP_Info.Fil_Rewards_on_daily_basis
	
	FMP_Info.Total_FRP=prevTotalFRP+FMP.Info.Increased_FRP_on_daily_basis
    
	FMP_Info.Total_Quality_adjP_with_increased_FRP_inv=prevTotal_Quality_AdjPow+FMP.Info.Increased_FRP_on_daily_basis
	
	db.Create(&FMP_Info)
}


