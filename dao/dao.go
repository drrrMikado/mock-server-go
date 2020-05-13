package dao

import (
	"github.com/drrrMikado/mock-server-go/conf"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	c  *conf.Config
	db *gorm.DB
}

// New init mysql db
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		c: c,
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	_ = d.db.Close()
}
