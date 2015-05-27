package hello

import (
	"fmt"
	"go-diff/diffmatchpatch"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
	fmt.Fprintln(w, "recently deployed at 2015-05-27 circle ci")
	fmt.Fprintln(w, "using submodule go-diff below\n")

	dmp := diffmatchpatch.New()
	a, b, c := dmp.DiffLinesToChars("string before", "string after")
	diffs := dmp.DiffMain(a, b, false)
	result := dmp.DiffCharsToLines(diffs, c)
	fmt.Fprintln(w, result)

}

// comment
