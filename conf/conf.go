package conf

import (
	"flag"
	"log"

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
	flag.StringVar(&ConfPath, "conf", "./cmd/mainsite.toml", "default config path")
}

// Init conf.
func Init() (err error) {
	_, err = toml.DecodeFile(ConfPath, &Conf)
	if err!=nil{
		log.Fatalf("Init toml.DecodeFile error(%v)", err)
	}
	return
}
