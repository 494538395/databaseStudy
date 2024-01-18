package model

import (
	"context"

	"gorm.io/gorm"
)

type MetaDB struct {
	DB    *gorm.DB
	table any
}

func NewMetaDB(db *gorm.DB, table any) *MetaDB {
	return &MetaDB{
		DB:    db,
		table: table,
	}
}

func (m *MetaDB) db(ctx context.Context) *gorm.DB {
	return m.DB.WithContext(ctx).Model(m.table)
}
