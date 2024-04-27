package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const (
	ChefServiceHost = "chef-service:8090"
	BBProdHost      = "bb-productions:8091"
)

func SendRequest(ctx context.Context, url string) string {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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

func GetChefServiceURL(dishName string) string {
	return fmt.Sprintf("http://%s/%s", ChefServiceHost, dishName)
}

func GetBBBProdURL(dishName string) string {
	params := url.Values{}
	params.Add("dish_name", dishName)
	return fmt.Sprintf("http://%s/rate?%s", BBProdHost, params.Encode())
}
