package util

import "net/http"

const Token = "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e"

type ApiBodyParams struct {
	Exchange   string
	Start_date string
	End_date   string
	Is_open    string
}

type ApiBody struct {
	Api_name string
	Token    string
	Fields   string
	Params   map[string]string
}

func test() {
	body := ApiBody{
		Api_name: "trade_cal",
		Token:    "xxxxxxxx",
		Params: map[string]string{
			"exchange":   "",
			"start_date": "20180901",
			"end_date":   "20181001",
			"is_open":    "0",
		},
		Fields: "exchange,cal_date,is_open,pretrade_date",
	}
	http.Post("", "application/json")

}
