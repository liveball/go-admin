package permission

import "time"

//Domain .
type Domain struct {
	ID     int64
	Name   string    `gorm:"column:name" form:"name" json:"name"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}

// DomainAppId 用Did和Aid计算出来一个值，然后代表当前唯一id,这样方便查询在插入数据的时候查询是否存在
type DomainApp struct {
	ID     int64
	Name   string    `gorm:"column:name" form:"name" json:"name"`
	Did    string    `gorm:"column:did" form:"did" json:"did"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}

//DomainAppRole domain app role关联关系(admin用)

type DomainAppRole struct {
	ID     int64     `gorm:"column:id" form:"id" json:"id"`
	Did    string    `gorm:"column:did" form:"did" json:"did"`
	Aid    string    `gorm:"column:aid" form:"aid" json:"aid"`
	Rid    string    `gorm:"column:rid" form:"rid" json:"rid"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}

// DomainAppUser domain app user关联关系表(admin用)
type DomainAppUser struct {
	ID     int64
	Did    string    `gorm:"column:did" form:"did" json:"did"`
	Aid    string    `gorm:"column:aid" form:"aid" json:"aid"`
	Uid    string    `gorm:"column:uid" form:"uid" json:"uid"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}
