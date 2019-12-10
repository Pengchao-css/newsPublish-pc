package routers

import (
	"th_beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:ShowGet;post:Post")
	beego.Router("/register",&controllers.UserController{},"get:ShowRegister;post:HandlePost")
	beego.Router("/login",&controllers.UserController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/showArticleList",&controllers.ArticleController{},"get:ShowArticleList")
	beego.Router("/addArticle",&controllers.ArticleController{},"get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/showArticleDetail",&controllers.ArticleController{},"get:ShowArticleDetail")
	beego.Router("/updateArticle",&controllers.ArticleController{},"get:ShowUpdateArticle;post:HandleUpdateArticle")
	beego.Router("/deleteArticle",&controllers.ArticleController{},"get:DeleteArticle")
	beego.Router("addType",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddType")


}

