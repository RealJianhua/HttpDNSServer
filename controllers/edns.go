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

func (c *EDNSController) Get() {
	edns.Init()
	clientIP := c.Input().get("ip")
	fmt.Println("clientip:", clientIP)
	ednsModel := edns.Find(c.GetString("domain"), c.GetString("ip"))
	fmt.Println(ednsModel)
	c.Data["json"] = &ednsModel
	c.ServeJSON()
}
