package dnsbl

import "context"

// Codes represent error codes and their meanings
type Code string

const (
	CodeSpamhausSBL       = "127.0.0.2"
	CodeSpamhausSBLCSS    = "127.0.0.3"
	CodeSpamhausCBL       = "127.0.0.4"
	CodeSpamhausDROP      = "127.0.0.10"
	CodeISP               = "127.0.0.10"
	CodeSpamhaus          = "127.0.0.11"
	CodeErrorTyping       = "127.255.255.252"
	CodeErrorOpenResolver = "127.255.255.254"
	CodeRateLimit         = "127.255.255.255"
	CodeUnknown           = "0.0.0.0"
)

// Response represents a Response from the DNSBL service
type Response struct {
	Codes []Code
}

// newResponse returns a new empty response
func newResponse() *Response {
	return &Response{
		Codes: []Code{},
	}
}

// DNSBL represents a service interface for interacting with DNSBL services
type DNSBL interface {
	Query(ctx context.Context, ip string) (*Response, error)
}
