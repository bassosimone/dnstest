// SPDX-License-Identifier: GPL-3.0-or-later

package dnstest

import (
	"net"
	"net/netip"
	"slices"
	"testing"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
)

func collectAddrs(resp *dns.Msg) (output []string) {
	for _, rec := range resp.Answer {
		switch rec := rec.(type) {
		case *dns.A:
			output = append(output, rec.A.String())
		case *dns.AAAA:
			output = append(output, rec.AAAA.String())
		}
	}
	slices.Sort(output)
	return
}

func TestUDP(t *testing.T) {
	// create config
	config := NewHandlerConfig()
	config.AddNetipAddr("www.example.com", netip.MustParseAddr("104.20.34.220"))
	config.AddNetipAddr("www.example.com", netip.MustParseAddr("172.66.144.113"))

	// create handler
	handler := NewHandler(config)

	// create server
	srv := MustNewUDPServer(&net.ListenConfig{}, "127.0.0.1:0", handler)
	defer srv.Close()

	// create query
	query := &dns.Msg{}
	query.Question = append(query.Question, dns.Question{
		Name:   dns.CanonicalName("www.example.com"),
		Qtype:  dns.TypeA,
		Qclass: dns.ClassINET,
	})

	// exchange
	resp, err := dns.Exchange(query, srv.Address())
	assert.NoError(t, err)

	// get results
	addrs := collectAddrs(resp)
	expect := []string{"104.20.34.220", "172.66.144.113"}
	assert.Equal(t, expect, addrs)
}
