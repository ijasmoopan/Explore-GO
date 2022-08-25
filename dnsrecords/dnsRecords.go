package main

import (
	"fmt"
	"net"
)

func main() {
	// ------DNS--------
	// DNS records are mapping files that associate with DNS server whichever IP addresses each domain
	// is associated with, and they handle requests sent to each domain.
	// The net package contains various methods to find general details of DNS records.
	// Let's run some examples, to collect information about the DNS servers and the corresponding
	// records of a target domain:


	// IPrecords : The net.LookupIP() function accepts a string(domain-name) and 
	// returns a slice of net.IP objects that contains host's IPv4 and IPv6 addresses.
	// iprecords, _ := net.LookupIP("google.com")
	// for _, ip := range iprecords {
	// 	fmt.Println(ip)
	// }
	fmt.Println("Printing IP records of facebook.com")
	iprecords , _ := net.LookupIP("facebook.com")
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	// CNAME : CNAMEs are essentially domain and subdomain text aliases to bind traffic. 
	// The net.LookupCNAME() function accepts a host-name(m.facebook.com) as a string and 
	// returns a single canonical domain name for the given host.
	fmt.Println("Printing CNAME of facebook.com")
	cname , _ := net.LookupCNAME("m.facebook.com")
	fmt.Println(cname)

	// PTR (pointer) : These records provide the reverse binding from addresses to names. 
	// PTR records should exactly match the forward maps. 
	// The net.LookupAddr() function performs a reverse finding for the address and 
	// returns a list of names that map to the given address.
	fmt.Println("Printing Address")
	ptr, _ := net.LookupAddr("6.8.8.8")
	for _, ptrvalue := range ptr {
		fmt.Println(ptrvalue)
	}


	// Name Server : The NS records describe the authorized name servers for the zone.
	// The NS also delegates subdomains to other organizations on zone files.
	// The net.LookupNS() function takes a domain name(facebook.com) as a string and
	// returns DNS-NS records as a slice of NS structs.
	fmt.Println("Printing Name Server:")
	nameserver, _ := net.LookupNS("facebook.com")
	for _, ns := range nameserver {
		fmt.Println(ns)
	}

	// MX records : These records identify the servers that can exchange emails.
	// The net.LookupMX() function takes a domain name as a string and
	// returns a slice of MX structs sorted by preference.
	// An MX struct is made up of a Host as a string and Pref as a uint16.
	fmt.Println("Printing MX records:")
	mxrecords, _ := net.LookupMX("facebook.com")
	for _, mx := range mxrecords {
		fmt.Println(mx.Host, mx.Pref)
	}

	// SRV service record : The LookupSRV function tries to resolve an SRV query
	// of the given service, protocol, and domain name.
	// The second parameter is "tcp" or "udp".
	// The returned records are sorted by priority and randomized by weight within a priority.
	fmt.Println("Printing Service record:")
	cname, srvs, err := net.LookupSRV("xmpp-server", "tcp", "golang.org")
	if err != nil {
		panic(err)
	}
 
	fmt.Printf("\ncname: %s \n\n", cname)
 
	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}

	// TXT records : This text record stores information about the SPF that can identify the
	// authorized server to send email on behalf of your organization.
	// The net.LookupTXT() function takes a domain name(facebook.com) as a string and
	// returns a list of DNS TXT records as a slice of strings.
	fmt.Println("Printing TXT records:")
	txtrecords, _ := net.LookupTXT("facebook.com")
 
	for _, txt := range txtrecords {
		fmt.Println(txt)
	}
}