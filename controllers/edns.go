package controllers

import (
	"fmt"

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
	edns.Init()
	clientIP := this.Ctx.Input.IP
	fmt.Println("clientip:", clientIP)
	ednsModel := edns.Find(this.GetString("domain"), this.GetString("ip"))
	fmt.Println(ednsModel)
	this.Data["json"] = &ednsModel
	this.ServeJSON()
}
