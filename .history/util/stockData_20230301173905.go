package util

import "net/http"

const Token = "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e"

func test() {
	body := map[string]string{"api_name": "trade_cal", "token": "xxxxxxxx", "params": map[string]string{"exchange": "", "start_date": "20180901", "end_date": "20181001", "is_open": "0"}, "fields": "exchange,cal_date,is_open,pretrade_date"}
	http.Post("", "application/json")

}
