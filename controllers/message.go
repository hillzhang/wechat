package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type MessageController struct {
	beego.Controller
}

func(c *MessageController) Post(){
	data,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	d,err := json.Marshal(string(data))
	if err != nil {
		fmt.Println(err)
	}
	c.Data["json"] = string(d)
	c.ServeJSON(true)
}
