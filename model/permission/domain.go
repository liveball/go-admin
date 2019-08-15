package permission

import "time"

/**
域
*/
type Domain struct {
	ID     int64
	Name   string    `gorm:"column:name" form:"name" json:"name"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}
