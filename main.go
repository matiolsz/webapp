package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r http.Request) {

}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("bites written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}
