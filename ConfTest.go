package main

import (
	"fmt"

	"github.com/go-ini/ini"
)

func main() {
	cfg, _ := ini.Load("pkg/conf/config.ini")
	fmt.Println(cfg.Section("ssh").Key("key").MustString("hoge"))
}
