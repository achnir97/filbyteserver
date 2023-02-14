package api 
 import (
	"fmt"
	"time"
	"github.com"
	"net/http"
 )

 
 type Fetched_Info struct{

 }
 
 type FMP_Investment_Info_From_API_on_Daily struct  {
	Fil_Price float32 
	Current_Sector_Initial_Pledge_32GB float32 
	Fil_Rewards_f01624021_node_1 float32
	Fil_Rewards_f01918123_node_2 float32	
	Fil_Rewards_f01987994_node_3 float32
	FRP_f01624021_node_1 float32 // updataed once every 25th of the month  i:e fixed for one month 
	FRP_f01918123_node_2 float32 // updataed once every 25th of the month  i:e fixed for one month
	FRP_f01987994_node_3 float32 // updataed once every 25th of the month  i:e fixed for one month
}

func getAPI(context *fiber.Ctx) { 

	response, err:=http.Get("https://api.coingecko.com/api/v3/simple/price?ids=fil%20&vs_currencies=KRW%2C%20USD")
	if err!=nil {
		fmt.Printf("The Htpp request failed with errp %s\n", err)
		c.SendStatus(http.StatusBadRequest)
		return 
}
	data,_=ioutil.ReadAll(response.Body)
	c.JSON(data)

}

func  GetRewards(contex *fiber.Ctx){
	response, err=htpp.Get("https://filfox.info/api/v1/address/f01624021")
	if err !=nil {
		fmt. Printf("The http Request failed with error %s\n", err)
		c.SendStatus(http.StatusBadRequest)
		return
	}
	data, _=ioutil.ReadAll(repsonce.Body)
	c.JSON(data)
	if data.TotalRewards > Saved_Total_Rewards{
		A=:Saved_totol_Rewards
		Saved_Total_Rewards= data.Total_Rewards
		block_reward_in_last_24_hrs=Saved_Total_Rewards-A
		db.update(block_rewardds_last_24_hr).where("id=?" ,id )
	}
	return 
}
