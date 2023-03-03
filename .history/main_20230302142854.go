package main

import (
	fmt
	similarStock/model

	github.com/go-resty/resty/v2
)

const Token = bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e
const Api = http://api.tushare.pro

var c = resty.New()

var stockBasicParam = `{api_name: stock_basic, token: bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e, params: {exchange:, list_status:L}, fields: ts_code,symbol,name,area,industry,list_date}`
var stockDailyParam = `{api_name: daily, token: bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e, params: {ts_code:600000.SH, start_date:20230201, end_date:20230302}, fields: ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount}`
var indexBasicParam = `{api_name: index_basic, token: bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e, params: {market: SSE}, fields: ts_code,name,fullname}`
var indexDailyParam = `{api_name: index_daily, token: bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e, params: {ts_code:000001.SH, start_date:20230301, end_date:20230302}, fields: ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount}`

func main() {
	// req(stockBasicParam)
	stockDaily := req(stockDailyParam)
	handleStockDailyData(stockDaily)

	// req(indexBasicParam)
	// req(indexDailyParam)
}

func req(paramStr string) model.ResData {
	res := model.Res{}
	_, e := c.R().SetHeader(Content-Type, application/json).SetBody(paramStr).SetResult(&res).Post(Api)
	if e != nil {
		fmt.Println(error:, e.Error())
	}
	fmt.Println(-------)
	fmt.Println(res)
	return res.Data
}

func handleStockDailyData(data model.ResData) (s []model.StockDailyData) {
	for _, v := range data.Items {
		index := 
		ts_code, _ := v[0].(string)
		trade_date, _ := v[1].(string)
		open, _ := v[2].(string)
		high, _ := v[3].(string)
		low, _ := v[4].(string)
		close, _ := v[5].(string)
		pre_close, _ := v[6].(string)
		change, _ := v[7].(string)
		pct_chg, _ := v[8].(string)
		vol, _ := v[9].(string)
		amount, _ := v[10].(string)
		s = append(s, model.StockDailyData{
			Ts_Code: ts_Code,
		})
	}

	return s
}
