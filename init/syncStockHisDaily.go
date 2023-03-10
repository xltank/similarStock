package main

/*
Sync daily data of all stocks in last 100 days.


// var stockBasicParam = `{"api_name": "stock_basic", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"exchange":"", "list_status":"L"}, "fields": "ts_code,symbol,name,area,industry,list_date"}`
// var stockHisParam = `{"api_name": "daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"ts_code":"600000.SH", "start_date":"20230201", "end_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`
// var allStockDailyParam = `{"api_name": "daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"trade_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`
// var indexBasicParam = `{"api_name": "index_basic", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"market": "SSE"}, "fields": "ts_code,name,fullname"}`
// var indexDailyParam = `{"api_name": "index_daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"ts_code":"000001.SH", "start_date":"20230301", "end_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`
*/

import (
	"context"
	"fmt"
	"similarStock/model"
	"strings"
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

var startDateStr = "20230303"
var endDateStr = ""

var c = resty.New()

var col *mongo.Collection

func main() {

	client, e := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/test"))
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	col = client.Database("test").Collection("stockDaily")

	startDate, _ := time.Parse("20060102", startDateStr)
	endDate := time.Now()
	if endDateStr != "" {
		endDate, _ = time.Parse("20060102", endDateStr)
	}

	for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
		dateStr := d.Format("20060102")
		fmt.Println(dateStr)
		p := strings.Replace(allStockDailyParam, "$date", dateStr, -1)
		stockDaily := req(p)
		fmt.Println("daily data length:", dateStr, len(stockDaily.Items))
		if len(stockDaily.Items) == 0 {
			continue
		}
		stockDailyList := handleStockDailyData_2(stockDaily)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		r, e := col.InsertMany(ctx, stockDailyList)
		if e != nil {
			fmt.Println("error:", e.Error())
		}
		if r != nil {
			fmt.Println("Inserted doc count:", len(r.InsertedIDs))
		}
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
