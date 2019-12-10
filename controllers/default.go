package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["data"] = "CD SJ"
	c.TplName = "test.html"
}

func (c*MainController) ShowGet() {
	/*
		o := orm.NewOrm()
		var user models.User
		user.Name = "chuanzhi"
		user.PassWorld = "1234"
		count,err  := o.Insert(&user)
		if err != nil{
			logs.Info("wwrong",err)
		}
		logs.Info(count)
	*/
	/*
		o := orm.NewOrm()
		var UpDb models.User
		UpDb.Id = 1
		err := o.Read(&UpDb)
		if err != nil {
			logs.Error("no have data",err)
		}
		UpDb.Name = "Updata1"
		count,err := o.Update(&UpDb)
		if err != nil {
			logs.Error("wwrong",err)
		}
		logs.Info(count)
		logs.Error(UpDb)
	*/
	/*
	o := orm.NewOrm()
	var DelDb models.User
	DelDb.Id = 4
	count, err := o.Delete(&DelDb)
	if err != nil {
		logs.Error("delete is wrong", err)
	}
	logs.Info(count)
	logs.Error(DelDb)
	 */
	c.Data["data"] = "get-get"
	c.TplName = "test.html"
}
func (c*MainController) Post(){
	/*
	o := orm.NewOrm()
	var SeDb models.User
	SeDb.Name = "chuanzhi"
	err := o.Read(&SeDb,"Name")
	if err != nil{
		logs.Error("wwrong",err)
	}
	logs.Info(SeDb)
	 */
	c.Data["data"] = "POST fangfa"
	c.TplName = "post.html"
}

