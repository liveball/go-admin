package users

// User  for user
type User struct {
	ID     int64  `gorm:"column:id"    form:"id" json:"id"`
	Name   string `gorm:"column:name"  form:"name" json:"name"`
	Nick   string `json:"name" gorm:"type:varchar(100) "`
	Pwd    string `json:"pwd" xorm:"type:varchar(64)"`
	Phone  string `json:"phone" gorm:"column:phone"`
	Status int    `json:"status" gorm:"status"`
	Did    string `json:"did" gorm:"type:varchar(64) "` //冗余字段，为了查询方便
	Aid    string `json:"aid" gorm:"type:varchar(64) "`
}
