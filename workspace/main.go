package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"workspace/internal"
)

func main() {

	fmt.Print()
	internal.MyFunc()

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		internal.EnableCors(w)
		if r.Method == "OPTIONS" {
			return
		}
		res, err := internal.ReturnResponse(r)
		if err != nil {
			fmt.Println("Error:", err)
		} else {

			fmt.Println(res)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(strconv.Itoa(res)))
			return
		}
	})
	url.Parse("")
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
