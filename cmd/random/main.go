package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/tbruyelle/emojis"
)

func main() {
	emojis, err := emojis.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(emojis[r.Intn(len(emojis))])
}
