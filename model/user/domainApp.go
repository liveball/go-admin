package adminm

import "time"

/**
Id 用Did和Aid计算出来一个值，然后代表当前唯一id,这样方便查询在插入数据的时候查询是否存在

*/
type DomainApp struct {
	ID     int64
	Name   string    `json:"name"`
	Did    string    `json:"did" `
	Status int       `json:"status" `
	CTime  time.Time `json:"ctime" `
	MTime  time.Time `json:"mtime" `
}
