package conf

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/yangjian/lib/database/mysql"
	"github.com/yangjian/lib/grpc"
	"github.com/yangjian/lib/kafka"
)

// Conf info.
var (
	ConfPath string
	Conf     = &Config{}
)

// Config struct.
type Config struct {
	// http
	HTTP *HTTP
	// db
	DB *DB
	//kafka
	Kafka        *kafka.Config
	ClientConfig *grpc.ClientConfig
}

// HTTP conf.
type HTTP struct {
	Addr string
}

// DB conf.
type DB struct {
	// yangjian
	Yangjian *mysql.Config
}

func init() {
	flag.StringVar(&ConfPath, "conf", "/media/xiaohan/mk/code/go_code/src/github.com/yangjian/mainsite/cmd/mainsite.toml", "default config path")
}

// Init conf.
func Init() (err error) {
	_, err = toml.DecodeFile(ConfPath, &Conf)
	return
}
