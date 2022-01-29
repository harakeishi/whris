package whereis

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/likexian/whois"
)

type NetworkAdomin struct {
	IpRange  string
	NetName  string
	Country  string
	Descript string
}

func ResolveIP(domain string) {
	addr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		fmt.Println("Resolve error ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%+v\n", addr.String())
	result, err := whois.Whois(addr.String(), "whois.iana.org")
	if err != nil {
		fmt.Println("whois error ", err.Error())
		os.Exit(1)
	}
	splittedStr := strings.Split(result, "\n\n")
	for _, v := range splittedStr {
		tmp := NetworkAdomin{}
		t := strings.Split(v, "\n")
		for _, s := range t {
			y := strings.Split(s, ":")
			switch y[0] {
			case "inetnum", "NetRange":
				tmp.IpRange = strings.TrimSpace(y[1])
			case "netname", "NetName":
				tmp.NetName = strings.TrimSpace(y[1])
			case "country":
				tmp.Country = strings.TrimSpace(y[1])
			case "descr", "Organization", "organisation":
				if tmp.Descript == "" {
					tmp.Descript = strings.TrimSpace(y[1])
				}
			}
		}
		if tmp.IpRange != "" {
			fmt.Printf("%+v\n", tmp)
		}
	}

}
