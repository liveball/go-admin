package users

import "time"

/**
域  用字符串id目的是MD5后查询id方便,中文名字查询是否存在有可能会出问题？
*/
type Domain struct {
	Id     string    `json:"id" gorm:"type:varchar(64) notnull unique 'id'"`
	Name   string    `json:"name" xorm:"varchar(256) notnull unique 'name'"`
	Status int       `json:"status" xorm:"default -1"`
	CTime  time.Time `json:"ctime" xorm:"created"`
	MTime  time.Time `json:"mtime" xorm:"updated" `
}
