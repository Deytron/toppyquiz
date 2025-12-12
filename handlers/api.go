package handlers

import (
	"fmt"
	"os"

	"resty.dev/v3"
)

func APICall(method string, endpoint string, body any) (any, error) {
	client := resty.New()
	defer client.Close()

	req := client.R().SetAuthToken(os.Getenv("API_TOKEN"))
	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		resp, err = req.Get(endpoint)
	case "POST":
		resp, err = req.SetBody(body).Post(endpoint)
	}
	if os.Getenv("DEBUG") == "true" {
		fmt.Println("Request to:", endpoint)
	}
	return resp.Result(), err
}
