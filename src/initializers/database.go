package initializers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	// dbを作成
	// NOTE: "@tcp(XX)"の"XX"には、docker-compose.ymlのサービス名orコンテナ名が入る
	dsn := "root:password@tcp(db)/go_gin_project_db?charset=utf8mb4&parseTime=true&loc=Local"
	// DBインスタンスを初期化
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
