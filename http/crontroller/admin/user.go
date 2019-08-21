package admin

import (
	"comadmin/model/admin"
	"comadmin/pkg/app"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
	"net/http"
)

//用户注册(对外暴露),不需要带token等信息
func (h HttpAdminHandler) Register(c app.GContext) {
	type P struct {
		Name string `json:"name"  binding:"required"`
		Pwd  string `json:"pwd"  binding:"required" `
		Did  string `json:"did"  binding:"required" ` //属于哪个域
		Aid  string `json:"aid"  binding:"required"`  //属于哪个app
	}

	var p P
	code := e.Success
	g := app.G{c}

	if !utils.CheckError(c.ShouldBindJSON(&p), "createUser") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	//注册的时候校验did aid是否合法
	domainApp := &admin.DomainApp{Did: p.Did, Id: p.Aid}
	code = h.logic.FindById(domainApp)
	if code != e.Success {
		code = e.ParamError
		g.Json(http.StatusOK, code, "分配id异常")
		return
	}
	user := admin.User{Name: p.Name, Id: utils.EncodeMd5(utils.StringJoin(p.Name, p.Did, p.Aid)), Did: p.Did, Aid: p.Aid, Pwd: utils.EncodeMd5(p.Pwd), Status: 1}
	appUser := admin.DomainAppUser{Id: utils.EncodeMd5(utils.StringJoin(p.Name, p.Did, p.Aid)), Did: p.Did, Aid: p.Aid, Uid: utils.EncodeMd5(utils.StringJoin(p.Name, p.Did, p.Aid)), Status: 1}

	//todo 是否考虑开启一个goroutine去处理
	code = h.logic.Create(appUser)
	code = h.logic.Create(user)
	g.Json(http.StatusOK, code, "")
	return
}

//登录的时候需要判断did和app id是否存在
func (h HttpAdminHandler) Login(c app.GContext) {
	type P struct {
		Name string `json:"name"  binding:"required"`
		Pwd  string `json:"pwd"  binding:"required" `
		Did  string `json:"did"  binding:"required" ` //属于哪个域
		Aid  string `json:"aid"  binding:"required"`  //属于哪个app
	}

	var p P
	code := e.Success
	g := app.G{c}

	if !utils.CheckError(c.ShouldBindJSON(&p), "createUser") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	user := admin.User{Name: p.Name, Id: utils.EncodeMd5(utils.StringJoin(p.Name, p.Did, p.Aid)), Did: p.Did, Aid: p.Aid, Pwd: utils.EncodeMd5(p.Pwd)}
	code = h.logic.Login(user)
	g.Json(http.StatusOK, code, "")
	return
}

//admin手动添加用户
func (h HttpAdminHandler) RegisterAdminUser(c app.GContext) {
	type P struct {
		Name string `json:"name"  binding:"required"`
		Pwd  string `json:"pwd"  binding:"required" `
		Did  string `json:"did"  binding:"required" ` //属于哪个域
		Aid  string `json:"aid"  binding:"required"`  //属于哪个app
	}

	var p P
	code := e.Success
	g := app.G{c}

	if !utils.CheckError(c.ShouldBindJSON(&p), "createUser") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	domainApp := admin.DomainApp{Did: p.Did, Id: p.Aid}
	code = h.logic.FindById(domainApp)
	if code != e.Success {
		code = e.ParamError
		g.Json(http.StatusOK, code, "分配id异常")
	}
	user := admin.User{Name: p.Name, Id: utils.EncodeMd5(utils.StringJoin(p.Name, p.Did, p.Aid)), Did: p.Did, Aid: p.Aid, Pwd: utils.EncodeMd5(p.Pwd)}
	code = h.logic.Create(user)
	g.Json(http.StatusOK, code, "")
	return
}

//删除用户
func (h HttpAdminHandler) DeleteUser(c app.GContext) {
	g := app.G{c}

	type P struct {
		Id string `json:"id"  binding:"required"`
	}

	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "deleteDomain") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	domain := admin.User{Id: p.Id}
	code = h.logic.Delete(domain)
	g.Json(http.StatusOK, code, "")
	return

}

//更新用户
//更新用户密码
//更新手机号码
//更新头像 分开写
//禁用用户登录

func (h HttpAdminHandler) UpdateUserPhone(c app.GContext) {
	g := app.G{c}
	type P struct {
		Id    string `json:"id" binding:"required"`
		Phone string `json:"name" binding:"required"`
	}

	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updatePhone") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	domain := admin.User{Id: p.Id, Phone: p.Phone}
	cols := []string{"phone"}

	code = h.logic.Update(domain, cols)
	g.Json(http.StatusOK, code, "")
	return
}

//更新密码
func (h HttpAdminHandler) UpdateUserPwd(c app.GContext) {
	g := app.G{c}
	type P struct {
		Id  string `json:"id" binding:"required"`
		Pwd string `json:"name" binding:"required"`
	}

	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updatePwd") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	domain := admin.User{Id: p.Id, Pwd: utils.EncodeMd5(p.Pwd)}
	cols := []string{"pwd"}

	code = h.logic.Update(domain, cols)
	g.Json(http.StatusOK, code, "")
	return
}

//禁用账户

func (h HttpAdminHandler) UpdateUserStatus(c app.GContext) {
	g := app.G{c}
	type P struct {
		Id     string `json:"id" binding:"required"`
		Status int    `json:"name" binding:"required"`
	}
	var p P
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "updateStatus") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	domain := admin.User{Id: p.Id, Status: p.Status}
	cols := []string{"status"}

	code = h.logic.Update(domain, cols)
	g.Json(http.StatusOK, code, "")
	return
}

/**
查询用户
*/

func (h HttpAdminHandler) FindUser(c app.GContext) {
	g := app.G{c}

	var p admin.UserParams
	code := e.Success
	if !utils.CheckError(c.ShouldBindJSON(&p), "findUser") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	domain := admin.User{Id: p.Id, Name: p.Name, Status: p.Status, Phone: p.Phone, Did: p.Did, Aid: p.Aid}

	list, count := h.logic.Find(domain, p.Pn, p.Ps)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, code, m)
	return
}
