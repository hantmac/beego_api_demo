package controllers
import (
	"fmt"
	"github.com/astaxie/beego"
)
import (
	"encoding/json"
	"firstAPI/models"
)

type StudentController struct {
beego.Controller
}
// @Title 获得所有学生
// @Description 返回所有的学生数据
// @Success 200 {object} models.Student
// @router / [get]
func (u *StudentController) GetAll() {
	ss := models.GetAllStudents()
	u.Data["json"] = ss
	fmt.Println(ss)
	u.ServeJSON()
}
// @Title 获得一个学生
// @Description 返回某学生数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Student
// @router /:id [get]
func (u *StudentController) GetById() {
	id ,_:= u.GetInt(":id")
	fmt.Println(id)
	s := models.GetStudentById(id)
	u.Data["json"] = s
	u.ServeJSON()
}
// @Title 创建用户
// @Description 创建用户的描述
// @Param      body          body   models.Student true          "body for user content"
// @Success 200 {struct} models.Student
// @Failure 403 body is empty
// @router / [post]
func (u *StudentController) Post() {
	var s models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	uid := models.AddStudent(&s)
	u.Data["json"] = uid
	fmt.Println(uid)
	u.ServeJSON()
}
// @Title 修改用户
// @Description 修改用户的内容
// @Param      body          body   models.Student true          "body for user content"
// @Success 200 {int} models.Student
// @Failure 403 body is empty
// @router / [put]
func (u *StudentController) Update() {
	var s models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	models.UpdateStudent(&s)
	u.Data["json"] = s
	u.ServeJSON()
}
// @Title 删除一个学生
// @Description 删除某学生数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Student
// @router /:id [delete]
func (u *StudentController) Delete() {
	id ,_:= u.GetInt(":id")
	models.DeleteStudent(id)
	u.Data["json"] = true
	u.ServeJSON()
}

