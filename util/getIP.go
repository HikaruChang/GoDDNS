package util

import (
	"io/ioutil"
	"net/http"
)

type URL struct {
	IPV4API string
	IPV6API string
}

var url *URL = &URL{
	IPV4API: "https://api-ipv4.ip.sb/ip",
	IPV6API: "https://api-ipv6.ip.sb/ip",
}

func IPv4() string {
	resp, err := http.Get(url.IPV4API)
	CheckErr(err)
	respBody, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	return string(respBody)
}

func IPv6() string {
	resp, err := http.Get(url.IPV6API)
	CheckErr(err)
	respBody, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	return string(respBody)
}
