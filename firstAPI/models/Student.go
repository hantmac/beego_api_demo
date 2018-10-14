package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)


//在models模块中创建一个struct，目的是使用beego的orm框架，使struct与数据库中的字段产生对应关系
type Student struct {
	Id int`orm:"column(Id)"` //column()括号中的字段就是在定义数据库时的相应字段，这一段必须严格填写，不然在API读写数据时就会出现读不到或者写不进去的问题
	Name string  `orm:"column(Name)"`
	BirthDate string `orm:"column(Birthdate)"`
	Gender bool `orm:"column(Gender)"`
	Score int `orm:"column(Score)"`
}


//该函数获得数据库中所有student的信息，返回值是一个结构体数组指针
func GetAllStudents() []*Student {
	o := orm.NewOrm() //产生一个orm对象
	o.Using("default") //这句话的意思是使用定义的默认数据库，与main.go中的orm.RegisterDataBase()对应
	var students []*Student //定义指向结构体数组的指针
	q := o.QueryTable("student")//获得一个数据库表的请求
	q.All(&students)//取到这个表中的所有数据

	return students

}


//该函数根据student中的Id，返回该学生的信息
func GetStudentById(id int) Student {
	u := Student{Id:id}//根据所传入的Id得到对应student的对象
	o := orm.NewOrm()//new 一个orm对象
	o.Using("default")//使用最开始定义的default数据库
	err := o.Read(&u)//读取Id=id的student的信息

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")//对应操作，不一定是print
	} else if err == orm.ErrMissPK {
		fmt.Println("没有主键")
	}

	return u
}


//添加一个学生的信息到数据库中，参数是指向student结构题的指针
func AddStudent(student *Student) Student {
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(student)//插入数据库

	return *student
}

func UpdateStudent(student *Student) {
	o := orm.NewOrm()
	o.Using("default")
	o.Update(student)//更新该student的信息
}

func DeleteStudent(id int) {
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&Student{Id:id})//删除对应id的student的信息
}

func init()  {
	orm.RegisterModel(new(Student))//将数据库注册到orm
}