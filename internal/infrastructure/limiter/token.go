package limiter

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Token struct {
	id     uuid.UUID
	task   func()
	done   bool
	status bool

	createdAt time.Time

	timeout time.Duration
	cancel  context.CancelFunc

	context context.Context
}

func NewTokenWithTimeout(task func(), timeout time.Duration) *Token {
	token := &Token{task: task, id: uuid.New(), createdAt: time.Now(), timeout: timeout}
	token.Context()
	go func() {
		for !token.done {
			<-token.context.Done()
			token.done = true
		}
	}()

	return token
}

func NewToken(task func()) *Token {
	return &Token{task: task, id: uuid.New(), createdAt: time.Now()}
}

func (t *Token) Id() string {
	return t.id.String()
}

func (t *Token) Resume() {
	go func() {
		t.task()
		t.done = true
		t.status = true
		if t.cancel != nil {
			t.cancel()
		}
	}()
}

func (t *Token) Done() bool {
	return t.done
}

func (t *Token) Status() bool {
	return t.status
}

func (t *Token) Timeout() time.Duration {
	return t.timeout
}

func (t *Token) Context() context.Context {
	if t.context != nil {
		return t.context
	}

	t.context = context.WithValue(throttlerContext, "meta", ContextMetadata{
		Object:    "token",
		Id:        t.Id(),
		Timestamp: t.createdAt,
	})
	if t.timeout > 0 {
		contextWithTimeout, cancel := context.WithTimeout(t.context, t.timeout)
		t.context = contextWithTimeout
		t.cancel = cancel
	}

	return t.context
}

type ContextMetadata struct {
	Object    string
	Id        string
	Timestamp time.Time
}
