package dns

import (
	"fmt"
	"sort"

	"github.com/lixiangzhong/dnsutil"
)

// Dig digs
func Dig(edgeDNSServer, fqdn string) ([]string, error) {
	var dig dnsutil.Dig
	if edgeDNSServer == "" {
		return nil, fmt.Errorf("empty edgeDNSServer")
	}
	err := dig.SetDNS(edgeDNSServer)
	if err != nil {
		return nil, err
	}
	a, err := dig.A(fqdn)
	if err != nil {
		return nil, err
	}
	var IPs []string
	for _, ip := range a {
		IPs = append(IPs, fmt.Sprint(ip.A))
	}
	sort.Strings(IPs)
	return IPs, nil
}
