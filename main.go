package main

import (
	_ "wechat/routers"
	"github.com/astaxie/beego"
	"wechat/proc"
)

func main() {
	proc.Menu()
	beego.Run()

}

