package permission

import "time"

/**
domain app role关联关系(admin用)
*/

type DomainAppRole struct {
	ID     int64     `gorm:"column:id" form:"id" json:"id"`
	Did    string    `gorm:"column:did" form:"did" json:"did"`
	Aid    string    `gorm:"column:aid" form:"aid" json:"aid"`
	Rid    string    `gorm:"column:rid" form:"rid" json:"rid"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}
