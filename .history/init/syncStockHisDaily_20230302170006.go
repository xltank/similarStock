package main

import (
	"context"
	"fmt"
	"similarStock/model"
	"time"

	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Token = "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e"
const Api = "http://api.tushare.pro"

var c = resty.New()

var allStockDailyParam = `{"api_name": "daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"trade_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`

func main() {

	client, e := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/test"))
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	col := client.Database("test").Collection("stockDaily")

	// req(stockBasicParam)
	// stockDaily := req(stockHisParam)
	stockDaily := req(allStockDailyParam)
	fmt.Println("-------")
	fmt.Println(stockDaily)
	stockDailyList := handleStockDailyData_2(stockDaily)
	fmt.Println("-------")
	fmt.Println(stockDailyList)
	// req(indexBasicParam)
	// req(indexDailyParam)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r, e := col.InsertMany(ctx, stockDailyList)
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	fmt.Println(r)
}

func req(paramStr string) model.ResData {
	res := model.Res{}
	_, e := c.R().SetHeader("Content-Type", "application/json").SetBody(paramStr).SetResult(&res).Post(Api)
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	// fmt.Println("-------")
	// fmt.Println(res)
	return res.Data
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

type Dictionary[K comparable, V any] map[K]V
