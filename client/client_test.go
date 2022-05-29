package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetInfo(t *testing.T) {
	GetInfo()
}

func TestAPI(t *testing.T) {
	client := &http.Client{}
	// resp, err := client.Get("http://example.com")
	// ...
	req, _ := http.NewRequest("GET", "https://aws.okx.com/api/v5/market/ticker?instId=BTC-USD-SWAP", nil)
	// ...

	bodyBytes := []byte(`{"instType" : "SPOT"}`)
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	resp, _ := client.Do(req)
	body, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		t.Logf("")
	}
	t.Logf(string(body))
}

func TestData(t *testing.T) {
	client := &http.Client{}
	// resp, err := client.Get("http://example.com")
	// ...
	// baseURL := "https://aws.okx.com/api"
	// // apiURL := "/v5/market/ticker?instId=BTC-USD-SWAP"
	// apiURL := "api/v5/rubik/stat/margin/loan-ratio"

	// parameter := "?ccy=ETH"

	req, _ := http.NewRequest("GET", "https://fapi.binance.com/futures/data/topLongShortPositionRatio?symbol=ETHUSDT&period=5m", nil)
	// ...

	// bodyBytes := []byte(`{"instType" : "SPOT"}`)
	// req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	resp, _ := client.Do(req)
	body, error := ioutil.ReadAll(resp.Body)

	if error != nil {
		t.Logf("")
	}
	t.Logf(string(body))
}
