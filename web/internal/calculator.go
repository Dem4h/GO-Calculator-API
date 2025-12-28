package internal

import (
	"fmt"
	"net/http"
	"net/url"
)

func myFunc() {
	fmt.Println("test")
}

func main() {
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request")
	})
	url.Parse("")
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
