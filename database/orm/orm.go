package orm

import (
	time2 "time"

	"github.com/drrrMikado/mock-server-go/internal/time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Config struct {
	Addr        string
	DSN         string
	Active      int
	Idle        int
	IdleTimeout time.Duration
}

func NewMySQL(c *Config) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time2.Duration(c.IdleTimeout) / time2.Second)
	return
}
