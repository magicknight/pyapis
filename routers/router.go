package routers

import (
	"pyapis/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.IndexController{}, "GET:Index")

	beego.Router("/poxiao/index", &controllers.PoxiaoController{}, "GET:Index")
	beego.Router("/poxiao/detail", &controllers.PoxiaoController{}, "GET:Detail")
	beego.Router("/poxiao/movie", &controllers.PoxiaoController{}, "GET:Movie")

	beego.Router("/piaohua/index", &controllers.PiaohuaController{}, "GET:Index")
	beego.Router("/piaohua/detail", &controllers.PiaohuaController{}, "GET:Detail")
	beego.Router("/piaohua/movie", &controllers.PiaohuaController{}, "GET:Movie")
	beego.Router("/piaohua/lastest", &controllers.PiaohuaController{}, "GET:Lastest")

}
