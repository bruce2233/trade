package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"
)

const (
	url = "https://aws.okx.com"
)

func GetInfo() {
	client := &http.Client{}
	// resp, err := client.Get("http://example.com")
	// ...
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)

	resp, _ := client.Do(req)
	fmt.Println(resp.StatusCode)
}

func Signature(timeStamp string, method string, requestPath string, body string) {
	timeStamp = time.Stamp
	date := timeStamp
	fmt.Println(date)
	method = "POST"
	secret := "B5B2C8F36FE229F42E9C8B2463E84595"
	requestPath = "/api/v5/account/balance"
	body = ""
	final := timeStamp + method + requestPath + body
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(final))
}
