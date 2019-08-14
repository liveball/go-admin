package permission

import "time"

/**
user role
*/

type UserRole struct {
	ID     int64
	Rid    string    `json:"rid"  `
	Uid    string    `json:"uid"  `
	Status int       `json:"status" `
	CTime  time.Time `json:"ctime"  `
	MTime  time.Time `json:"mtime"  `
}
