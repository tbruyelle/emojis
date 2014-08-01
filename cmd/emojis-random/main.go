package main

import (
	"fmt"
	"os"

	"github.com/tbruyelle/emojis"
)

func main() {
	emojis, err := emojis.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(emojis.Random())
}
