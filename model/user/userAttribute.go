package users

/**
用户属性,abac
*/

type UserAttr struct {
	Uid  string `json:"uid" xorm:"varchar(64) "`
	Attr string `json:"arrt"` //数据库存json数据
}
