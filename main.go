package main

import(
	_"fmt"
    _"time" 
	_"context"
	"github.com/gofiber/fiber/v2"
	_"net/http"
	"github.com/achnir97/go_lang_filbytes/api"
	_"os"
	"github.com/gofiber/fiber/v2/middleware/cors"

)

/*
type FMP_Investment_Info_From_API_on_Daily struct  {
	Date Date()
	Days  int 
	Fil_Price float32 
	Current_Sector_Initial_Pledge_32GB float32 
	Fil_Rewards_f01624021_node_1 float32
	Fil_Rewards_f01918123_node_2 float32	
	Fil_Rewards_f01987994_node_3 float32
	FRP_f01624021_node_1 float32 // updataed once every 25th of the month  i:e fixed for one month 
	FRP_f01918123_node_2 float32 // updataed once every 25th of the month  i:e fixed for one month
	FRP_f01987994_node_3 float32 // updataed once every 25th of the month  i:e fixed for one month
}
type FRP_upated_Once_every_month struct{
	
	FRP_TiB_Vogo float32 
}

type FMP__Info_VOGO_Calculated_Based_ON_API struct{
	
	Total_FIL_Reward_Vogo float32
	FRP_Investment  float32
	
}


type FMP_Info_Investor_Based_ON_API struct {
	
	Total_FIL_Rewards_Investor float32
	Total_Net_FRP_Value_Investor float32
	Total_Staking_FIL  float32 
	Total_Reward_Value float32
	Total_FRP_Investor float32

}
type Daily_Dyanamic_Value_Investor  struct{
	
	Daily_FIL_Reward_Investor float32 
	Daily_Staking float32

	Daily_Increased_FRP_Investor  float32
	
	Paid_Reward_Investor float32
	
	Total_FIL_Paid_Investor float32
	Value_of_FIL_Paid_Investor float32
	Value_of_Total_FIL_paid float32
}

func Create_FMP_Investment_Static(context *fiber.Context) error {
	FMP_Investment_Info_From_API := &FMP_Investment_Info_From_AP{}
	err:=context.BodyParser(FMP_Investment_Info_From_AP)
	if err!=nil {
		context.Status(http.StatusBadRequest(&fiber.Map{
		"message":"Could not parse the information"
		}))
		return error 
	}

}

func(FMP_Invest_Dynamic *FMP_Investment_Info_Calculated_Based_ON_API{}) Update(FMP_Invest_Static *FMP_Investment_Info_From_API{}){}


func() Update_Node_Adjusted_Power_OnlyOnce_Month(db *gorm.DB, context *fiber.Context) (error)  {

		    FMP_Investment_Info_From_API_on_Daily:= &FMP_Investment_Info_From_API_on_Daily{}
			err:=context.BodyParser(FMP_Investment_Info_From_API_on_Daily)
			if err!=nil {
				fmt.Printf("The Data cannot be parsed")
				return err
			}
			err:=db.Create(FMP_Investment_Info_From_API_on_Daily).Error()
			if err !=nil {
				fmt.Printf("The database cannot be created, check your error correctly")
				return err
			}
			Context.Status(http.StatusOK(&fiber.Map{
				"Message":"Data created successfully",
				"Data":FMP_Investment_Info_From_API_on_Daily{}
			}))
           return nil 
}	
*/


func main() {
	app:=fiber.New()
	app.Use(cors.New(cors.Config{
	AllowOrigins:"*",
	AllowHeaders:"Origin, Content-Type, Accept",
	AllowMethods:"GET. POST, PUT, DELETE",
}))
	app.Get("/filPrice", api.GetFIL_Price_on_24Hour_basis)
	app.Get("/filReward_node1", api.GetRewards_For_Each_Node_f01624021)
	
	// defer app.ReleaseCtx(Ctx)

	// api.GetRewards_For_Each_Node(context *fiber.Ctx)
	// if err!=nil {
	// 	fmt.Printf("The error occured, try to resolve error ")
	// }
	// fmt.Print("Successfuly Fetched Rewards for each Node.")
	

	// err := api.GetFIL_Price_on_24Hour_basis(&context)
	// if err!=nil {

	// }

	// err:= (".env")
	// if err!=nil {
	// 	fmt.Printf("Enviromente Varible from the env file cannot be loaded, Check for the error")
	// 	return err 
	// }
	
	// 	c:=cron.New()
	// 	c.Add("0 0 24 * * ", Update_Node_Adjusted_Power_OnlyOnce_Month()) // implement once every month 
	//     c.Add("0, 0, 0, * * * ",api.GetRewards()) // implement once every 24hrs 
	// 	if err!=nil {
	// 		fmt.Println("Error scheduling task:", err)
	// 	}
	// 	c.Start()
	// 	select{}
	app.Listen(":4000")
	
	}
