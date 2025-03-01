package db

import (
	"context"
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
)

type DataAccess struct {
	ctx context.Context
	db  *gorm.DB
}

func NewDataAccess(ctx context.Context, db *gorm.DB) *DataAccess {
	once.Do(initDatabases(db))
	return &DataAccess{
		ctx: ctx,
		db:  db,
	}
}

func initDatabases(db *gorm.DB) func() {
	return func() {
		if err := db.AutoMigrate(&userModel); err != nil {
			panic(err)
		}
		if err := db.AutoMigrate(&commnutiyModel); err != nil {
			panic(err)
		}
		if err := db.AutoMigrate(&postModel); err != nil {
			panic(err)
		}
	}
}
