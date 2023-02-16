package api 
 import (
	"fmt"
	_"time"
	"net/http"
	_"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"gorm.io/gorm"
	_"encoding/json"
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
	Miner *MinerDetails
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

type Repository struct{
	db *gorm.DB 
}

type Response_from_node_1 struct {
	
}
func GetFIL_Price_on_24Hour_basis(context *fiber.Ctx) error { 
	response, err:=http.Get("https://api.coingecko.com/api/v3/simple/price?ids=Filecoin&vs_currencies=KRW")
	if err!=nil {
		fmt.Printf("The Htpp request failed with errp %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch FIl_Price"})
			return err
	}
	defer response.Body.Close()
	data,err:=ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body with error :%s\n", err)
		context.SendStatus(http.StatusInternalServerError)
		return err
}
 		context.JSON(data)
		fmt.Printf("The price of filecoin so obtained is %s\n", string(data))
		return nil 
}

//Get FIL_Rewards and Quality adjusted power of node f01624021 on daily basis 
func GetRewards_For_Each_Node_f01624021(context *fiber.Ctx)error{ 

	response, err:=http.Get("https://filfox.info/api/v1/address/f01624021")
	if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch Total_Rewards_For_Each_Node"})
		return err
	}
	defer response.Body.Close()
	//var Fetched_data Fetched_Info
	//data:=context.JSON(response.Body)
	data, err:=ioutil.ReadAll(response.Body)
	if err !=nil {
		fmt.Printf("failed to read Response body with error")
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Could not read information from the Response\n"},
		)
		return err
	}
	 //var Fetched_Info *Fetched_Info
	//totalRewards:=Fetched_data.Miner.totalRewards
	 context.JSON(data)
	// if err:= json.Unmarshal(data, &Fetched_Info); err!=nil {
	// 	fmt.Printf("Error Occured, Try to solve that error\n")
	// 	return err
	// }
	fmt.Println(string(data))
	return nil
}

//Get FIL_Rewards and Quality adjusted power of node f01918123 on daily basis 
func( r* Repository) GetRewards_For_Each_Node_f01918123(context *fiber.Ctx) {

	response, err:=http.Get("https://filfox.info/api/v1/address/f01819003")
	if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch Total_Rewards_For_Node_f01918123"})
	}
	defer response.Body.Close()
	data, err:=ioutil.ReadAll(response.Body)
	if err !=nil {
		fmt.Printf("failed to read Response body with error")
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Could not read information from the Response\n"},
		)
	}
	context.JSON(data)
	fmt.Printf("The Reward for the node_ff01819003 are%d\n",data)
}


//Get FIL_Rewards and Quality adjusted power of node f01987994 on daily basis 
func(r *Repository) GetRewards_For_Each_Node_f01987994(context *fiber.Ctx) {

	response, err:=http.Get("https://filfox.info/api/v1/address/f01987994")
	if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch Total_Rewards_For_Each_Node"})	
	}
	defer response.Body.Close()
	data, err:=ioutil.ReadAll(response.Body)
	if err !=nil {
		fmt.Printf("failed to read Response body with error")
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Could not read information from the Response\n"},
		)
	}
	context.JSON(data)
	fmt.Printf("The Reward for the node_ff01819003 are %d\n",data)
}

// Task  Get the data from the API 
// Calculate the value from the API and s
// Store the value so calculated in sql database. 
// Fetch value from sql database to mysql and render to to the FrontEnd. 
// Render all the required value in the FrontEnd. 

