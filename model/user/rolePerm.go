package adminm

import "time"

/**
role  perms
*/

type RolePerms struct {
	ID     int64
	Rid    string    `json:"rid"  `
	Pid    string    `json:"pid"  `
	Status int       `json:"status"  `
	CTime  time.Time `json:"ctime"  `
	MTime  time.Time `json:"mtime"   `
}
