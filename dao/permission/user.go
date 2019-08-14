package permission

import (
	"log"
	"time"

	"go-admin/model/permission"
)

//AddUser add user.
func (d *Dao) AddUser(user *permission.User) (err error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	user.CTime = now
	user.MTime = now

	if err = d.db.Create(user).Error; err != nil {
		log.Printf("Add user error(%v)", err)
	}

	return
}
