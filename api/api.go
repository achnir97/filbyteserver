package api 
 import (
	"fmt"
	_"time"
	"net/http"
	_"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
 )

 
 type Fetched_Info struct{

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

func getFIL_Price_on_24Hour_basis(context *fiber.Ctx) error { 

	response, err:=http.Get("https://api.coingecko.com/api/v3/simple/price?ids=fil%20&vs_currencies=KRW%2C%20USD")
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
		fmt.Printf("The price of filecoin so obtained is %d\n", data)
		return nil 
}

func GetRewards_For_Each_Node(context *fiber.Ctx) error{

	response, err:=http.Get("https://filfox.info/api/v1/address/f01624021")
	if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Couldnot fetch Total_Rewards_For_Each_Node"})
		return err
	}
	defer response.Body.Close()
	data, err:=ioutil.ReadAll(response.Body)
	if err !=nil {
		fmt.Printf("failed to read Response body with error")
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message":"Could not read information from the Response\n"},
		)
		return err
	
	}
	
	context.JSON(data)
	fmt.Printf("The Reward for the node_f01624021 are%d\n",data)
	return nil 
}
	/*if data.TotalRewards > Saved_Total_Rewards{
		A=:Saved_totol_Rewards
		Saved_Total_Rewards= data.Total_Rewards
		block_reward_in_last_24_hrs=Saved_Total_Rewards-A
		db.update(block_rewardds_last_24_hr).Where("id=?" ,id )
	}*/

	
