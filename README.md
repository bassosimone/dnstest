# Golang DNS test helpers

[![GoDoc](https://pkg.go.dev/badge/github.com/bassosimone/dnstest)](https://pkg.go.dev/github.com/bassosimone/dnstest) [![Build Status](https://github.com/bassosimone/dnstest/actions/workflows/go.yml/badge.svg)](https://github.com/bassosimone/dnstest/actions) [![codecov](https://codecov.io/gh/bassosimone/dnstest/branch/main/graph/badge.svg)](https://codecov.io/gh/bassosimone/dnstest)

The `dnstest` Go package is like [net/http/httptest](https://pkg.go.dev/net/http/httptest)
but focuses on testing DNS clients using various protocols.

Basic usage is like:

```Go
import (
	"log"
	"net"

	"github.com/bassosimone/dnstest"
)

// 1. create handler config and handler
config := dnstest.NewHandlerConfig()
config.AddCNAME("www.example.com", "example.com")
config.AddNetipAddr("example.com", netip.MustParseAddr("104.20.34.220"))
handler := dnstest.NewHandler(config)

// 2a. DNS-over-UDP server
srv := MustNewUDPServer(&net.ListenConfig{}, "127.0.0.1:0", handler)
log.Print(srv.Address()) // UDP address to use

// 3b. DNS-over-TCP server
srv := MustNewTCPServer(&net.ListenConfig{}, "127.0.0.1:0", handler)
log.Print(srv.Address()) // TCP address to use

// 3c. DNS-over-TLS server
cert := tls.Certificate{} // TODO: configure
srv := MustNewTLSServer(&net.ListenConfig{}, "127.0.0.1:0", cert, handler)
log.Print(srv.Address()) // TCP address to use

// 3d. DNS-over-HTTPS server
srv := MustNewHTTPSServer(&net.ListenConfig{}, "127.0.0.1:0", cert, handler)
log.Print(srv.URL()) // URL to use

// 4. Close when done
defer srv.Close()
```

## Features

- **Supports multiple protocols:** Currently, UDP, TCP, TLS, and HTTPS.

- **Supports multiple query types:** Currently, A, AAAA, and CNAME.

- **Compatible with pkitest:** Can use [github.com/bassosimone/pkitest](
https://pkg.go.dev/github.com/bassosimone/pkitest) to generate self-signed
certificates and cert pools.

- **Concurrency safe:** Safe for concurrent use in parallel tests.

- **Test friendly:** Panic on failure to avoid unnecessary `if err != nil` checks.

## Installation

To add this package as a dependency to your module:

```sh
go get github.com/bassosimone/dnstest
```

## Development

To run the tests:
```sh
go test -v .
```

To measure test coverage:
```sh
go test -v -cover .
```

## License

```
SPDX-License-Identifier: GPL-3.0-or-later
```

## History

Adapted from [ooni/netem](https://github.com/ooni/netem).
