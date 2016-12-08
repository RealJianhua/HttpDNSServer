package edns

import (
	"fmt"

	"github.com/miekg/dns"
)

var DEFAULT_RESOLV_FILE = "/etc/resolv.conf"
var OPEN_DNS_SERVER = "8.8.8.8:53"
var cf *dns.ClientConfig

func Init() {
	var el error
	cf, el = dns.ClientConfigFromFile(DEFAULT_RESOLV_FILE)
	if el != nil {
		fmt.Println("DEFAULT_RESOLV_FILE ERR")
	}
}

func Find(domain string, ip string, ns string) (model EDNSModel, err error) {

	// 判断域名是否是标准domain
	domain = dns.Fqdn(domain)

	// 存储对象
	ednsModel := EDNSModel{domain, ip, nil, nil, nil, nil}

	if len(ns) > 0 {
		ednsModel.NS = []string{ns}[:]
	}

	fmt.Println("server dns servers:", cf.Servers)

	// 查询SOA记录
	err = FindSoaNs(&ednsModel)
	if err != nil {
		return ednsModel, err
	}

	fmt.Println("ednsModel:", ednsModel)

	// 查询A记录
	err = FindA(&ednsModel)
	if err != nil {
		return ednsModel, err
	}

	return ednsModel, nil
}

type DomainA struct {
	A   string `json:"ip"`
	Ttl int    `json:"ttl"`
}

type EDNSModel struct {
	Domain   string    `json:"domain"`
	ClientIP string    `json:"clientip"`
	A        []DomainA `json:"a"`
	CName    []string  `json:"cname,omitempty"`
	SOA      []string  `json:"soa,omitempty"`
	NS       []string  `json:"ns,omitempty"`
}

func (e *EDNSModel) String() string {
	str := ""
	str += fmt.Sprint("Domain :", e.Domain, "\n")
	str += fmt.Sprint("ClientIP :", e.ClientIP, "\n")
	str += fmt.Sprint("A : ", e.A, "\n")
	str += fmt.Sprint("CName :", e.CName, "\n")
	str += fmt.Sprint("SOA : ", e.SOA, "\n")
	str += fmt.Sprint("NS :", e.NS, "\n\n")
	return str
}
