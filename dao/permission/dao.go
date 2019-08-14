package permission

import (
	"context"

	"go-admin/conf"
	"go-admin/lib"

	"github.com/jinzhu/gorm"
)

//Dao for user dao
type Dao struct {
	db *gorm.DB
}

// New init dao
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		db: lib.NewMySQL(c.ORM),
	}
	d.initORM()
	return
}

func (d *Dao) initORM() {
	d.db.LogMode(true)
}

// Ping check connection of db , mc.
func (d *Dao) Ping(c context.Context) (err error) {
	if d.db != nil {
		err = d.db.DB().PingContext(c)
	}
	return
}

// Close db
func (d *Dao) Close() (err error) {
	return d.db.Close()
}
