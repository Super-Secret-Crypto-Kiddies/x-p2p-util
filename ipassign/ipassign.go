package ipassign

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type APIResponse struct {
	Query string `json:"query"`
}

func GetIPv4Address() string {
	resp, err := http.Get("http://ip-api.com/json/")

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var apiResp APIResponse
	json.Unmarshal(body, &apiResp)

	return apiResp.Query
}

func GetIPv6Address() string {
	resp, err := http.Get("https://v6.ident.me/")

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return (string)(body)

}
