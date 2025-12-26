package main

import (
	"fmt"
	"net/http"
	"strconv"
	"workspace/internal"
)

func main() {

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
			w.Write([]byte(strconv.FormatFloat(res, 'f', 2, 64)))
			return
		}
	})
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
