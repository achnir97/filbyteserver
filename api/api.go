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
/*type FMP_Investment_Info_From_API_on_Daily struct  {
	Date Date `json:"date" validate:"required"`
	Fil_Price float32 `json:"fil_price" validate:"required,gte=0"`
	Current_Sector_Initial_Pledge_32GB float32 `json:"current_sector_initial_pledge_32gb" validate:"required,gte=0"`
	Fil_Rewards_f01624021_node_1 float32 `json:"fil_rewards_f01624021_node_1" validate:"required,gte=0"`
	Fil_Rewards_f01918123_node_2 float32 `json:"fil_rewards_f01918123_node_2" validate:"required,gte=0"`	
	Fil_Rewards_f01987994_node_3 float32 `json:"fil_rewards_f01987994_node_3" validate:"required,gte=0"`
	FRP_f01624021_node_1 float32 `json:"frp_f01624021_node_1" validate:"required,gte=0"`
	FRP_f01918123_node_2 float32 `json:"frp_f01918123_node_2" validate:"required,gte=0"`
	FRP_f01987994_node_3 float32 `json:"frp_f01987994_node_3" validate:"required,gte=0"`
}
*/
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



//Get FIL_Rewards and Quality adjusted power of node f01624021 on daily basis 
func FIL_Price_n_Block_rewards_for_Each_Node(context *fiber.Ctx)error{ 
    var wg sync.WaitGroup
     
    go func() {
	wg.Add(1)
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
	fmt.Printf("The Price of Filecoin is %f\n", FilPrice.Filecoin.Krw)
	}()

// calculate Adjusted power and blocks reward fror f01624021
	go func() {
		wg.Add(1)
		response, err:=http.Get("https://filfox.info/api/v1/address/f01624021")
		if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch Total_Rewards_For_Each_Node"})
		return 
	}
	defer response.Body.Close()
	var Miner_Info   Node_Related_Info
	if err:=json.NewDecoder(response.Body).Decode(&Miner_Info);err!=nil{
		return 
	}
	defer wg.Done()
	fmt.Printf("Miner Id : %s\n", Miner_Info.Id)
	fmt.Printf("The total_qualityAdj for the node_f01624021 is %s\n",Miner_Info.Miner.QualityAdjPower)
	fmt.Printf("The total_blocks mined for the node_f01624021 are %d\n",Miner_Info.Miner.BlocksMined)
	
	return 
	}()

// calculate Adjusted power and blocks reward fror f01819003
	go func() {
		wg.Add(1)
		response, err:=http.Get("https://filfox.info/api/v1/address/f01819003")
		if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch Total_Rewards_For_Each_Node"})
		return 
	}
	defer response.Body.Close()
	var Miner_Info   Node_Related_Info
	if err:=json.NewDecoder(response.Body).Decode(&Miner_Info);err!=nil{
		return 
	}
	defer wg.Done()
	fmt.Printf("Miner Id : %s\n", Miner_Info.Id)
	fmt.Printf("The total_qualityAdj for the node_f01819003 is %s\n",Miner_Info.Miner.QualityAdjPower)
	fmt.Printf("The total_blocks mined for the node_f01819003 are %d\n",Miner_Info.Miner.BlocksMined)
	return 
	}()

// calculate Adjusted power and blocks reward fror f01987994
	go func() {
		wg.Add(1)
		response, err:=http.Get("https://filfox.info/api/v1/address/f01987994")
		if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch Total_Rewards_For_Each_Node"})
		return 
	}
	defer response.Body.Close()
	var Miner_Info   Node_Related_Info
	if err:=json.NewDecoder(response.Body).Decode(&Miner_Info);err!=nil{
		return 
	}
	defer wg.Done()
	fmt.Printf("Miner Id : %s\n", Miner_Info.Id)
	fmt.Printf("The total_qualityAdj for the node_f01987994 is %s\n",Miner_Info.Miner.QualityAdjPower)
	fmt.Printf("The total_blocks mined for the node_f01987994 are %d\n",Miner_Info.Miner.BlocksMined)
	
	return 
	}()

	wg.Wait()
	return nil
}
 

