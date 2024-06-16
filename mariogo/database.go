package mariogo

import (
	"os"

	"github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnect() {
	db, err := gorm.Open(gormmysql.New(gormmysql.Config{
		DSNConfig: &mysql.Config{
			User:      os.Getenv("DB_USER"),
			Net:       Getenv("DB_NET", "tcp"),
			Addr:      os.Getenv("DB_HOST") + ":" + Getenv("DB_PORT", "3306"),
			DBName:    os.Getenv("DB_NAME"),
			ParseTime: true,
		},
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
