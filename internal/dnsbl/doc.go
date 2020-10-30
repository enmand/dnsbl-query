// Package dnsbl provides an interface to query DNSBL services, and any supported
// implementations.
//
// Currently only Spamhaus is directly supported. Spamhaus can be configured to
// use the system DNS resolvers, Go resolver, or a customized resolver that allows
// configuration of the DNS server to use to preform the request.
package dnsbl
