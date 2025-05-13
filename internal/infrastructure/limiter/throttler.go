package limiter

import (
	"time"
)

type ThrottlerCommand string

const (
	notifyCommand ThrottlerCommand = "notify"
	dummyCommand  ThrottlerCommand = "dummy"
)

type Throttler struct {
	limit       int32
	timeWindow  time.Duration
	commands    chan ThrottlerCommand
	isAvailable bool
	//tokens      map[string]*Token
	tokens   *tokenDeque
	disposed bool
}

func NewThrottler(limit int32, timeWindow time.Duration) *Throttler {
	throttler := &Throttler{
		limit:      limit,
		timeWindow: timeWindow,
		commands:   make(chan ThrottlerCommand),
		tokens:     newTokenDeque(),
	}

	go func(ch <-chan ThrottlerCommand) {
		for !throttler.disposed {
			select {
			case cmd := <-ch:
				if cmd == notifyCommand {
					throttler.notify()
				}
			}
		}
	}(throttler.commands)

	return throttler
}

func (throttler *Throttler) notify() {

}

func (throttler *Throttler) SendRequest(task func()) {
	if throttler.isAvailable {
		go task()
	}

	token := NewToken(task)
	throttler.tokens.append(token)
}

func (throttler *Throttler) Dispose() {
	throttler.disposed = true
	throttler.commands <- dummyCommand
}
