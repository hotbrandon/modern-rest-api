package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Payload struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func main() {
	payload := Payload{
		UserName: "jack",
		Password: "1234567",
	}

	// Marshal struct to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Make POST request
	response, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		data, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Response:", string(data))
		fmt.Println("Status OK")
	} else {
		fmt.Println(response.StatusCode)
	}

}
