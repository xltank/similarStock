package model

type Params struct {
	Is_hs       string `json:"is_hs,omitempty"`       //	是否沪深港通标的，N否 H沪股通 S深股通
	List_status string `json:"list_status,omitempty"` //	上市状态 L上市 D退市 P暂停上市，默认是L
	Exchange    string `json:"exchange,omitempty"`    //	交易所 SSE上交所 SZSE深交所 BSE北交所
	ts_code     string `json:"ts_code,omitempty"`     //	TS股票代码
	market      string `json:"market,omitempty"`      //	市场类别 （主板/创业板/科创板/CDR/北交所）
	limit       int    `json:"limit,omitempty"`       //
	offset      int    `json:"offset,omitempty"`      //
	name        string `json:"name,omitempty"`        //	名称
}

type Resp struct {
}
