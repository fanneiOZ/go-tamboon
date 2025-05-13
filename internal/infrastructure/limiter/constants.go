package limiter

import "context"

var throttlerContext = context.WithValue(context.Background(), "group", "throttler")
