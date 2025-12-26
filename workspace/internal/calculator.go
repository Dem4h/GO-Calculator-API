package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var operators = map[string]string{
	"/": "",
	"+": "",
	"-": "",
	"*": "",
}

func EnableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
}

func ReturnResponse(r *http.Request) (int, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var data map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}
	result, err := parseOperation(data["foo"])
	if err != nil {
		return 0, err
	}
	return result, nil

}
func parseOperation(operation string) (int, error) {
	num1, num2 := "", ""
	op := ""
	for _, k := range operation {
		if _, prs := operators[string(k)]; prs {
			op = string(k)
		} else if op == "" {
			num1 += string(k)
		} else {
			num2 += string(k)
		}
	}
	return calc(num1, num2, op)
}
func calc(sn1, sn2, op string) (int, error) {
	if sn1 == "" || sn2 == "" || op == "" {
		err := errors.New("missing 1 number/operator")
		return 0, err
	}
	n1, err := strconv.Atoi(sn1)
	if err != nil {
		return 0, err
	}
	n2, err := strconv.Atoi(sn2)
	if err != nil {
		return 0, err
	}
	switch op {
	case "/":
		if n2 == 0 {
			return 0, errors.New("Cant divide by 0")
		}
		return n1 / n2, nil
	case "*":
		return n1 * n2, nil
	case "-":
		return n1 - n2, nil

	case "+":
		return n1 + n2, nil
	default:
		return 0, nil
	}
}

func MyFunc() {
	fmt.Println("test")
}
