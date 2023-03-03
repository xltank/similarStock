package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

const Token = "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e"
const Api = "http://api.tushare.pro"

func getStockList() {

	params := `{"api_name": "stock_basic", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"exchange":"", "list_status":"L"}, "fields": "ts_code,symbol,name,area,industry,list_date"}`

	/* r, e := http.Post(Api, "application/json", strings.NewReader(params))
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	body, e := io.ReadAll(r.Body)
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	fmt.Println(string(body)) */

	resty.Client().Post()
}

func main() {
	getStockList()
}
