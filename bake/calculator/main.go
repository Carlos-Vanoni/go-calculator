package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/calc/{op}/{number1}/{number2}", calculate)
	r.HandleFunc("/calc/history", calcHistory)
	http.ListenAndServe(":8080", r)
}

func calculate(w http.ResponseWriter, r *http.Request) {
	var result string
	vars := mux.Vars(r)
	num1, _ := strconv.Atoi(vars["number1"])
	num2, _ := strconv.Atoi(vars["number2"])

	if num1 == 0 && vars["number1"] != "0" {
		result = "first parameter is not a number, try again"
	} else if num2 == 0 && vars["number2"] != "0" {
		result = "second is not a number, try again"
	} else {
		calculator := calculator{num1, num2}
		switch vars["op"] {
		case "sum":
			result = calculator.sum()
		case "sub":
			result = calculator.sub()
		case "mul":
			result = calculator.mul()
		case "div":
			if num2 == 0 {
				result = "Error, cannot be divided by 0"
			} else {
				result = calculator.div()
			}
		default:
			fmt.Fprintf(w, "There was an error in the operation. try again.")
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)
}

func calcHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "History: %v\n", getHistory())
}
