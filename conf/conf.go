package conf

import (
	"flag"

	"github.com/BurntSushi/toml"
	"go-admin/lib"
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
	// orm
	ORM *lib.Config
}

// HTTP conf.
type HTTP struct {
	Addr string
}

func init() {
	flag.StringVar(&ConfPath, "conf", "/media/xiaohan/mk/code/go_code/src/github.com/yangjian/mainsite/cmd/mainsite.toml", "default config path")
}

// Init conf.
func Init() (err error) {
	_, err = toml.DecodeFile(ConfPath, &Conf)
	return
}
