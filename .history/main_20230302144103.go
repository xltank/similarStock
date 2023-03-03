package main

import (
	"fmt"
	"similarStock/model"

	"github.com/go-resty/resty/v2"
)

const Token = "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e"
const Api = "http://api.tushare.pro"

var c = resty.New()

var stockBasicParam = `{"api_name": "stock_basic", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"exchange":"", "list_status":"L"}, "fields": "ts_code,symbol,name,area,industry,list_date"}`
var stockDailyParam = `{"api_name": "daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"ts_code":"600000.SH", "start_date":"20230201", "end_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`
var indexBasicParam = `{"api_name": "index_basic", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"market": "SSE"}, "fields": "ts_code,name,fullname"}`
var indexDailyParam = `{"api_name": "index_daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"ts_code":"000001.SH", "start_date":"20230301", "end_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`

func main() {
	// req(stockBasicParam)
	stockDaily := req(stockDailyParam)
	stockDailyList := handleStockDailyData(stockDaily)
	fmt.Println("-------")
	fmt.Println(stockDailyList)
	// req(indexBasicParam)
	// req(indexDailyParam)
}

func req(paramStr string) model.ResData {
	res := model.Res{}
	_, e := c.R().SetHeader("Content-Type", "application/json").SetBody(paramStr).SetResult(&res).Post(Api)
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	fmt.Println("-------")
	fmt.Println(res)
	return res.Data
}

func handleStockDailyData(data model.ResData) (s []model.StockDailyData) {
	for _, v := range data.Items {
		ts_code, _ := v[0].(string)
		trade_date, _ := v[1].(string)
		open, _ := v[0].(float32)
		high, _ := v[1].(float32)
		low, _ := v[2].(float32)
		close, _ := v[3].(float32)
		pre_close, _ := v[4].(float32)
		change, _ := v[5].(float32)
		pct_chg, _ := v[6].(float32)
		vol, _ := v[7].(float32)
		amount, _ := v[8].(float32)
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
