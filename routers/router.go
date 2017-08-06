package routers

import (
	"wechat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/we_connect",&controllers.WeChatController{})
	beego.Router("/message",&controllers.MessageController{})
}
