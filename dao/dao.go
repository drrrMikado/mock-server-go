package dao

import (
	"github.com/drrrMikado/mock-server-go/conf"
	"github.com/drrrMikado/mock-server-go/database/orm"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	c  *conf.Config
	db *gorm.DB
}

// New init mysql db
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c: c,
	}
	d.db = orm.NewMySQL(c.Mysql)
	return
}

// Close close the resource.
func (d *Dao) Close() {
	_ = d.db.Close()
}
