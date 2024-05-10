package db

import "gorm.io/gorm"

type Db struct {
	*gorm.DB
}

var dbInstance *Db

func GetCluster() *Db {
	return dbInstance
}

func SetCluster(db *gorm.DB) {
	dbInstance = &Db{db}
}
