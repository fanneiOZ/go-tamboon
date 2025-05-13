package limiter_test

import (
	"github.com/google/uuid"
	"go-tamboon/internal/infrastructure/limiter"
	"sync"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	t.Run("task", func(t *testing.T) {
		t.Parallel()
		t.Run("should create a new token properly", func(t *testing.T) {
			token := limiter.NewToken(func() {})

			if _, err := uuid.Parse(token.Id()); err != nil {
				t.Errorf("token id should be a valid uuid. received: %s", err.Error())
			}

			if token.Done() == true {
				t.Errorf("token should not be done. received: %v", token.Done())
			}

			if token.Status() == true {
				t.Errorf("token should not be succeeded. received: %v", token.Status())
			}
		})

		t.Run("should create a new timeout token properly", func(t *testing.T) {
			token := limiter.NewTokenWithTimeout(func() {}, 2*time.Second)

			if _, err := uuid.Parse(token.Id()); err != nil {
				t.Errorf("token id should be a valid uuid. received: %s", err.Error())
			}

			if token.Timeout() == 0 {
				t.Error("token should have a timeout")
			}

			if token.Context() == nil {
				t.Error("token should have a context")
			}

			if token.Done() == true {
				t.Errorf("token should not be done. received: %v", token.Done())
			}

			if token.Status() == true {
				t.Errorf("token should not be succeeded. received: %v", token.Status())
			}
		})

		t.Run("should hit the timeout", func(t *testing.T) {
			givenTimeout := 50 * time.Millisecond
			spyTaskIsCalled := false
			spyTask := func() {
				spyTaskIsCalled = true
			}
			token := limiter.NewTokenWithTimeout(spyTask, givenTimeout)
			time.Sleep(givenTimeout + 10*time.Millisecond)

			if token.Done() == false {
				t.Errorf("token should be done after timeout. received: %v", token.Done())
			}

			if spyTaskIsCalled == true {
				t.Errorf("spyTask should not be called")
			}

			if token.Status() == true {
				t.Errorf("token should not be succeeded. received: %v", token.Status())
			}
		})

		t.Run("should not execute task when create token", func(t *testing.T) {
			spyTaskIsCalled := false
			spyTask := func() {
				spyTaskIsCalled = true
			}
			limiter.NewToken(spyTask)

			if spyTaskIsCalled == true {
				t.Error("token task should not be called")
			}
		})

		t.Run("should execute task when resume", func(t *testing.T) {
			wg := sync.WaitGroup{}
			wg.Add(1)
			spyTaskIsCalled := false
			spyTask := func() {
				spyTaskIsCalled = true
				wg.Done()
			}

			token := limiter.NewToken(spyTask)
			token.Resume()
			wg.Wait()

			if token.Done() == false {
				t.Errorf("token should be marked done. received: %v", token.Done())
			}

			if token.Status() == false {
				t.Errorf("token should succeeded. received: %v", token.Status())
			}

			if spyTaskIsCalled == false {
				t.Error("token task should be called")
			}
		})

		t.Run("should execute task for token with timeout when resume", func(t *testing.T) {
			wg := sync.WaitGroup{}
			wg.Add(1)
			spyTaskIsCalled := false
			spyTask := func() {
				spyTaskIsCalled = true
				wg.Done()
			}

			token := limiter.NewTokenWithTimeout(spyTask, 10*time.Second)
			token.Resume()
			wg.Wait()

			if token.Done() == false {
				t.Errorf("token should be marked done. received: %v", token.Done())
			}

			if token.Status() == false {
				t.Errorf("token should succeeded. received: %v", token.Status())
			}

			if spyTaskIsCalled == false {
				t.Error("token task should be called")
			}
		})

		t.Run("should not be timeout when token is resumed and program still alive", func(t *testing.T) {
			givenDelay := 50 * time.Millisecond
			spyTaskIsCalled := false
			spyTask := func() {
				spyTaskIsCalled = true
				time.Sleep(givenDelay)
			}

			givenTimeout := 10 * time.Millisecond
			token := limiter.NewTokenWithTimeout(spyTask, givenTimeout)
			token.Resume()
			time.Sleep(givenTimeout + givenDelay)

			if token.Done() == false {
				t.Errorf("token should be marked done. received: %v", token.Done())
			}

			if token.Status() == false {
				t.Errorf("token should succeeded. received: %v", token.Status())
			}

			if spyTaskIsCalled == false {
				t.Error("token task should be called")
			}
		})
	})
}
