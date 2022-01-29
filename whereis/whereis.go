package whereis

import (
	"fmt"
	"net"
	"strings"

	"github.com/likexian/whois"
)

type NetworkAdomin struct {
	IpRange string
	NetName string
	Country string
	Admin   string
}

type Summary struct {
	TargetDomain  string
	TargetIp      string
	WhoisResponse string
	ParseResult   []NetworkAdomin
}

func Resolve(domain string, verbose bool) error {
	var summary Summary
	summary.TargetDomain = domain
	addr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		return err
	}
	summary.TargetIp = addr.String()
	summary.WhoisResponse, err = whois.Whois(summary.TargetIp, "whois.iana.org")
	if err != nil {
		return err
	}
	summary.ParseWhoisResponse()
	summary.EchoResult(verbose)
	return nil
}

func (s *Summary) ParseWhoisResponse() error {
	paragraph := strings.Split(s.WhoisResponse, "\n\n")
	for _, v := range paragraph {
		tmp := NetworkAdomin{}
		row := strings.Split(v, "\n")
		for _, val := range row {
			col := strings.Split(val, ":")
			switch col[0] {
			case "inetnum", "NetRange":
				tmp.IpRange = strings.TrimSpace(col[1])
			case "netname", "NetName":
				tmp.NetName = strings.TrimSpace(col[1])
			case "country":
				tmp.Country = strings.TrimSpace(col[1])
			case "descr", "Organization", "organisation":
				if tmp.Admin == "" {
					tmp.Admin = strings.TrimSpace(col[1])
				}
			}
		}
		if tmp.IpRange != "" {
			s.ParseResult = append(s.ParseResult, tmp)
		}
	}
	return nil
}

func (s *Summary) EchoResult(verbose bool) {
	fmt.Printf("Target domain:%s\n", s.TargetDomain)
	fmt.Printf("Target ip    :%s\n\n", s.TargetIp)
	fmt.Printf("Network Admin:%s\n", s.ParseResult[len(s.ParseResult)-1].Admin)
	fmt.Printf("Network name :%s\n", s.ParseResult[len(s.ParseResult)-1].NetName)
	fmt.Printf("ip range     :%s\n", s.ParseResult[len(s.ParseResult)-1].IpRange)
	fmt.Printf("country      :%s\n", s.ParseResult[len(s.ParseResult)-1].Country)
	if verbose {
		fmt.Printf("\n=========Network Details=========\n")
		for i, v := range s.ParseResult {
			fmt.Printf("%d:\n", i)
			fmt.Printf(" Network Admin:%s\n", v.Admin)
			fmt.Printf(" Network name :%s\n", v.NetName)
			fmt.Printf(" ip range     :%s\n", v.IpRange)
			fmt.Printf(" country      :%s\n", v.Country)
		}
	}
}
