package permission

import "time"

/**
role  perms
*/

type RolePerms struct {
	ID     int64
	Rid    string    `gorm:"column:rid" form:"rid" json:"rid"`
	Pid    string    `gorm:"column:pid" form:"pid" json:"pid"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}

//`gorm:"column:mtime" form:"mtime" json:"mtime"`
