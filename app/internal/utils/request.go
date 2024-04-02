package utils

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const (
	ChefServiceURL = "http://chef-service:8090/chef"
	BBBServiceURL  = "http://bbb-service:8091/bbb"
)

func SendRequest(ctx context.Context, serviceURL string, dishName string) string {
	params := url.Values{}
	params.Add("dish_name", dishName)
	req, err := http.NewRequestWithContext(ctx, "GET", serviceURL+"?"+params.Encode(), nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return ""
	}

	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ReadAll: ", err)
		return ""
	}

	return string(body)
}
