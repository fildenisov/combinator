package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fildenisov/combinator"
)

func main() {
	res := make(chan []rune)
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)

	go combinator.Recombo(ctx, []rune("1234567890"), res)

	for r := range res {
		fmt.Println(r)
	}
}
