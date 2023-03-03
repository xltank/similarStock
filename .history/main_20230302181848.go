package main

import (
	"context"
	"fmt"
	"similarStock/model"

	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var c = resty.New()

// var stockBasicParam = `{"api_name": "stock_basic", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"exchange":"", "list_status":"L"}, "fields": "ts_code,symbol,name,area,industry,list_date"}`
// var stockHisParam = `{"api_name": "daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"ts_code":"600000.SH", "start_date":"20230201", "end_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`
// var allStockDailyParam = `{"api_name": "daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"trade_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`
// var indexBasicParam = `{"api_name": "index_basic", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"market": "SSE"}, "fields": "ts_code,name,fullname"}`
// var indexDailyParam = `{"api_name": "index_daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"ts_code":"000001.SH", "start_date":"20230301", "end_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`

func main() {
	client, e := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/test"))
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	col := client.Database("test").Collection("stockDaily")

}

func handleStockDailyData(data model.ResData) (s []model.StockDailyData) {
	for _, v := range data.Items {
		ts_code, _ := v[0].(string)
		trade_date, _ := v[1].(string)
		open, _ := v[2].(float64)
		high, _ := v[3].(float64)
		low, _ := v[4].(float64)
		close, _ := v[5].(float64)
		pre_close, _ := v[6].(float64)
		change, _ := v[7].(float64)
		pct_chg, _ := v[8].(float64)
		vol, _ := v[9].(float64)
		amount, _ := v[10].(float64)
		s = append(s, model.StockDailyData{
			Ts_Code:    ts_code,
			Trade_Date: trade_date,
			Open:       open,
			High:       high,
			Low:        low,
			Close:      close,
			Pre_Close:  pre_close,
			Change:     change,
			Pct_Chg:    pct_chg,
			Vol:        vol,
			Amount:     amount,
		})
	}

	return s
}
