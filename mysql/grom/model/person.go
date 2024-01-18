package model

type Person struct {
	Id      int    `gorm:"column:Id;primaryKey;autoIncrement"`
	Name    string `gorm:"column:name;type:varchar(255);uniqueIndex:unique01"`
	Age     int    `gorm:"column:age;uniqueIndex:unique01"`
	Address string `gorm:"column:address"`
}

func (p *Person) TableName() string {
	return "person"
}
