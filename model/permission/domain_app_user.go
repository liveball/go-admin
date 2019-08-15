package permission

import "time"

/**
domain app user关联关系表(admin用)
*/

type DomainAppUser struct {
	ID     int64
	Did    string    `gorm:"column:did" form:"did" json:"did"`
	Aid    string    `gorm:"column:aid" form:"aid" json:"aid"`
	Uid    string    `gorm:"column:uid" form:"uid" json:"uid"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}
