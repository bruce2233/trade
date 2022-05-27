package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetInfo() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
		fmt.Println("get url error")
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
}
