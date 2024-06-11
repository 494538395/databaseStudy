package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 设置 MongoDB 连接选项
	clientOptions := options.Client().ApplyURI("mongodb://10.0.1.84:38077")

	// 连接到 MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// 选择要操作的数据库和集合
	db := client.Database("89t_im")

	coll := db.Collection("group_info_250") // 获取 collection 实例

	// 设置过滤条件
	filter := bson.M{"group_id": 7}

	// 设置更新操作
	//fieldValue := map[string]interface{}{
	//	"test": "jerry",
	//	"ex": map[string]interface{}{
	//		"name": "jack",
	//	},
	//}

	fieldValue := map[string]interface{}{
		"test":       "jerry",
		"ex.field01": "value01",
		"ex.field02": "value02",
	}

	updateCol := bson.M{"$set": fieldValue}

	// 执行更新操作
	result := coll.FindOneAndUpdate(context.Background(), filter, updateCol)
	if result.Err() != nil {
		log.Fatal(err)
	}

	// 输出更新结果
	fmt.Println("更新完成")
}
