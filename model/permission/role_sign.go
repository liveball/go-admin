package permission

/**
Role签名token签名表,用于对比权限变化  后台更新权限，对应更新这个角色的Sign
*/

type RoleSign struct {
	ID   int64
	Uid  string `gorm:"column:uid" form:"uid" json:"uid"`
	Sign string `gorm:"column:sign" form:"sign" json:"sign"`
}
