package main

import (
	"context"
	"fmt"
	"math"
	"similarStock/model"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var col *mongo.Collection

var sourceList = make([]model.StockDailyData, 0)
var pool = make(map[string][]model.StockDailyData)

const Period = 40

var result = make([]model.Pearson, 0)

func main() {
	connectMongo()
	getSourceStockData("002326.SZ")

	for i, v := range pool {
		value := calcPearsonCoefficient(sourceList, pool[i])
		result = append(result, model.Pearson{Ts_Code: v[0].Ts_Code, Value: value})
	}
	sort.SliceStable(result, func(i, j int) bool { return result[i].Value > result[j].Value })
	fmt.Println("top:", result[0:5])
	fmt.Println("tail:", result[len(result)-5:])
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
			if len(sourceList) >= Period {
				continue
			}
			sourceList = append(sourceList, item)
		} else {
			s, ok := pool[item.Ts_Code]
			if !ok {
				pool[item.Ts_Code] = []model.StockDailyData{item}
			} else {
				if len(s) >= Period {
					continue
				}
				pool[item.Ts_Code] = append(s, item)
			}
		}
		// fmt.Println(item.Ts_Code, item.Trade_Date)
	}

	fmt.Println("sourceList len:", len(sourceList), "pool len:", len(pool))

	return
}

func calcPearsonCoefficient(source, target []model.StockDailyData) float64 {
	size := len(source)
	if size != len(target) {
		fmt.Println("Error, Lengthes not match:", target[0].Ts_Code, len(source), len(target))
		return -1
	}

	sum_xy := 0.0
	sum_x := 0.0
	sum_y := 0.0
	sum_x2 := 0.0
	sum_y2 := 0.0

	for i, v := range source {
		x := v.Close
		y := target[i].Close
		sum_xy += x * y
		sum_x += x
		sum_y += y
		sum_x2 += x * x
		sum_y2 += y * y
	}

	value := (float64(size)*sum_xy - sum_x*sum_y) / (math.Sqrt(float64(size)*sum_x2-sum_x*sum_x) * math.Sqrt(float64(size)*sum_y2-sum_y*sum_y))
	return math.Ceil(value*100000) / 100000
}
