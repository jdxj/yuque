package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/jdxj/yuque/play"
)

func init() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"yuque.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func main() {
	c, err := play.NewCounter()
	if err != nil {
		logs.Error("create counter error: %s", err)
		return
	}

	c.Users()
	c.Stop()
}
