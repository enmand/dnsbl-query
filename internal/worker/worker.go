package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/gammazero/workerpool"
	"go.uber.org/zap"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/operation"
)

const Workers = 100

// Worker represents a worker that can select tasks from the database que, and
// insert them to be worked on, tracking their state to completion.
//
// IMPORTANT NOTE: In the "real work" this would be a much more robust and dynamic
// worker system, probably using Cadence, or a real state-machine to progress jobs
// but in the interest of time and simplicity, this worker will work fine.
// This is not safe for production use.
type Worker interface {
	Start() error
	Stop()
	Register(operation.Type, Job)
}
type worker struct {
	client *ent.Client
	pool   *workerpool.WorkerPool
	jobs   map[operation.Type]Job
	log    *zap.SugaredLogger
}

// Job represents an in-memory job for the worker to work on
// TODO: in the "real world" this should be polymophic and include state
type Job interface {
	Execute(ctx context.Context, state string) error
}

// State represents the state of a Job
type State interface {
	Get() map[string]interface{}
}

// New returns a new Worker pool
func New(cl *ent.Client, l *zap.SugaredLogger) Worker {
	pool := workerpool.New(Workers)

	return &worker{
		jobs:   make(map[operation.Type]Job),
		client: cl,
		pool:   pool,
		log:    l,
	}
}

// Start begins processing background jobs from the database
func (w *worker) Start() error {
	for {
		w.log.Info("checking for jobs")
		ctx := context.Background()
		op, err := w.client.Operation.Query().
			Where(
				operation.StatusEQ(operation.StatusWAITING),
			).First(ctx)
		if ent.IsNotFound(err) {
			w.log.Info("none found, sleeping")
			// sleep and wait for the next job
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if err != nil {
			w.log.Error(err.Error())
			return fmt.Errorf("finding operations: %w", err)
		}

		if j, ok := w.jobs[op.Type]; ok {
			w.log.Infof("job:%s found, submitting", op.Type)

			pj, err := w.prepJob(ctx, j, op)
			if err != nil {
				return fmt.Errorf("preparing job: %w", err)
			}
			w.pool.Submit(pj)
		}
	}
}

// Stop stops the worker pool gracefully if possible
func (w *worker) Stop() {
	w.pool.Stop()
}

// Register registers a new Job to be worked
func (w *worker) Register(opType operation.Type, j Job) {
	w.jobs[opType] = j
}

// prepJob prepares a Job to run run asyncronously
func (w *worker) prepJob(ctx context.Context, j Job, op *ent.Operation) (func(), error) {
	w.log.Infof("job:%s progressed", op.Type)

	_, err := op.Update().SetStatus(operation.StatusIN_PROGRESS).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("progressing operation: %w", err)
	}

	return func() {
		w.log.Infof("running job:%s", op.Type)
		ctx := context.Background() // each task gets it's own context
		err := j.Execute(ctx, op.IPAddress)

		if err != nil {
			w.log.Errorf("job:%s state:%s failed: %s", op.Type, op.IPAddress, err)
			_, _ = op.Update(). // TODO: we can't deal with the error
						SetStatus(operation.StatusERROR).
						SetError(err.Error()).
						Save(ctx)
		}

		w.log.Info("job worked")
		_, _ = op.Update().
			SetStatus(operation.StatusDONE).Save(ctx)
	}, nil
}
