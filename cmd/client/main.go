package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	fmt.Println(response)
}
