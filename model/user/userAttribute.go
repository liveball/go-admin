package adminm

/**
用户属性,abac
*/

type UserAttr struct {
	ID   int64
	Uid  string `json:"uid" `
	Attr string `json:"attr"` //数据库存json数据
}
