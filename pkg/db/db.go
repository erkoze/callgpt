package db

import (
	"callgpt/configs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.DbConfig) *Db {
	db, err := gorm.Open(
		postgres.Open(buildDSN(conf)),
		&gorm.Config{},
	)

	if err != nil {
		panic(err.Error())
	}

	return &Db{db}
}

func buildDSN(conf *configs.DbConfig) string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		conf.Host,
		conf.User,
		conf.Password,
		conf.Name,
		conf.Port,
		conf.SslMode,
	)
}
