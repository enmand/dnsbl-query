// Package worker provides a very basic, naive background worker library.
//
// The worker will continually poll for operations, and if an Operation is
// found with a matching Job definition, the worker works the Job, and updates
// the operational status.
//
// This worker should not be used in production, and something like Cadence should
// be used instead.
package worker
