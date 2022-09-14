package main

import (
	"context"
	"fmt"

	"github.com/fildenisov/combinator"
)

func main() {
	res := make(chan []rune)
	go combinator.Recombo(context.Background(), []rune("abcd"), res)

	for r := range res {
		fmt.Println(string(r))
	}
}
