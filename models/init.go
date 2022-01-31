package models

import (
	"time"

	"github.com/coding-and-typing/xhup-club-api-go/internal/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"

	"github.com/rs/zerolog/log"
)

// DB ...
var DB *gorm.DB

// Database ...
func Database(dbConfig *config.DBConfig) error {
	log.Info().Msgf("connect to db: {%#v}", dbConfig.Uri)
	db, err := gorm.Open(sqlite.Open(dbConfig.Uri), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Silent),
	})
	if err != nil {
		log.Error().Msgf("conect db err:%v", err)
		return err
	}
	//设置连接池
	//空闲
	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Msgf("get sqlDB err:%v", err)
		return err
	}

	//打开
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
	//超时
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(dbConfig.ConnMaxLifeTime))

	DB = db
	migration()
	return nil
}
