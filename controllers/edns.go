package controllers

import (
	"fmt"
	"strings"

	"../dns/edns"
	"github.com/astaxie/beego"
)

type EDNSController struct {
	beego.Controller
}

type Result struct {
	EdnsResult interface{} `json:"data,omitempty"`
	Error      int         `json:"ec"`
	ErrMsg     string      `json:"em"`
}

func (this *EDNSController) Get() {
	clientIP := strings.Split(this.Ctx.Request.RemoteAddr, ":")[0]
	fmt.Println("clientip:", clientIP)

	result := Result{nil, 0, ""}

	edns.Init()
	ednsModel, err := edns.Find(this.GetString("domain"), clientIP)
	if err != nil {
		result.Error = 500
		result.ErrMsg = err.Error()
	} else {
		result.EdnsResult = &ednsModel
		fmt.Println(ednsModel)
	}

	this.Data["json"] = &result
	this.ServeJSON()
}
