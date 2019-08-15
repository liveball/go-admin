package permission

/**
用户属性,abac
*/

type UserAttr struct {
	ID   int64
	Uid  string `gorm:"column:uid" form:"uid" json:"uid"`
	Attr string `gorm:"column:attr" form:"attr" json:"attr"` //数据库存json数据
}
