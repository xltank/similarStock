package model

type Params struct {
	Is_hs       string `json:"is_hs,omitempty"`       //	是否沪深港通标的，N否 H沪股通 S深股通
	List_status string `json:"list_status,omitempty"` //	上市状态 L上市 D退市 P暂停上市，默认是L
	Exchange    string `json:"exchange,omitempty"`    //	交易所 SSE上交所 SZSE深交所 BSE北交所
	Ts_code     string `json:"ts_code,omitempty"`     //	TS股票代码
	Market      string `json:"market,omitempty"`      //	市场类别 （主板/创业板/科创板/CDR/北交所）
	Limit       int    `json:"limit,omitempty"`       //
	Offset      int    `json:"offset,omitempty"`      //
	Name        string `json:"name,omitempty"`        //	名称
}

type Resp struct {
	ts_code     str //	TS代码
	symbol      str //	股票代码
	name        str //	股票名称
	area        str //	地域
	industry    str //	所属行业
	fullname    str //	股票全称
	enname      str //	英文全称
	cnspell     str //	拼音缩写
	market      str //	市场类型（主板/创业板/科创板/CDR）
	exchange    str //	交易所代码
	curr_type   str //	交易货币
	list_status str //	上市状态 L上市 D退市 P暂停上市
	list_date   str //	上市日期
	delist_date str //	退市日期
	is_hs       str //	是否沪深港通标的，N否 H沪股通 S深股通
}
