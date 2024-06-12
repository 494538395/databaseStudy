package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database   = "nut-test"
	collection = "u_level"

	col *mongo.Collection
)

func init() {
	// 设置 MongoDB 客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://10.0.1.84:40077")

	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!")

	// 选择数据库和集合
	col = client.Database(database).Collection(collection)
}

func main() {

	deleteDoc("2001")

}

func updateDoc(docId string, updateFields map[string]interface{}) {
	// 定义要更新的文档过滤条件
	filter := bson.M{"_id": docId}

	// 在更新之前，查询并打印文档，以确保过滤条件正确匹配
	var result bson.M
	err := col.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatalf("Failed to find document: %v", err)
	}
	fmt.Printf("Document before update: %+v\n", result)

	// 定义更新操作
	//update := bson.M{
	//	"$set": bson.M{
	//		"maxFinishLevelId": 0,
	//	},
	//}

	// 将要更新的字段添加到更新操作中
	update := bson.M{"$set": updateFields}

	// 执行更新操作
	updateResult, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatalf("Failed to update document: %v", err)
	}

	fmt.Println("updateResult-->", updateResult)
}

func deleteDoc(docId string) {
	// 定义要更新的文档过滤条件
	filter := bson.M{"_id": docId}

	delResult, err := col.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("delResult.DeletedCount-->", delResult.DeletedCount)

}
