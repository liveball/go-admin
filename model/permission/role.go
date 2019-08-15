package permission

import "time"

/**
role

角色分可读 可写 读写
*/

type Role struct {
	ID     int64
	Name   string    `gorm:"column:name" form:"name" json:"name"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	Rw     int       `gorm:"column:rw" form:"rw" json:"rw"` //读写标志  1为读 2为写 3可读可写
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}
