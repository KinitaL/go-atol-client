package atol_client

import (
	"bytes"
	"github.com/avast/retry-go/v4"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

const (
	applicationJsonContentType = "application/json"
)

type (
	ATOLHttpClient struct {
		http.Client
		url           string
		login         string
		password      string
		retryAttempts uint
	}
)

func NewATOLHttpClient(url, login, pass string, retryAttempts uint) ATOLClient {
	return &ATOLHttpClient{
		url:           url,
		login:         login,
		password:      pass,
		retryAttempts: retryAttempts,
	}
}

func (client *ATOLHttpClient) PostReceipt() {
	authResp, err := client.auth()
	if err != nil {

	}

}

func (client *ATOLHttpClient) Callback() {
}

func (client *ATOLHttpClient) PutReceipt() {
	authResp, err := client.auth()
	if err != nil {

	}

}

func (client *ATOLHttpClient) GetReceipt() {
	authResp, err := client.auth()
	if err != nil {

	}

}

func (client *ATOLHttpClient) auth() (*authResponseMessage, error) {
	// request body: go struct to array or bytes
	reqBody, err := json.Marshal(&authRequestMessage{
		Login: client.login,
		Pass:  client.password,
	})
	if err != nil {
		return nil, err
	}

	// try request retryAttempts times
	body, err := retry.DoWithData(
		func() ([]byte, error) {
			resp, err := client.Post(client.url, applicationJsonContentType, bytes.NewReader(reqBody))
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			return body, nil
		},
		retry.Attempts(client.retryAttempts),
	)
	if err != nil {
		// handle error
		return nil, err
	}

	// parse response to golang struct
	var resp authResponseMessage
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	// check inner error
	if resp.Error != nil {
		// TODO handle error
	}

	return &resp, nil
}
