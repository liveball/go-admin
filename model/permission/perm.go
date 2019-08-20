package permission

import "time"

/**
perms
设计做成一级目录 二级目录 三级目录
	api
api 分 读 写
这个是针对角色的
*/

//Menu .
type Menu struct {
	Id         string    `gorm:"column:id" form:"id" json:"id"`                            //主键id
	Category   int       `gorm:"column:category" form:"category" json:"category"`          //类型  1为目录 2为api
	Name       string    `gorm:"column:name" form:"name" json:"name"`                      //名称
	Icon       string    `gorm:"column:icon" form:"icon" json:"icon"`                      //目录的icon
	ParentId   string    `gorm:"column:parent_id" form:"parent_id" json:"parent_id"`       //父id
	ParentPath string    `gorm:"column:parent_path" form:"parent_path" json:"parent_path"` //父路径
	Sort       int       `gorm:"column:sort" form:"sort" json:"sort"`                      //子菜单排序
	Hidden     int       `gorm:"column:hidden" form:"hidden" json:"hidden"`                //是否展示
	Status     int       `gorm:"column:status" form:"status" json:"status"`
	CTime      time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime      time.Time `gorm:"column:mtime" form:"mtime" json:"name"`
	Action     Actions   `gorm:"column:action" form:"action" json:"action"`
	Source     Sources   `gorm:"column:source" form:"source" json:"source"`
	Did        string    `gorm:"column:did" form:"did" json:"did"` //冗余字段，为了查询方便
	Aid        string    `gorm:"column:aid" form:"aid" json:"aid"`
}

type Actions []*Action
type Sources []*Source

//Action 菜单动作对象 id和上面的id是同一个id
type Action struct {
	Id    int64     `gorm:"column:id" form:"id" json:"id"`
	Mid   string    `gorm:"column:mid" form:"mid" json:"mid"`    //权限id
	Code  string    `gorm:"column:code" form:"code" json:"code"` //动作编号
	Name  string    `gorm:"column:name" form:"name" json:"name"` //名称
	CTime time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}

//Source 菜单资源对象
type Source struct {
	Id     int64     `gorm:"column:id" form:"id" json:"id"`
	Mid    string    `gorm:"column:mid" form:"mid" json:"mid"`          //权限id
	Code   string    `gorm:"column:code" form:"code" json:"code"`       //资源编号
	Name   string    `gorm:"column:name" form:"name" json:"name"`       //名称
	Method string    `gorm:"column:method" form:"method" json:"method"` //请求方式
	Uri    string    `gorm:"column:uri" form:"uri" json:"uri"`          //请求uri
	CTime  time.Time `gorm:"column:ctime" form:"ctime" json:"ctime"`
	MTime  time.Time `gorm:"column:mtime" form:"mtime" json:"mtime"`
}

/**
结构体对象转换
*/

//func (m Menu)ToMenu()*Menu  {
//	item := &Menu{
//		Id:m.Id,
//	}
//}
