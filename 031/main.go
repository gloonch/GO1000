package main

import (
	"context"
	"errors"
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

	ctx, cancel := context.WithCancelCause(rootContext)

	go worker(ctx, 2)
	//cancel()
	go worker(ctx, 22)

	go func(cancel context.CancelCauseFunc) {
		time.Sleep(5 * time.Second)
		cancel(errors.New("context canceled by me"))
	}(cancel)

	fmt.Scanln()

}
