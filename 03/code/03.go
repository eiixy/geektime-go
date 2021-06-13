package code

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func Run03() {
	d := time.Now().Add(shortDuration)

	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	default:
		fmt.Println("default")
	}
}
