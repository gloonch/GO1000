package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: context canceled with error: %s, Exiting... \n", id, context.Cause(ctx))

			return
		default:
			fmt.Printf("worker %d: processing message... \n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	rootContext := context.Background()

	timeoutCtx, _ := context.WithTimeout(rootContext, time.Second*10)
	deadlineCtx, _ := context.WithDeadline(timeoutCtx, time.Now().Add(time.Second*3))
	valueCtx := context.WithValue(timeoutCtx, "foo", "bar")

	go worker(timeoutCtx, 2)
	go worker(deadlineCtx, 22)
	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Printf("worker with value %v: context canceled... \n", ctx.Value("foo"))
	}(valueCtx)
	fmt.Scanln()

}
