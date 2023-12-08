package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Get() *sql.DB {
	dsn := "root:xl123456?@tcp(127.0.0.1:3306)/orm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 禁用默认事务
		PrepareStmt:            true, // 缓存预编译
	})
	if err != nil {
		fmt.Println("数据库连接错误:", err)
	}

	db, _ := gormDb.DB()
	return db
}
