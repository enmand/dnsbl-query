package dnsbl

import (
	"context"
	"fmt"
	"net"
	"strings"
)

// spamhaus is the service for making dnsbl queries
type spamhaus struct {
	network  string
	resolver resolver
}

// options are the options of the spamhaus dnsbl querier
type options struct {
	network    string
	dnsServer  string
	goResolver bool
}

// Option is the public interface for optional parameters
type Option func(*options)

// resolver is a resolver type that can look up an IP address set from a hostname
type resolver interface {
	LookupIP(context.Context, string, string) ([]net.IP, error)
}

// WithDNSServer sets a custom DNS server to use
func WithDNSServer(server string) Option {
	return func(o *options) {
		o.dnsServer = server
	}
}

func NewSpamhaus(opts ...Option) DNSBL {
	options := &options{}
	for _, opt := range opts {
		opt(options)
	}
	if options.network == "" {
		options.network = "ip"
	}

	r := &net.Resolver{
		PreferGo: options.goResolver,
	}
	if options.dnsServer != "" {
		r.Dial = customDialer(options.network, options.dnsServer)
	}

	return &spamhaus{
		network:  options.network,
		resolver: r,
	}
}

// Query does a DNSBL query for a single IP address to Spamhaus
func (sh *spamhaus) Query(ctx context.Context, ip string) (*Response, error) {
	ip = reverseIP(ip)
	ip = fmt.Sprintf("%s.zen.spamhaus.org", ip)

	resp, err := sh.resolver.LookupIP(ctx, sh.network, ip)
	if err != nil {
		return nil, fmt.Errorf("spamhaus query: %w", err)
	}

	r := newResponse()
	for _, ip := range resp {
		r.Codes = append(r.Codes, Code(ip.String()))
	}

	return r, nil
}

type dialerFunc func(ctx context.Context, network, addr string) (net.Conn, error)

// customDialer returns
func customDialer(network, server string) dialerFunc {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return net.Dial(network, server)
	}
}

// reverseIP returns a reverse-notation for given IP address
func reverseIP(str string) string {
	octs := strings.Split(str, ".")

	last := len(octs) - 1
	for i := 0; i < len(octs)/2; i++ {
		octs[i], octs[last-i] = octs[last-i], octs[i]
	}

	return strings.Join(octs, ".")
}
