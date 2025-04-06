package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Base context
	baseCtx := context.Background()
	fmt.Println("🌍 Base context:", baseCtx)

	// Add a value to the context
	ctxWithValue := context.WithValue(baseCtx, "username", "srikanth")
fmt.Println("🔑 Context with value:", ctxWithValue)
	// Manual cancellation context
	ctxWithCancel, cancelFunc := context.WithCancel(ctxWithValue)

	// Set a timeout context
	ctxWithTimeout, timeoutCancel := context.WithTimeout(ctxWithCancel, 3*time.Second)

	// Set an absolute deadline (5 seconds from now)
	deadline := time.Now().Add(5 * time.Second)
	ctxWithDeadline, deadlineCancel := context.WithDeadline(ctxWithTimeout, deadline)

	// Cleanup
	defer cancelFunc()
	defer timeoutCancel()
	defer deadlineCancel()

	// Run worker with full context
	doSomething(ctxWithDeadline)
}

func doSomething(ctx context.Context) {
	// Get value
	if username := ctx.Value("username"); username != nil {
		fmt.Println("🔑 Username from context:", username)
	}

	// Simulate cancel scenario (uncomment to see it in action)
	// go func() {
	//     time.Sleep(1 * time.Second)
	//     fmt.Println("🛑 Cancelling context manually...")
	//     cancelFunc() // Would need to be passed in to do this
	// }()

	// Wait on either work completion or cancellation/deadline
	select {
	case <-time.After(4 * time.Second):
		fmt.Println("✅ Work finished without interruption")
	case <-ctx.Done():
		fmt.Println("⚠️ Context canceled or deadline exceeded:", ctx.Err())
	}
}
