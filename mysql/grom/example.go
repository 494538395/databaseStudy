package main

import (
	"fmt"

	"gorm.io/gorm"
)

// region  Create

// 方法用于查找匹配条件的记录，如果找不到则创建一条新的记录。
func FirstOrCreate() {
	mysqlClient.FirstOrCreate(&Student{}, &Student{Name: "jack"})

	fmt.Println("FirstOrCreate finished")
}

func Create() {
	mysqlClient.Create(&Student{Name: "jack", Age: 17})
	// INSERT INTO `students` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`address`)
	// VALUES ('2024-01-04 10:25:43.915','2024-01-04 10:25:43.915',NULL,'jack',17,'')
	fmt.Println("Create finished")
}

func Save() {
	mysqlClient.Save(&Student{Model: gorm.Model{ID: 1}, Name: "jack", Age: 23})

	fmt.Println("Save finished")

}

func CreateWithSelect() {
	s := &Student{
		Name:    "mike",
		Age:     21,
		Address: "beijing",
	}

	mysqlClient.Select("Name").Create(s)
	// INSERT INTO `students` (`created_at`,`updated_at`,`name`)
	// VALUES ('2024-01-04 10:25:43.919','2024-01-04 10:25:43.919','mike')
	fmt.Println("CreateWithSelect finished")
}

func CreateInBatch() {
	students := []Student{
		{Name: "john"},
		{Name: "tom"},
		{Name: "sam"},
	}

	mysqlClient.CreateInBatches(students, 3)
	/*
	   INSERT INTO `students` (`created_at`,`updated_at`,`deleted_at`,`name`,`age`,`address`)
	   VALUES ('2024-01-04 10:29:45.828','2024-01-04 10:29:45.828',NULL,'john',0,''),
	          ('2024-01-04 10:29:45.828','2024-01-04 10:29:45.828',NULL,'tom',0,''),
	          ('2024-01-04 10:29:45.828','2024-01-04 10:29:45.828',NULL,'sam',0,'')

	*/

	fmt.Println("CreateInBatch finished")
}

func CreateWithMap() {
	mysqlClient.Model(&Student{}).Create(map[string]interface{}{
		"Name": "kid",
		"age":  25,
	})

	fmt.Println("CreateWithMap finished")

}

// endregion  Create

// region Select

func FirstByID() {
	res := mysqlClient.First(&Student{}, 2)
	if res.Error == gorm.ErrRecordNotFound {
		fmt.Println("record not found")
	} else if res.Error != nil {
		panic(res.Error)
	} else {
		var stu Student
		res.Scan(&stu)
		//fmt.Println(stu)
	}
	/*
	   SELECT * FROM `students`
	   WHERE `students`.`id` = 2
	   AND `students`.`deleted_at` IS NULL
	   ORDER BY `students`.`id` LIMIT 1
	*/

	fmt.Println("First finished")
}

func FirstByExpr() {
	res := mysqlClient.First(&Student{}, "age = ?", "25")
	if res.Error == gorm.ErrRecordNotFound {
		fmt.Println("record not found")
	} else if res.Error != nil {
		panic(res.Error)
	} else {
		var stu Student
		res.Scan(&stu)
		fmt.Println(stu)
	}

	/*
		    SELECT * FROM `students`
		   	WHERE age = '25'
		   	AND `students`.`deleted_at` IS NULL
			AND `students`.`id` = 10
		   	ORDER BY `students`.`id` LIMIT 1
	*/
	fmt.Println("FirstByExpr finished")
}

func Take() {
	res := mysqlClient.Take(&Student{})
	if res.Error == gorm.ErrRecordNotFound {
		fmt.Println("record not found")
	} else if res.Error != nil {
		panic(res.Error)
	} else {
		var stu Student
		res.Scan(&stu)
		fmt.Println(stu)
	}

	/*
		SELECT * FROM `students`
		WHERE `students`.`deleted_at` IS NULL
		AND `students`.`id` = 1 LIMIT 1
	*/

	fmt.Println("Task finished")
}

func Find() {
	res := mysqlClient.Find(&[]Student{})
	if res.Error == gorm.ErrRecordNotFound {
		fmt.Println("record not found")
	} else if res.Error != nil {
		panic(res.Error)
	} else {
		var stus []Student
		res.Scan(&stus)
		fmt.Println(res)
	}

	/*
		SELECT * FROM `students` WHERE `students`.`deleted_at` IS NULL
	*/
	fmt.Println("Find finished")
}

// endregion

// region Delete

func Delete() {
	mysqlClient.Delete(&Student{}, "age = ?", 25)

	// 如果 students 存在 deleted_at 字段，则执行：
	/*
	   UPDATE `students`
	   SET `deleted_at`='2024-01-04 11:22:27.196' WHERE age = 25
	   AND `students`.`deleted_at` IS NULL
	*/

	// 如果 students 表没有 deleted_at 字段，则执行：
	/*
		DELETE FROM `students` WHERE age = 25
	*/
	fmt.Println("Delete finished")
}

// endregion
