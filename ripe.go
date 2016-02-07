package ripego

import (
	"strings"
)

const (
	RIPE_WHOIS_SERVER = "whois.ripe.net"
)

type Ripe struct {
}

func (r Ripe) Check(search string) (w WhoisInfo, err error) {
	content, err := GetTcpContent(search, RIPE_WHOIS_SERVER)

	if err != nil {
		return w, err
	}

	w = r.ParseData(content)

	return w, err
}

func (r Ripe) ParseData(whoisData string) WhoisInfo {

	wi := WhoisInfo{}
	wi.Inetnum = ParseRPSLValue(whoisData, "inetnum", "inetnum")
	wi.Netname = ParseRPSLValue(whoisData, "inetnum", "netname")
	wi.AdminC = ParseRPSLValue(whoisData, "inetnum", "admin-c")
	wi.Country = ParseRPSLValue(whoisData, "inetnum", "country")
	wi.Created = ParseRPSLValue(whoisData, "inetnum", "created")
	wi.Descr = ParseRPSLValue(whoisData, "inetnum", "descr")
	wi.LastModified = ParseRPSLValue(whoisData, "inetnum", "last-modified")
	wi.MntBy = ParseRPSLValue(whoisData, "inetnum", "mnt-by")
	wi.MntLower = ParseRPSLValue(whoisData, "inetnum", "mnt-lower")
	wi.MntRoutes = ParseRPSLValue(whoisData, "inetnum", "mnt-routes")
	wi.Source = ParseRPSLValue(whoisData, "inetnum", "source")
	wi.TechC = ParseRPSLValue(whoisData, "inetnum", "tech-c")
	wi.Organization = ParseRPSLValue(whoisData, "inetnum", "org")

	p := WhoisPerson{}
	p.Name = ParseRPSLValue(whoisData, "person", "person")
	p.AbuseMailbox = ParseRPSLValue(whoisData, "person", "abuse-mailbox")
	p.Address = ParseRPSLValue(whoisData, "person", "address")
	p.Created = ParseRPSLValue(whoisData, "person", "created")
	p.LastModified = ParseRPSLValue(whoisData, "person", "last-modified")
	p.MntBy = ParseRPSLValue(whoisData, "person", "mnt-by")
	p.NicHdl = ParseRPSLValue(whoisData, "person", "nic-hdl")
	p.Phone = ParseRPSLValue(whoisData, "person", "phone")
	p.Source = ParseRPSLValue(whoisData, "person", "source")

	rt := WhoisRoute{}
	rt.Origin = ParseRPSLValue(whoisData, "route", "origin")
	rt.Created = ParseRPSLValue(whoisData, "route", "created")
	rt.Descr = ParseRPSLValue(whoisData, "route", "descr")
	rt.LastModified = ParseRPSLValue(whoisData, "route", "last-modified")
	rt.Route = ParseRPSLValue(whoisData, "route", "route")
	rt.Source = ParseRPSLValue(whoisData, "route", "source")

	wi.Person = p
	wi.Route = rt

	return wi
}

func (r Ripe) hasIP(ipaddr string) bool {

	ips := []string{"141", "51"}

	for i := range ips {
		if strings.HasSuffix(ipaddr, ips[i]) {
			return true
		}
	}

	return false
}