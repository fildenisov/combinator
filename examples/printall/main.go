package main

import (
	"context"
	"fmt"

	"github.com/fildenisov/combinator"
)

func main() {
	res := make(chan []int)
	go combinator.Recombo(context.Background(), []int{1, 2, 3, 4}, res)

	for r := range res {
		fmt.Println(r)
	}
}
