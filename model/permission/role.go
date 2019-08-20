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

//RolePerms .
type RolePerms struct {
	ID     int64
	Rid    string    `gorm:"column:rid" form:"rid" json:"rid"`
	Pid    string    `gorm:"column:pid" form:"pid" json:"pid"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}

//RoleSign Role签名token签名表,用于对比权限变化  后台更新权限，对应更新这个角色的Sign
type RoleSign struct {
	ID   int64
	Uid  string `gorm:"column:uid" form:"uid" json:"uid"`
	Sign string `gorm:"column:sign" form:"sign" json:"sign"`
}
