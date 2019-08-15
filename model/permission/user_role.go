package permission

import "time"

/**
user role
*/

type UserRole struct {
	ID     int64
	Rid    string    `gorm:"column:rid" form:"rid" json:"rid"`
	Uid    string    `gorm:"column:uid" form:"uid" json:"uid"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}
