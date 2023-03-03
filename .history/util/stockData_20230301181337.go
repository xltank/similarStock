package util

import (
	"net/http"
	"strings"
)

const Token = "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e"

/* type ApiBodyParams struct {
	Exchange   string `json:"exchange,omitempty"`
	Start_date string `json:"start_date,omitempty"`
	End_date   string `json:"end_date,omitempty"`
	Is_open    string `json:"is_open,omitempty"`
}

type ApiBody struct {
	Api_name string `json:"api_name,omitempty"`
	Token    string `json:"token,omitempty"`
	Fields   string `json:"fields,omitempty"`
	Params   map[string]string
} */

func test() {
	/* body := ApiBody{
		Api_name: "trade_cal",
		Token:    "xxxxxxxx",
		Params: map[string]string{
			"exchange":   "",
			"start_date": "20180901",
			"end_date":   "20181001",
			"is_open":    "0",
		},
		Fields: "exchange,cal_date,is_open,pretrade_date",
	} */

	params := `{"api_name": "trade_cal", "token": "$Token", "params": {"exchange":"", "start_date":"$start_date", "end_date":"$end_date", "is_open":"0"}, "fields": "exchange,cal_date,is_open,pretrade_date"}`

	http.Post("", "application/json", strings.NewReader(params))

}
