package main

/*
Sync daily data of all stocks in last 100 days.
*/

import (
	"context"
	"fmt"
	"similarStock/model"
	"time"

	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Api = "http://api.tushare.pro"
const Token = "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e"

var allStockDailyParam = `{
	"api_name": "daily", 
	"token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", 
	"params": {"trade_date":"$date"}, 
	"fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"
	}`

var startDate = "2023-01-01"
var endDate = ""

var c = resty.New()

var col *mongo.Collection

func main() {

	client, e := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/test"))
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	col = client.Database("test").Collection("stockDaily")

	for {
		stockDaily := req(allStockDailyParam)
		stockDailyList := handleStockDailyData_2(stockDaily)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		r, e := col.InsertMany(ctx, stockDailyList)
		if e != nil {
			fmt.Println("error:", e.Error())
		}
		fmt.Println("Inserted doc count:", len(r.InsertedIDs))
		defer cancel()
	}
}

func req(paramStr string) model.ResData {
	res := model.Res{}
	_, e := c.R().SetHeader("Content-Type", "application/json").SetBody(paramStr).SetResult(&res).Post(Api)
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	return res.Data
}

func handleStockDailyData_2(data model.ResData) (s []any) {
	for _, v := range data.Items {
		s = append(s, model.StockDailyData_2{
			Ts_Code:    v[0],
			Trade_Date: v[1],
			Open:       v[2],
			High:       v[3],
			Low:        v[4],
			Close:      v[5],
			Pre_Close:  v[6],
			Change:     v[7],
			Pct_Chg:    v[8],
			Vol:        v[9],
			Amount:     v[10],
		})
	}

	return s
}
