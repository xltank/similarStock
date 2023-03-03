package main

import "golang.org/x/tools/godoc/util"

import (
	"fmt"
	"net/http"
	"strings"
)

const Token = "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e"

func test() {

	params := `{"api_name": "stock_basic", "token": "$Token", "params": {"exchange":"", "list_status":"L"}, "fields": "exchange,cal_date,is_open,pretrade_date"}`

	r, e := http.Post("", "application/json", strings.NewReader(params))
	if e != nil {
		fmt.Println("error:", e.Error())
	}

	fmt.Println(r)
}


func main() {
	util.
}
