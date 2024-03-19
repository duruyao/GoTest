package args

import (
	"flag"
	"github.com/duruyao/gotest/conf"
)

func init() {
	flag.StringVar(&a.host, "host", conf.DefaultHost, "Host address to listen")
	flag.BoolVar(&a.wantHelp, "h", false, "Display this help message")
	flag.BoolVar(&a.wantHelp, "help", false, "Display this help message")
	flag.BoolVar(&a.wantVersion, "v", false, "Print version information and quit")
	flag.BoolVar(&a.wantVersion, "version", false, "Print version information and quit")
}
