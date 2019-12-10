package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"th_beego/models"
	"time"
)

type ArticleController struct {
	beego.Controller
}
func(this*ArticleController)ShowArticleList(){
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	var articles []models.Article
	//_,err := qs.All(&articles)
	//if err != nil {
	//	logs.Error("select DB error")
	//}
	count,_ := qs.Count()
	pageSize := 2
	pageCount := math.Ceil(float64(count)/float64(pageSize))
	pageIndex,err:= this.GetInt("pageIndex")
	if err != nil{
		pageIndex = 1
	}
	start := (pageIndex - 1)*pageSize
	qs.Limit(pageSize,start).RelatedSel("ArticleType").All(&articles)

	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	this.Data["types"] = types
	typeName := this.GetString("select")
	logs.Info(typeName)
	qs.Limit(pageSize,start).RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).All(&articles)

	this.Data["pageIndex"] = pageIndex
	this.Data["pageCount"] = int(pageCount)
	this.Data["count"] = count
	this.Data["articles"] = articles
	this.TplName = "index.html"
}
func(this*ArticleController)ShowAddArticle(){
	 o := orm.NewOrm()
	 var types []models.ArticleType
	 o.QueryTable("ArticleType").All(&types)

	 this.Data["types"] = types
	this.TplName = "add.html"
}
func(this*ArticleController)HandleAddArticle(){
	articleName := this.GetString("articleName")
	content := this.GetString("content")
	if articleName == "" || content == ""{
		this.Data["errmsg"] = "input data is wrong"
		this.TplName = "add.html"
		return
	}
	//logs.Info(articleName,content)
	file,head,err := this.GetFile("uploadname")
	defer file.Close()
	if err !=  nil{
		this.Data["errmsg"] = "file upload is error"
		this.TplName = "add.html"
		return
	}
	if head.Size > 5000000{
		this.Data["errmsg"] = "file is big"
		this.TplName = "add.html"
		return
	}
	ext := path.Ext(head.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg"{
		this.Data["errmsg"] = "file format error"
		this.TplName = "add.html"
		return
	}
	fileName := time.Now().Format("2006-01-02-15:04:05") + ext
	this.SaveToFile("uploadname","./static/img/"+fileName)

	o := orm.NewOrm()
	var article models.Article
	article.ArtiName = articleName
	article.Acontent = content
	article.Aimg = "/static/img"+fileName
	typeName := this.GetString("select")
	var articleType models.ArticleType
	articleType.TypeName = typeName
	o.Read(&articleType,"TypeName")
	article.ArticleType = &articleType

	o.Insert(&article)

	this.Redirect("/showArticleList",302)
}
func(this*ArticleController)ShowArticleDetail(){
	id,err := this.GetInt("articleId")
	if err!= nil {
		logs.Error("link error")
	}
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	o.Read(&article)

	logs.Error(article.Acount)

	article.Acount += 1
	o.Update(&article)

	this.Data["article"] = article
	this.TplName = "content.html"

}

func(this*ArticleController)ShowUpdateArticle(){
	//huoqushuji
	id,err := this.GetInt("articleId")
	if err != nil {
		logs.Error("qingqiu wenzhang error")
		return
	}
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	o.Read(&article)
	this.Data["article"] = article
	this.TplName = "update.html"
}

func UploadFile(this*beego.Controller,filePath string)string{
	file,head,err := this.GetFile(filePath)
	if head.Filename == ""{
		return "NoImg"
	}
	if err !=  nil{
		this.Data["errmsg"] = "file upload is error"
		this.TplName = "add.html"
		return ""
	}
	defer file.Close()

	if head.Size > 5000000{
		this.Data["errmsg"] = "file is big"
		this.TplName = "add.html"
		return ""
	}
	ext := path.Ext(head.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg"{
		this.Data["errmsg"] = "file format error"
		this.TplName = "add.html"
		return ""
	}
	fileName := time.Now().Format("2006-01-02-15:04:05") + ext
	this.SaveToFile(filePath,"./static/img/"+fileName)
	return "/static/img/"+fileName
}

func(this*ArticleController)HandleUpdateArticle(){

	id,err := this.GetInt("articleId")
	articleName := this.GetString("articleName")
	content := this.GetString("content")
	filePath := UploadFile(&this.Controller,"uploadname")
	if err != nil || articleName == "" || content == "" || filePath == ""{
		logs.Error("way error")
		return
	}
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	err = o.Read(&article)
	if err != nil{
		logs.Error("nohave")
		return
	}
	article.ArtiName = articleName
	article.Acontent = content
	if filePath != "NoImg"{
		article.Aimg = filePath
	}
	article.Aimg = filePath
	o.Update(&article)
	this.Redirect("/showArticleList",302)
}

func(this*ArticleController)DeleteArticle(){
	id,err := this.GetInt("articleId")
	if err != nil {
		logs.Info("erroe")
		return
	}
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	o.Delete(&article)
	this.Redirect("/showArticleList",302)

}

func(this*ArticleController)ShowAddType(){
	o := orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)

	this.Data["types"] = types
	this.TplName = "addType.html"
}

func(this*ArticleController)HandleAddType(){
	typeName := this.GetString("typeName")
	if typeName == ""{
		logs.Info("massige is not full")
		return
	}
	o := orm.NewOrm()
	var articleType models.ArticleType
	articleType.TypeName = typeName
	o.Insert(&articleType)

	this.Redirect("/addType",302)
}