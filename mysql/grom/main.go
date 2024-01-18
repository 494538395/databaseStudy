package main

import (
	"context"
	"fmt"
	"time"

	"database-study/mysql/grom/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mysqlClient *gorm.DB

var (
	username      = "root"
	password      = "123456"
	address       = "localhost:3306"
	defaultDBName = "mysql"
	testDBName    = "jerry"
)

var globalCtx = context.Background()

func init() {
	// 1.连接默认库
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		username, password, address, defaultDBName)

	var err error
	preDB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	// 2.创建测试库
	sql := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS %s default charset utf8mb4 COLLATE utf8mb4_unicode_ci;",
		testDBName,
	)
	if err = preDB.Exec(sql).Error; err != nil {
		panic(err)
	}

	// 3.连接测试库
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, address, testDBName)

	mysqlClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	// 4.自动迁移
	mysqlClient.AutoMigrate(&model.Person{})

	sqlDB, err := mysqlClient.DB()
	if err != nil {
		panic(err)
	}
	// 5.设置连接池
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(10)
}

func main() {
	fmt.Println("MySQL started")

	personDB := model.NewPersonDB(mysqlClient)

	//personDB.CreateWithSelect(globalCtx, &model.Person{
	//	Name:    "jerry",
	//	Age:     21,
	//	Address: "xian",
	//}, []string{"name"})

	//personDB.CreateInBatch(globalCtx, []*model.Person{
	//	{Name: "sam"},
	//	{Name: "mike"},
	//	{Name: "tom"}},
	//)

	goCtx, _ := context.WithTimeout(globalCtx, time.Second)

	personDB.Create(goCtx, &model.Person{
		Name: "jerry",
		Age:  22,
	})

	fmt.Println("main process")

	select {}

}
