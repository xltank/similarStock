package main

import (
	"fmt"
	"similarStock/model"

	"github.com/go-resty/resty/v2"
)

const Token = "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e"
const Api = "http://api.tushare.pro"

var c = resty.New()

func main() {
	// getStockList()
	getStockData()
}

func getStockList() {
	params := `{"api_name": "stock_basic", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"exchange":"", "list_status":"L"}, "fields": "ts_code,symbol,name,area,industry,list_date"}`
	res := model.Res{}
	_, e := c.R().SetHeader("Content-Type", "application/json").SetBody(params).SetResult(&res).Post(Api)
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	fmt.Println(res)
}

func getStockData() {
	params := `{"api_name": "daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"ts_code":"000001.SZ", "start_date":"20230301", "end_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`
	res := model.Res{}
	r, e := c.R().SetHeader("Content-Type", "application/json").SetBody(params).SetResult(&res).Post(Api)
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	fmt.Println(string(r.Body()))
	// fmt.Println(res)
}
