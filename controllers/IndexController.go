package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

// index page
func (c *IndexController) Index(){
	c.TplName = "index.tpl"
}