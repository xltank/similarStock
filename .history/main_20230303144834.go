package main

import (
	"context"
	"fmt"
	"similarStock/model"
	"time"

	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var c = resty.New()
var col *mongo.Collection

// var stockBasicParam = `{"api_name": "stock_basic", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"exchange":"", "list_status":"L"}, "fields": "ts_code,symbol,name,area,industry,list_date"}`
// var stockHisParam = `{"api_name": "daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"ts_code":"600000.SH", "start_date":"20230201", "end_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`
// var allStockDailyParam = `{"api_name": "daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"trade_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`
// var indexBasicParam = `{"api_name": "index_basic", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"market": "SSE"}, "fields": "ts_code,name,fullname"}`
// var indexDailyParam = `{"api_name": "index_daily", "token": "bf0569836c92efcb6c52aa66063bbb923dbf958e54eb1431894db06e", "params": {"ts_code":"000001.SH", "start_date":"20230301", "end_date":"20230302"}, "fields": "ts_code,trade_date,open,high,low,close,pre_close,change,pct_chg,vol,amount"}`

func main() {
	connectMongo()
	getSourceStockData("002326.SZ")
}

func connectMongo() {
	client, e := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/test"))
	if e != nil {
		fmt.Println("error:", e.Error())
	}
	col = client.Database("test").Collection("stockDaily")
}

func getSourceStockData(code string) (data []model.StockDailyData) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	cursor, e := col.Find(ctx, bson.M{"trade_date": bson.M{"$gte": "20230101"}})
	if e != nil {
		fmt.Println("Get data error:", e.Error())
	}

	for cursor.Next(context.TODO()) {
		item := model.StockDailyData{}
		cursor.Decode(&item)
		fmt.Println(item)
	}

	return
}
