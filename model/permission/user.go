package permission

//TableName get table name
func (u *User) TableName() string {
	return "user"
}

//User def user.
type User struct {
	//自增id 添加不用传 编辑传
	ID    int64  `form:"id" json:"id"`
	State int8   `gorm:"column:state" form:"state" json:"state"`
	Name  string `gorm:"column:name" form:"name" json:"name"`
	Nick  string `gorm:"column:nick" form:"nick" json:"nick"`
	Pwd   string `gorm:"column:pwd" form:"pwd" json:"pwd"`
	Phone string `gorm:"column:phone" form:"phone" json:"phone"`
	//创建时间 不用传
	CTime string `gorm:"column:ctime" form:"ctime" json:"ctime"`
	//修改时间 不用传
	MTime string `gorm:"column:mtime" form:"mtime" json:"mtime"`

	//加上冗余两个字段
	//DID string `gorm:"column:did" form:"did" json:"did"`
	//AID string `gorm:"column:aid" form:"aid" json:"aid"`
}
