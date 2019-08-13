package adminm

import "time"

/**
加上冗余两个字段
*/
type User struct {
	ID int64
	//Id     string    `json:"id" xorm:"varchar(64) notnull unique 'id'"`
	Name   string    `json:"name"  `
	Nick   string    `json:"name" `
	Pwd    string    `json:"pwd"  `
	Phone  string    `json:"phone"`
	Status int       `json:"status"  `
	CTime  time.Time `json:"ctime"  `
	MTime  time.Time `json:"mtime"  `
	Did    string    `json:"did"  ` //冗余字段，为了查询方便
	Aid    string    `json:"adi"  `
}
