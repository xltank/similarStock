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

var sourceList = make([]model.StockDailyData, 0)
var pool = make(map[string][]model.StockDailyData)

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
	options := options.Find().SetSort(bson.D{{"ts_code", 1}, {"trade_date", -1}})
	cursor, e := col.Find(ctx, bson.M{"trade_date": bson.M{"$gte": "20221201"}}, options)
	if e != nil {
		fmt.Println("Get data error:", e.Error())
	}

	for cursor.Next(context.TODO()) {
		item := model.StockDailyData{}
		cursor.Decode(&item)
		if item.Ts_Code == code {
			sourceList = append(sourceList, item)
		} else {
			s, ok := pool[item.Ts_Code]
			if !ok {
				pool[item.Ts_Code] = []model.StockDailyData{item}
			} else {
				pool[item.Ts_Code] = append(s, item)
			}
		}
		// fmt.Println(item.Ts_Code, item.Trade_Date)
	}

	fmt.Println("sourceList len:", len(sourceList), "pool len:", len(pool))

	return
}
