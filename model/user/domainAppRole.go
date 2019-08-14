package adminm

import "time"

/**
domain app role关联关系(admin用)
*/

type DomainAppRole struct {
	ID     int64     `json:"id"`
	Did    string    `json:"did"`
	Aid    string    `json:"aid" `
	Rid    string    `json:"rid"`
	Status int       `json:"status" `
	CTime  time.Time `json:"ctime" `
	MTime  time.Time `json:"mtime" `
}
