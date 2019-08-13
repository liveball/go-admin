package user

import (
	"database/sql"
	"time"

	"github.com/yangjian/api"
	pb "github.com/yangjian/api/grpc/proto/user"
	"github.com/yangjian/lib/database/mysql"
	"github.com/yangjian/lib/grpc"
	"github.com/yangjian/mainsite/conf"
)

//Dao for user dao
type Dao struct {
	db  *sql.DB
	rpc pb.UserServiceClient
}

// New init dao
func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		db: mysql.NewMySQL(c.DB.Yangjian),
	}
	cli, err := api.NewClient(&grpc.ClientConfig{
		Dial:    time.Second * 50,
		Block:   true,
		Service: "grpc.srv.user",
		Target:  "127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379",
	})
	if err != nil {
		panic(err)
	}
	d.rpc = cli
	return
}

// Ping db
func (d *Dao) Ping() (err error) {
	return d.db.Ping()
}

// Close db
func (d *Dao) Close() (err error) {
	return d.db.Close()
}
