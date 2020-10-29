package worker

import (
	"github.com/gammazero/workerpool"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
)

const Workers = 100

// Worker represents a worker that can select tasks from the database que, and
// insert them to be worked on, tracking their state to completion.
//
// IMPORTANT NOTE: In the "real work" this would be a much more robust and dynamic
// worker system, probably using Cadence, but in the interest of time and
// simplicity, this worker will work fine.
type Worker struct {
	client *ent.Client
	pool   *workerpool.WorkerPool
}

func New(cl *ent.Client) *Worker {
	pool := workerpool.New(Workers)

	return &Worker{
		client: cl,
		pool:   pool,
	}
}
