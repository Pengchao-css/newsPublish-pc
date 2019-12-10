package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"th_beego/models"
)
type UserController struct {
	beego.Controller
}
func (this*UserController)ShowRegister(){
	this.TplName = "register.html"
}
func (this*UserController)HandlePost(){
	userName := this.GetString("userName")
	pwd := this.GetString("password")
	//logs.Info(userName,pwd)
	if userName == "" || pwd == ""{
		this.Data["errmsg"] = "Register is wrong"
		logs.Error("register is wrong")
		this.TplName = "register.html"
		return
	}
	o := orm.NewOrm()
	var register models.User
	register.Name = userName
	register.PassWorld = pwd
	o.Insert(&register)
	//this.Ctx.WriteString("Registing is seccessful")
	this.Redirect("/login",302)
	//this.TplName = ("/login.html")
}
func (this*UserController)ShowLogin(){
	initName := this.Ctx.GetCookie("userName")
	if initName == ""{
		this.Data["userName"] = ""
		this.Data["checked"] = ""
	}else{
		this.Data["userName"] = initName
		this.Data["checked"] = "checked"
	}
	this.TplName = "login.html"
}
func (this*UserController)HandleLogin(){
	initName := this.GetString("userName")
	initpwd := this.GetString("password")
	if initName == "" || initpwd == ""{
		this.Data["errmsg"] = "Login is wrong"
		this.TplName = "login.html"
		return
	}
	o := orm.NewOrm()
	var login models.User
	login.Name = initName
	err := o.Read(&login,"Name")
	if err != nil{
		this.Data["errmsg"] = "User is not have "
		this.TplName = "login.html"
		return
	}
	if login.PassWorld !=  initpwd {
		this.Data["errmsg"] = "password is wrong"
		this.TplName = "login.html"
		return
	}
	data := this.GetString("remember")
	logs.Info(data)
	if data == "on"{
		this.Ctx.SetCookie("userName",initName,100)
	}else {
		this.Ctx.SetCookie("userName",initName,-1)
	}
	//this.Ctx.SetCookie("userName",userName,100)
	this.Redirect("/showArticleList",302)
	//this.Ctx.WriteString("login is seccessful")
}


