package model

type Params struct {
	is_hs       string //	是否沪深港通标的，N否 H沪股通 S深股通
	list_status string //	上市状态 L上市 D退市 P暂停上市，默认是L
	exchange    string //	交易所 SSE上交所 SZSE深交所 BSE北交所
	ts_code     string //	TS股票代码
	market      string //	市场类别 （主板/创业板/科创板/CDR/北交所）
	limit       int    //
	offset      int    //
	name        string //	名称
}

type Resp struct {
}
