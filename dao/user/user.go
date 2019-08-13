package user

import (
	pb "github.com/yangjian/api/grpc/proto/user"
	"github.com/yangjian/lib/log"
)

//Users get all the user.
func (d *Dao) Users() (res []*pb.User, err error) {
	rows, err := d.db.Query("SELECT id, username FROM user")
	defer rows.Close()
	if err != nil {
		log.Errorf("d.db.Query error(%v)", err)
		return
	}

	res = make([]*pb.User, 0)
	for rows.Next() {
		u := &pb.User{}
		if err = rows.Scan(&u.Id, &u.UserName); err != nil {
			return
		}
		res = append(res, u)
	}
	return
}
