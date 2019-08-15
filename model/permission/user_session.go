package permission

/**
user session 用户session时长
*/

type UserSession struct {
	ID      int64
	Uid     string `gorm:"column:uid" form:"uid" json:"uid"`
	Web     int    `gorm:"column:web" form:"web" json:"web"`
	Android int    `gorm:"column:android" form:"android" json:"android"`
	Ios     int    `gorm:"column:ios" form:"ios" json:"ios"`
}

//`gorm:"column:mtime" form:"mtime" json:"mtime"`
