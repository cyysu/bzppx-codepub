package controllers

import (
	"encoding/json"
	"strings"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type JsonResponse struct {
	Code     int                    `json:"code"`
	Message  interface{}            `json:"message"`
	Data     interface{}            `json:"data"`
	Redirect map[string]interface{} `json:"redirect"`
}

// prepare
func (this *BaseController) Prepare() {
	controllerName, _ := this.GetControllerAndAction()
	controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
}

// view layout title
func (this *BaseController) viewLayoutTitle(title, viewName, layout string) {
	this.Layout = "layout/" + layout + ".html"
	this.TplName = viewName + ".html"
	this.Data["title"] = title
	this.Render()
}

// view layout
func (this *BaseController) viewLayout(viewName, layout string) {
	this.Layout = "layout/" + layout + ".html"
	this.TplName = viewName + ".html"
	this.Data["title"] = ""
	this.Render()
}

// view
func (this *BaseController) view(viewName string) {
	this.Layout = "layout/default.html"
	this.TplName = viewName + ".html"
	this.Data["title"] = ""
	this.Render()
}

// error view
func (this *BaseController) viewError(errorMessage string, data ...interface{}) {
	this.Layout = "layout/install.html"
	redirect := "/"
	sleep := 2000
	if len(data) > 0 {
		redirect = data[0].(string)
	}
	if len(data) > 1 {
		sleep = data[1].(int)
	}
	_, actionName := this.GetControllerAndAction()
	methodName := strings.ToLower(actionName)
	this.TplName = "install/error.html"
	this.Data["title"] = "error"
	this.Data["method"] = methodName
	this.Data["message"] = errorMessage
	this.Data["redirect"] = redirect
	this.Data["sleep"] = sleep
	this.Render()
}

// view title
func (this *BaseController) viewTitle(title, viewName string) {
	this.Layout = "layout/default.html"
	this.TplName = viewName + ".html"
	this.Data["title"] = title
	this.Render()
}

// return json success
func (this *BaseController) jsonSuccess(message interface{}, data ...interface{}) {
	url := ""
	sleep := 300
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}
	if len(data) > 1 {
		url = data[1].(string)
	}
	if len(data) > 2 {
		sleep = data[2].(int)
	}
	this.Data["json"] = JsonResponse{
		Code:    1,
		Message: message,
		Data:    _data,
		Redirect: map[string]interface{}{
			"url":   url,
			"sleep": sleep,
		},
	}

	j, err := json.MarshalIndent(this.Data["json"], "", "\t")
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Abort(string(j))
	}
}

// return json error
func (this *BaseController) jsonError(message interface{}, data ...interface{}) {
	url := ""
	sleep := 2000
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}
	if len(data) > 1 {
		url = data[1].(string)
	}
	if len(data) > 2 {
		sleep = data[2].(int)
	}
	this.Data["json"] = JsonResponse{
		Code:    0,
		Message: message,
		Data:    _data,
		Redirect: map[string]interface{}{
			"url":   url,
			"sleep": sleep,
		},
	}
	j, err := json.MarshalIndent(this.Data["json"], "", " \t")
	if err != nil {
		this.Abort(err.Error())
	} else {
		this.Abort(string(j))
	}
}

// get client ip
func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// is post
func (this *BaseController) isPost() bool {
	return this.Ctx.Input.IsPost()
}