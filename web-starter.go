package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloFunction)
	http.ListenAndServe(":8080", nil)

}

func helloFunction(wr http.ResponseWriter, req *http.Request){
	fmt.Fprint(wr, `Who dey zuzu!!!`)
}
