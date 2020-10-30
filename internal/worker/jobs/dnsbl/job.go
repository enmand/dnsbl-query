package job

import (
	"context"
	"fmt"
	"time"

	"github.com/enmand/dnsbl-query/internal/dnsbl"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/ip"
)

// DNSBLJob represents a worker.Job to query a DNSBL service
type DNSBLJob struct {
	cl    *ent.Client
	dnsbl dnsbl.DNSBL
}

// NewDNSBLJob returns a new worker job for querying a DNSBL service
func NewDNSBLJob(cl *ent.Client, dnsbl dnsbl.DNSBL) *DNSBLJob {
	return &DNSBLJob{
		cl:    cl,
		dnsbl: dnsbl,
	}
}

// Execute a DNSBL query and store the results on our graph
func (s *DNSBLJob) Execute(ctx context.Context, ipaddr string) error {
	tx, err := s.cl.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }() // rollback on failure

	// TODO: this should be separate retryable states or workflows
	ipe, err := tx.IP.Query().Where(
		ip.IPAddressEQ(ipaddr),
	).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		ipe, err = tx.IP.Create().
			SetIPAddress(ipaddr).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("ceating new ip: %w", err)
		}
	case err != nil:
		return fmt.Errorf("ip lookup: %w", err)
	default:
		_, err = ipe.Update().SetUpdatedAt(time.Now()).Save(ctx)
		if err != nil {
			return fmt.Errorf("ip updated at: %w", err)
		}
	}

	// TODO we should check here to make sure some time passed (e.g. 1 day)
	// before we check an IP again

	query, err := tx.DNSBLQuery.Create().
		SetIPAddress(ipe).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("query update: %w", err)
	}

	resp, err := s.dnsbl.Query(ctx, ipaddr)
	if err != nil {
		return fmt.Errorf("job execution: %w", err)
	}

	for _, code := range resp.Codes {
		_, err := tx.DNSBLResponse.Create().
			SetQuery(query).
			SetCode(string(code)).
			SetDescription("spamhaus response").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("saving response: %w", err)
		}
	}

	return tx.Commit()
}
