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
	getStockList()
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

func getStockData(codes, startDate, endDate string) {
	params := `{"api_name": "daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {ts_code="000001.SZ", start_date="20180701", end_date="20180718"}, "fields": "ts_code,symbol,name,area,industry,list_date"}`
	res := model.Res{}
	_, e := c.R().SetHeader("Content-Type", "application/json").SetBody(params).SetResult(&res).Post(Api)
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	fmt.Println(res)
}


ts_code	    str	股票代码
trade_date	str	交易日期
open	float	开盘价
high	float	最高价
low	float	最低价
close	float	收盘价
pre_close	float	昨收价(前复权)
change	float	涨跌额
pct_chg	float	涨跌幅 （未复权，如果是复权请用 通用行情接口 ）
vol	float	成交量 （手）
amount	float	成交额 （千元）