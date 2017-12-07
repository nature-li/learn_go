package common

import (
	"errors"
	"os"
	"time"
	"os/signal"
)

var (
	ErrTimeOut   = errors.New("执行者执行超时")
	ErrInterrupt = errors.New("执行者被中断")
)

type Runner struct {
	tasks     []func(int)
	timeout   <-chan time.Time
	interrupt chan os.Signal
}

func NewRunner(tm time.Duration) *Runner {
	return &Runner{
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}

		if r.isTimeout() {
			return ErrTimeOut
		}

		task(id)
	}

	return nil
}

func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) isTimeout() bool {
	select {
	case <-r.timeout:
		return true
	default:
		return false
	}
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	complete := make(chan error)

	go func() {
		complete <- r.run()
	}()

	select {
	case err := <-complete:
		return err
	}
}
