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

//type dataModel struct {
//	Domain    string    `json:"domain"`
//	Device_ip string    `json:"device_ip"`
//	Device_sp string    `json:"device_sp"`
//	DNS       EDNSModel `json:"dns"`
//}

func (this *EDNSController) Get() {
	clientIP := this.Ctx.Request.RemoteAddr
	end := strings.Index(clientIP, ":")
	if end < 0 {
		end = len(clientIP)
	}
	clientIP = string(clientIP[0:end])
	fmt.Println("clientip:", clientIP)

	edns.Init()
	ednsModel := edns.Find(this.GetString("domain"), clientIP)
	fmt.Println(ednsModel)
	this.Data["json"] = &ednsModel
	this.ServeJSON()
}
