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
	Ts_Code     string //	TS代码
	Symbol      string //	股票代码
	Name        string //	股票名称
	Area        string //	地域
	Industry    string //	所属行业
	Fullname    string //	股票全称
	Enname      string //	英文全称
	Cnspell     string //	拼音缩写
	Market      string //	市场类型（主板/创业板/科创板/CDR）
	Exchange    string //	交易所代码
	Curr_Type   string //	交易货币
	List_Status string //	上市状态 L上市 D退市 P暂停上市
	List_Date   string //	上市日期
	Delist_Date string //	退市日期
	Is_Hs       string //	是否沪深港通标的，N否 H沪股通 S深股通
}
