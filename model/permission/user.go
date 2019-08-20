package permission

import "time"

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

//UserAttr 用户属性,abac
type UserAttr struct {
	ID   int64
	Uid  string `gorm:"column:uid" form:"uid" json:"uid"`
	Attr string `gorm:"column:attr" form:"attr" json:"attr"` //数据库存json数据
}

//UserRole .
type UserRole struct {
	ID     int64
	Rid    string    `gorm:"column:rid" form:"rid" json:"rid"`
	Uid    string    `gorm:"column:uid" form:"uid" json:"uid"`
	Status int       `gorm:"column:status" form:"status" json:"status"`
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}

//UserSession  user session 用户session时长
type UserSession struct {
	ID      int64
	Uid     string `gorm:"column:uid" form:"uid" json:"uid"`
	Web     int    `gorm:"column:web" form:"web" json:"web"`
	Android int    `gorm:"column:android" form:"android" json:"android"`
	Ios     int    `gorm:"column:ios" form:"ios" json:"ios"`
}
