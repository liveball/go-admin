package permission

import "time"

/**
Id 用Did和Aid计算出来一个值，然后代表当前唯一id,这样方便查询在插入数据的时候查询是否存在

*/
type DomainApp struct {
	ID     int64
	Name   string    `gorm:"column:name" form:"name" json:"name"`
	Did    string    `gorm:"column:did" form:"did" json:"did"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}

//`gorm:"column:mtime" form:"mtime" json:"mtime"`
