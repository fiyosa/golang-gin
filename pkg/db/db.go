package db

import (
	"fmt"
	"go-gin/pkg/secret"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var G *gorm.DB

var Models = []interface{}{
	&User{},
	&Contact{},
	&Address{},
}

func Setup() {
	dsn := fmt.Sprintf(
		`host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Jakarta`,
		secret.DB_HOST,
		secret.DB_USER,
		secret.DB_PASS,
		secret.DB_NAME,
		secret.DB_PORT,
		secret.DB_SSLMODE,
	)

	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: secret.DB_SCHEMA + ".",
		},
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().Local() // timestamps
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	G = connect
}
