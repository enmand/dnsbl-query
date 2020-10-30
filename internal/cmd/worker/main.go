package worker

import (
	"fmt"
	"log"

	gflags "github.com/jessevdk/go-flags"
	"go.uber.org/zap"

	"github.com/enmand/dnsbl-query/internal/dnsbl"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/operation"
	"github.com/enmand/dnsbl-query/internal/flags"
	"github.com/enmand/dnsbl-query/internal/worker"
	job "github.com/enmand/dnsbl-query/internal/worker/jobs/dnsbl"
)

type cmd struct {
	flags.Database
}

// Register registers the worker runner with the flags parser given
func Register(p *gflags.Parser) {
	_, err := p.AddCommand("worker", "worker service", "", &cmd{})
	if err != nil {
		log.Fatal(err)
	}
}

// Execute initializes and starts a new worker pool
func (c *cmd) Execute(args []string) error {
	cl, err := ent.Open(c.DatabaseDriver, c.DatabaseURI)
	if err != nil {
		return fmt.Errorf("opening database: %w", err)
	}

	log, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("getting logger: %w", err)
	}

	wp := worker.New(cl, log.Sugar())
	if err != nil {
		return fmt.Errorf("initializing worker: %w", err)
	}

	sh := dnsbl.NewSpamhaus()
	j := job.NewDNSBLJob(cl, sh)

	wp.Register(operation.TypeIPDNSBL, j)

	return wp.Start()
}
