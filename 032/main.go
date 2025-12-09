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

	ctx, _ := context.WithTimeout(rootContext, time.Second*10)
	deadlineCtx, _ := context.WithDeadline(ctx, time.Now().Add(time.Second*3))

	go worker(ctx, 2)
	go worker(deadlineCtx, 22)

	fmt.Scanln()

}
