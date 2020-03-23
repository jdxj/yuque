package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/astaxie/beego/logs"
	"github.com/jdxj/yuque/play"
)

func init() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"yuque_get_users.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	c, err := play.NewCounter()
	if err != nil {
		logs.Error("create counter error: %s", err)
		return
	}

	c.Users(390562)
	c.Remaining()
	logs.Debug("start counter")

	select {
	case <-ch:
		logs.Debug("receive stop signal")
	}

	if err := c.Stop(); err != nil {
		logs.Error("stop err: %s", err)
		return
	}
	logs.Debug("stop")
}
