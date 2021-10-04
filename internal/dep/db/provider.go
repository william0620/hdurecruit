package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func ProvideDb(dbName string) (*gorm.DB,error) {
	db,ok := s.DbMap[dbName]
	if !ok {
		return  nil,errors.New("no this db")
	}
	return db,nil
}
