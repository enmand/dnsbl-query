package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/gammazero/workerpool"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/operation"
)

const Workers = 100

// Worker represents a worker that can select tasks from the database que, and
// insert them to be worked on, tracking their state to completion.
//
// IMPORTANT NOTE: In the "real work" this would be a much more robust and dynamic
// worker system, probably using Cadence, but in the interest of time and
// simplicity, this worker will work fine. This is not safe for production use.
type Worker interface {
	Start() error
	Stop() error
}
type worker struct {
	client *ent.Client
	pool   *workerpool.WorkerPool
}

func New(cl *ent.Client) Worker {
	pool := workerpool.New(Workers)

	return &worker{
		client: cl,
		pool:   pool,
	}
}

// Start begins processing background jobs from the database
func (w *worker) Start() error {
	for {
		ctx := context.Background()
		op, err := w.client.Operation.Query().
			Where(
				operation.TypeEQ(operation.TypeIPDNSBL),
				operation.StatusEQ(operation.StatusWAITING),
			).First(ctx)
		if ent.IsNotFound(err) {
			// sleep and wait for the next job
			time.Sleep(200 * time.Millisecond)
			continue
		}

		fmt.Printf("%# v\n\n", op)
		time.Sleep(1 * time.Second)
	}

	return nil
}

// Stop stops the worker pool gracefully if possible
func (w *worker) Stop() error {
	return nil
}
