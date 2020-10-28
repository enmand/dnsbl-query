package main

import (
	"os"

	"github.com/enmand/dnsbl-query/internal/cmd/graphql"
	"github.com/jessevdk/go-flags"
)

func main() {
	p := flags.NewParser(nil, flags.Default)
	p.SubcommandsOptional = true

	graphql.Register(p)

	if _, err := p.Parse(); err != nil {
		os.Exit(1)
	}
}
