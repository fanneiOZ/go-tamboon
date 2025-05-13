package limiter_test

import (
	"errors"
	"go-tamboon/internal/infrastructure/limiter"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRate(t *testing.T) {
	t.Parallel()
	t.Run("should create a new token properly", func(t *testing.T) {
		givenQuota := uint32(10)
		givenDuration := 1 * time.Second
		rate := limiter.NewRate(givenQuota, givenDuration)

		if quota, duration := rate.Settings(); quota != givenQuota || duration != givenDuration {
			t.Errorf("quota or duration are incorrect. received %v, %v", quota, duration)
		}

		if rate.Disposed() == true {
			t.Errorf("new rate should not be disposed")
		}
	})

	t.Run("should determine the limit allocation within time window", func(t *testing.T) {
		wg := sync.WaitGroup{}
		rate := limiter.NewRate(10, 1*time.Second)
		passed, failed := int32(0), int32(0)
		nRequests := 50
		wg.Add(nRequests)

		for i := 1; i <= nRequests; i++ {
			go func() {
				result := rate.Allocate()
				if result {
					atomic.AddInt32(&passed, 1)
				} else {
					atomic.AddInt32(&failed, 1)
				}

				wg.Done()
			}()
		}

		wg.Wait()

		if passed != 10 {
			t.Errorf("allocation should pass 10 times. received %d", passed)
		}

		if expectedFailed := nRequests - int(passed); int(failed) != expectedFailed {
			t.Errorf("allocation should fail %d times. received %v", expectedFailed, failed)
		}
	})

	t.Run("should reset the rate and allocate after time window", func(t *testing.T) {
		wg := sync.WaitGroup{}
		rate := limiter.NewRate(10, 500*time.Millisecond)
		passed, failed := int32(0), int32(0)
		nRequests := 50
		wg.Add(nRequests)
		for i := 1; i <= nRequests; i++ {
			go func() {
				result := rate.Allocate()
				if result {
					atomic.AddInt32(&passed, 1)
				} else {
					atomic.AddInt32(&failed, 1)
				}

				wg.Done()
			}()
		}

		wg.Wait()
		time.Sleep(510 * time.Millisecond)

		result := rate.Allocate()

		if result == false {
			t.Error("allocation should pass after time window reset")
		}

		if passed != 10 {
			t.Errorf("allocation should pass 10 times. received %d", passed)
		}

		if expectedFailed := nRequests - int(passed); int(failed) != expectedFailed {
			t.Errorf("allocation should fail %d times. received %v", expectedFailed, failed)
		}
	})

	t.Run("should not allocate when rate is disposed", func(t *testing.T) {
		rate := limiter.NewRate(10, 1*time.Second)

		_ = rate.Dispose()
		result := rate.Allocate()

		if result == true {
			t.Error("rate should not allocate after disposed")
		}
	})

	t.Run("should be able to disposed once", func(t *testing.T) {
		rate := limiter.NewRate(10, 1*time.Second)

		err1stRun := rate.Dispose()
		disposed := rate.Disposed()
		err2ndRun := rate.Dispose()

		if disposed == false {
			t.Error("token should be visible as disposed")
		}

		if err1stRun != nil {
			t.Errorf("first dispose should not return error. received %v", err1stRun)
		}

		if !errors.Is(err2ndRun, limiter.ErrRateAlreadyDisposed) {
			t.Errorf("second dispose should return error. received %v", err2ndRun)
		}
	})

	t.Run("should be able to assign parent once", func(t *testing.T) {
		rate := limiter.NewRate(10, 1*time.Second)

		err1stRun := rate.AssignParent(&limiter.Throttler{})
		err2ndRun := rate.AssignParent(&limiter.Throttler{})

		if err1stRun != nil {
			t.Errorf("first assign parent should not return error. received %v", err1stRun)
		}

		if !errors.Is(err2ndRun, limiter.ErrRateIsAlreadyAssigned) {
			t.Errorf("second assign parent should return error. received %v", err2ndRun)
		}
	})
}
