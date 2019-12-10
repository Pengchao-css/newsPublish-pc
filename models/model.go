package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
)
/*
func init(){
	conn,err := sql.Open("mysql","root:123456@tcp(127.0.0.1)/TwoBeeGo?charset=utf8")
	if err != nil{
		logs.Info("DB_wrong",err)
		logs.Error("DB_wrong",err)
		return
	}
	defer conn.Close()
	//conn.Exec("create table user(name varchar(40),password varchar(40))")
	//conn.Exec("insert into user(name,password) values(?,?)","chuanzhi","heima")
	res,err := conn.Query("select name from user")
	var name string
	for res.Next(){
		res.Scan(&name)
		logs.Info(name)
	}
}
 */

type User struct {
	Id int
	Name string
	PassWorld string
	Articles []*Article `orm:"reverse(many)"`
}

type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(20)"`
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string `orm:"size(500)"`
	Aimg string `orm:"size(100)"`

	ArticleType *ArticleType `orm:"rel(fk)"`
	Users []*User `orm:"rel(m2m)"`
}
type ArticleType struct {
	Id int
	TypeName string `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"`
}

func init (){
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1)/TwoBeeGo_01?charset=utf8")
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	orm.RunSyncdb("default",false ,true)
}
