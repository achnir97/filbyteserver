package api 
 import (
	"fmt"
	_"time"
	"net/http"
	_"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_"io/ioutil"
	_"gorm.io/gorm"
	"encoding/json"
	"sync"
)

type MinerDetails struct {
   QualityAdjPower  string `json:"qualityAdjPower"`
   NetworkRawBytePower string `json:"networkRawBytePower"`
   BlocksMined string `json:"blocksMined"`
   WeightedBlocksMined string `json:"weightedBlocksMined"`
   TotalRewards  string `json:"totalRewards"`
}
 

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
	Fil_Price float32 
	Current_Sector_Initial_Pledge_32GB float32 
	Fil_Rewards_f01624021_node_1 int 
	Fil_Rewards_f01918123_node_2 int
	Fil_Rewards_f01987994_node_3 int
	FRP_f01624021_node_1_adjP string 
	FRP_f01918123_node_2_adjP string
	FRP_f01987994_node_3_adjP string
}

//Method to assign values to FMP_integrated_info
func(Miner *FMP_Investment_Integrated_info)Set_Miner_Info(Fil_Price float32, Fil_Rewards_f01624021_node_1 int, Fil_Rewards_f01918123_node_2 int,
	Fil_Rewards_f01987994_node_3 int, FRP_f01624021_node_1 string, FRP_f01918123_node_2 string, 
	FRP_f01987994_node_3 string){
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

	Node_Info:=make(chan FMP_Investment_Integrated_info)

	//var miner FMP_Investment_Integrated_info 
	var FIL_PRICE float32
	var FIL_REWARDS_f01624021_node_1  int 
	var FIL_REWARDS_f01819003_node_2  int
	var FIL_REWARDS_f01987994_node_3 int
	var QualityAdjPower_f01624021_node_1 string
	var QualityAdjPower_f01819003_node_2 string
	var QualityAdjPower_f01987994_node_3 string

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
	Fil_price:=<-Node_Info
	Fil_price.Fil_Price=FIL_PRICE
	Node_Info<-Fil_price

	fmt.Printf("The Price of Filecoin is %f\n", FilPrice.Filecoin.Krw)
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
	QualityAdjPower_f01624021_node_1=Miner_Info_f01624021.Miner.QualityAdjPower
	node_1_info:=<-Node_Info
	node_1_info.Fil_Rewards_f01624021_node_1=FIL_REWARDS_f01624021_node_1
	node_1_info.FRP_f01624021_node_1_adjP=QualityAdjPower_f01624021_node_1
	Node_Info<-node_1_info

	fmt.Printf("Miner Id : %s\n", Miner_Info_f01624021.Id)
	fmt.Printf("The total_qualityAdj for the node_f01624021 is %s\n",QualityAdjPower_f01624021_node_1)
	fmt.Printf("The total_blocks mined for the node_f01624021 are %d\n",FIL_REWARDS_f01624021_node_1)
	return 
	}()

// calculate Adjusted power and blocks reward fror f01819003
	wg.Add(1)
	go func() {
		response, err:=http.Get("https://filfox.info/api/v1/address/f01819003")
		if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch Total_Rewards_For_Each_Node"})
		return 
	}
	defer response.Body.Close()
	var Miner_Info_f01819003  Node_Related_Info
	if err:=json.NewDecoder(response.Body).Decode(&Miner_Info_f01819003);err!=nil{
		return 
	}
	defer wg.Done()
	FIL_REWARDS_f01819003_node_2=Miner_Info_f01819003.Miner.BlocksMined
	QualityAdjPower_f01819003_node_2=Miner_Info_f01819003.Miner.QualityAdjPower

	node_1_info:=<-Node_Info
	node_1_info.Fil_Rewards_f01918123_node_2 =FIL_REWARDS_f01819003_node_2
	node_1_info.FRP_f01918123_node_2_adjP=QualityAdjPower_f01819003_node_2
	Node_Info<-node_1_info
	
	fmt.Printf("Miner Id : %s\n", Miner_Info_f01819003.Id)
	fmt.Printf("The total_qualityAdj for the node_f01819003 is %s\n",QualityAdjPower_f01819003_node_2)
	fmt.Printf("The total_blocks mined for the node_f01819003 are %d\n",FIL_REWARDS_f01819003_node_2)
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
	QualityAdjPower_f01987994_node_3=Miner_Info_f01987994.Miner.QualityAdjPower
	node_1_info:=<-Node_Info
	node_1_info.Fil_Rewards_f01987994_node_3=FIL_REWARDS_f01987994_node_3
	node_1_info.FRP_f01987994_node_3_adjP=QualityAdjPower_f01987994_node_3
	Node_Info<-node_1_info
	fmt.Printf("Miner Id : %s\n", Miner_Info_f01987994.Id)
	fmt.Printf("The total_qualityAdj for the node_f01987994 is %s\n",QualityAdjPower_f01987994_node_3)
	fmt.Printf("The total_blocks mined for the node_f01987994 are %d\n",FIL_REWARDS_f01987994_node_3)

	return 
	}()

		
	wg.Wait()
	/*miner.Set_Miner_Info(FIL_PRICE,FIL_REWARDS_f01624021_node_1,FIL_REWARDS_f01819003_node_2,FIL_REWARDS_f01918123_node_3,
		QualityAdjPower_f01624021_node_1,QualityAdjPower_f01819003_node_2,QualityAdjPower_f01918123_node_3)*/
	
	fmt.Printf("The extracted information are %v\n",<-Node_Info )
	
	return nil
}
 