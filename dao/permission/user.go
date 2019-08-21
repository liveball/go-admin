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

//delete user
func (d *Dao) DeleteUser(id int64) (err error) {
	if err = d.db.Where(" id = ? ", id).Delete(permission.User{}).Error; err != nil {
		log.Printf("delete user error(%v)", err)
	}
	return
}

//update user

func (d *Dao) UpdateUser(user *permission.User) (err error) {
	u := permission.User{}
	if err = d.db.Model(&u).Updates(user).Error; err != nil {
		log.Printf("update user error(%v)", err)
	}
	return
}

//find user
/**
查询user  先保证是查询这个组下面
如果aid不为空就查询这个组下面这个app的用户(这个app的管理员)
如果aid为空就是查询这个组下面的所有app的用户(这个组的管理员)
*/
func (d *Dao) FindUser(did, aid int) {

}
