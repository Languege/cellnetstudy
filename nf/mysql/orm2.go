package main

import (
	//"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver"
	"time"
)

type Userinfo struct {
	Uid        int `orm:"PK"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username   string
	Departname string
	Created    time.Time
}

type User struct {
	Uid     int `orm:"PK"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8", 30)

	// 需要在init中注册定义的model
	orm.RegisterModel(new(Userinfo), new(User), new(Profile), new(Tag), new(Post))
}

func main() {
	orm.Debug = true
	o := orm.NewOrm()
	var user Userinfo
	user.Username = "zxxx"
	user.Departname = "zxxx"
	user.Created = time.Now()

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}
}
