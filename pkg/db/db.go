package db

import (
	"fmt"
	"go-gin/database/model"
	"go-gin/pkg/secret"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var G *gorm.DB

var Models = []interface{}{
	&model.User{},
	&model.Contact{},
	&model.Address{},
}

func Setup() {
	dsn := fmt.Sprintf(
		`host=%v user=%v password="%v" dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta`,
		secret.DB_HOST,
		secret.DB_USER,
		secret.DB_PASS,
		secret.DB_NAME,
		secret.DB_PORT,
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().Local() // timestamps
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	G = conn
}
