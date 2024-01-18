package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type PersonGorm struct {
	*MetaDB
}

func NewPersonDB(db *gorm.DB) *PersonGorm {
	return &PersonGorm{NewMetaDB(db, &Person{})}
}

// Create 根据 struct 进行 insert
func (p *PersonGorm) Create(ctx context.Context, person *Person) error {

	time.Sleep(3 * time.Second)

	fmt.Println("Create finished")

	res := p.db(ctx).Create(person)
	return res.Error
}

// CreateWithSelect 根据选中的字段进行 insert
func (p *PersonGorm) CreateWithSelect(ctx context.Context, person *Person, fields []string) error {
	res := p.db(ctx).Select(fields).Create(person)
	return res.Error
}

// CreateInBatch 批量创建
func (p *PersonGorm) CreateInBatch(ctx context.Context, args []*Person) error {
	res := p.db(ctx).CreateInBatches(args, len(args))
	return res.Error
}

// CreateWithMap 根据 map 创建
func (p *PersonGorm) CreateWithMap(ctx context.Context, param map[string]interface{}) error {
	res := p.db(ctx).Create(param)
	return res.Error
}
