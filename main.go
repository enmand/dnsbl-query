package main

import (
	"os"

	"github.com/enmand/dnsbl-query/internal/cmd/admin"
	"github.com/enmand/dnsbl-query/internal/cmd/graphql"
	"github.com/enmand/dnsbl-query/internal/cmd/worker"
	"github.com/jessevdk/go-flags"
)

func main() {
	p := flags.NewParser(nil, flags.Default)
	p.SubcommandsOptional = true

	worker.Register(p)
	graphql.Register(p)
	admin.Register(p)

	if _, err := p.Parse(); err != nil {
		os.Exit(1)
	}
}
