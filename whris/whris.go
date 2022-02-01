package whris

import (
	"fmt"
	"net"
	"strings"

	"github.com/likexian/whois"
)

type NetworkAdmin struct {
	IpRange string
	NetName string
	Country string
	Admin   string
}

type Summary struct {
	TargetDomain        string
	TargetIp            string
	WhoisResponseServer string
	WhoisResponse       string
	ParseResult         []NetworkAdmin
}

func Resolve(domain string, verbose bool) (Summary, error) {
	var summary Summary
	summary.WhoisResponseServer = "whois.apnic.net"
	summary.TargetDomain = domain
	addr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		return Summary{}, err
	}
	summary.TargetIp = addr.String()
	summary.WhoisResponse, err = whois.Whois(summary.TargetIp, "whois.iana.org")
	if err != nil {
		return Summary{}, err
	}
	summary.SetWhoisResponseServerFromWhoisResponse()
	summary.ParseWhoisResponse()
	if !summary.ParseCheck() {
		summary.ParseResult = summary.ParseResult[1:]
		summary.WhoisResponse, err = whois.Whois(summary.TargetIp, summary.WhoisResponseServer)
		if err != nil {
			return Summary{}, err
		}
		summary.ParseWhoisResponse()
	}
	summary.EchoResult(verbose)
	return summary, nil
}

func (s *Summary) ParseWhoisResponse() {
	paragraph := s.BreakDownWhoisResponseIntoParagraphs()
	for _, v := range paragraph {
		tmp := NetworkAdmin{}
		row := strings.Split(v, "\n")
		for _, val := range row {
			col := strings.Split(val, ":")
			switch col[0] {
			case "inetnum", "NetRange":
				tmp.IpRange = strings.TrimSpace(col[1])
			case "netname", "NetName":
				tmp.NetName = strings.TrimSpace(col[1])
			case "country", "Country":
				tmp.Country = strings.TrimSpace(col[1])
			case "descr", "Organization", "organisation", "owner":
				if tmp.Admin == "" {
					tmp.Admin = strings.TrimSpace(col[1])
				}
			}
		}
		if tmp.IpRange != "" {
			s.ParseResult = append(s.ParseResult, tmp)
		}
	}
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

func (s *Summary) BreakDownWhoisResponseIntoParagraphs() []string {
	switch s.WhoisResponseServer {
	case "whois.arin.net":
		return strings.Split(s.WhoisResponse, "#")
	case "whois.apnic.net", "whois.ripe.net", "whois.lacnic.net":
		return strings.Split(s.WhoisResponse, "%")
	default:
		return strings.Split(s.WhoisResponse, "\n\n")
	}
}

func (s *Summary) SetWhoisResponseServerFromWhoisResponse() {
	tmp := strings.Split(s.WhoisResponse, "\n")
	for _, v := range tmp {
		col := strings.Split(v, ":")
		if col[0] == "refer" {
			s.WhoisResponseServer = strings.TrimSpace(col[1])
			break
		}
	}
}

func (s *Summary) ParseCheck() bool {
	// 転送されていないかチェックする
	list := []string{"apnic", "arin", "ripe", "lacnic"}
	for _, v := range list {
		if strings.Contains(strings.ToLower(s.ParseResult[1].NetName), v) {
			s.WhoisResponseServer = fmt.Sprintf("whois.%s.net", v)
			return false
		}
	}
	return true
}
