package runner

import (
	"context"
	"log"
	"sync"

	"github.com/autoar/internal/config"
)

// Task represents a single recon task to be executed.
type Task struct {
	Target string
	Type   string
}

// Result holds the output of a completed task.
type Result struct {
	Task   Task
	Output string
	Err    error
}

// Runner manages concurrent execution of recon tasks.
type Runner struct {
	cfg     *config.Config
	tasks   chan Task
	results chan Result
	wg      sync.WaitGroup
}

// New creates a new Runner with the given config.
func New(cfg *config.Config) *Runner {
	return &Runner{
		cfg:     cfg,
		tasks:   make(chan Task, cfg.Concurrency*2),
		results: make(chan Result, cfg.Concurrency*2),
	}
}

// Start launches worker goroutines and begins processing tasks.
func (r *Runner) Start(ctx context.Context) {
	for i := 0; i < r.cfg.Concurrency; i++ {
		r.wg.Add(1)
		go r.worker(ctx)
	}
}

// Submit enqueues a task for execution.
func (r *Runner) Submit(t Task) {
	r.tasks <- t
}

// Results returns the read-only results channel.
func (r *Runner) Results() <-chan Result {
	return r.results
}

// Stop closes the task channel and waits for all workers to finish.
func (r *Runner) Stop() {
	close(r.tasks)
	r.wg.Wait()
	close(r.results)
}

func (r *Runner) worker(ctx context.Context) {
	defer r.wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Printf("worker shutting down: %v", ctx.Err())
			return
		case t, ok := <-r.tasks:
			if !ok {
				return
			}
			output, err := execute(ctx, t)
			r.results <- Result{Task: t, Output: output, Err: err}
		}
	}
}

// execute runs the task and returns its output.
func execute(ctx context.Context, t Task) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	log.Printf("executing task type=%s target=%s", t.Type, t.Target)
	// Placeholder: real tool invocation goes here.
	return "ok", nil
}
