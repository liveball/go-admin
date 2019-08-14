package permission

import "time"

/**
域
*/
type Domain struct {
	ID     int64
	Name   string    `json:"name"`
	Status int       `json:"status" `
	CTime  time.Time `json:"ctime"`
	MTime  time.Time `json:"mtime"  `
}
