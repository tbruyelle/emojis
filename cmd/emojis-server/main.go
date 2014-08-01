package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tbruyelle/emojis"
)

var ems emojis.Emojis

func init() {
	var err error
	ems, err = emojis.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/random", randomHandler)
	panic(http.ListenAndServe(":8080", nil))
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", ems.Random())
}
