package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!<br />Go script is written in utf8.<br />テスト")
	fmt.Fprint(w, "<p>recently deployed at 2015-03-25</p>")
}

// comment
