package limiter

import (
	"errors"
	"sync/atomic"
	"time"
)

type Rate struct {
	current  uint32
	disposed uint32
	quota    uint32

	previousWindowAt int64
	window           time.Duration
	parent           *Throttler
}

var (
	ErrRateAlreadyDisposed   = errors.New("rate is already disposed")
	ErrRateIsAlreadyAssigned = errors.New("rate is already assigned")
)

func NewRate(quota uint32, window time.Duration) *Rate {
	return &Rate{quota: quota, window: window, current: 0, previousWindowAt: time.Now().UnixNano()}
}

func (r *Rate) Settings() (uint32, time.Duration) {
	return r.quota, r.window
}

func (r *Rate) Allocate() bool {
	if r.Disposed() {
		return false
	}

	currentTimestamp := time.Now().UnixNano()
	if currentTimestamp-atomic.LoadInt64(&r.previousWindowAt) >= r.window.Nanoseconds() {
		atomic.StoreInt64(&r.previousWindowAt, currentTimestamp)
		atomic.StoreUint32(&r.current, 0)
	}

	return atomic.AddUint32(&r.current, 1) <= r.quota
}

func (r *Rate) AssignParent(parent *Throttler) error {
	if r.parent != nil {
		return ErrRateIsAlreadyAssigned
	}

	r.parent = parent

	return nil
}

func (r *Rate) Dispose() error {
	if r.Disposed() {
		return ErrRateAlreadyDisposed
	}

	atomic.StoreUint32(&r.disposed, 1)

	return nil
}

func (r *Rate) Disposed() bool {
	return atomic.LoadUint32(&r.disposed) == 1
}
