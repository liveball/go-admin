package user

// User  for user
type User struct {
	ID    int64  `gorm:"column:id"    form:"id" json:"id"`
	Name  int64  `gorm:"column:name"  form:"name" json:"name"`
	CTime string `gorm:"column:ctime" form:"ctime" json:"-"`
	MTime string `gorm:"column:mtime" form:"mtime" json:"-"`
}
