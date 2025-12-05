package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"GIN-APP/internal/config"
)

type Payload struct {
	Uri   string `json:"uri"`
	Token string `json:"token"`
}

func main() {
	config.Load()
	fmt.Println("Server running on port:", config.AppConfig.Port)
	

	http.ListenAndServe(":"+config.AppConfig.Port, router)
	// url := "http://localhost:4100/auth/go-api"
	// payload := []byte(`{"uri":"user/admin/permissions", "token":"asdfasjkdhfnladkfmalsdkfnasdf"}`)

	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	// req.Header.Set("Content-Type", "application/json")
	// if err != nil {
	// 	panic(err)
	// }

	// client := &http.Client{}
	// resp, err := client.Do(req)

	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// body, _ := io.ReadAll(resp.Body)

	// // fmt.Println("Response String:", string(body))

	// var jsonData map[string]any

	// json.Unmarshal(body, &jsonData)

	// fmt.Println("Response:", jsonData["message"])

	// jsonString, _ := json.Marshal(jsonData)

	// fmt.Println("Response:", string(jsonString))
}
