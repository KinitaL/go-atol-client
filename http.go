package atol_client

import (
	"bytes"
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

const (
	applicationJsonContentType = "application/json"
	pathToAuth                 = "getToken"
	pathToGetReceipt           = "report"

	authHeader = "Token"
)

type (
	ATOLHttpClient struct {
		http.Client
		url           string
		login         string
		password      string
		groupCode     string
		apiVersion    string
		retryAttempts uint
	}
)

func NewATOLHttpClient(url, login, pass, groupCode, apiVersion string, retryAttempts uint) ATOLClient {
	return &ATOLHttpClient{
		url:           url,
		login:         login,
		password:      pass,
		groupCode:     groupCode,
		apiVersion:    apiVersion,
		retryAttempts: retryAttempts,
	}
}

func (client *ATOLHttpClient) PostReceipt(request *PostReceiveRequestMessage) (*PostReceiptMessageResponse, error) {
	//authResp, err := client.auth()
	//if err != nil {
	//
	//}
	return nil, nil
}

func (client *ATOLHttpClient) Callback() {
}

func (client *ATOLHttpClient) PutReceipt() {
	//authResp, err := client.auth()
	//if err != nil {
	//
	//}

}

func (client *ATOLHttpClient) GetReceipt(request *GetReceiptRequestMessage) (*ReceiptMessage, error) {
	// authorize
	authResp, err := client.auth()
	if err != nil {
		return nil, NewAuthError(err.Error(), false)
	}

	// try request retryAttempts times
	body, err := retry.DoWithData(
		func() ([]byte, error) {
			req, err := http.NewRequest("GET", fmt.Sprintf(
				"%s/%s/%s/%s/%s",
				client.url,
				client.apiVersion,
				client.groupCode,
				pathToGetReceipt,
				request.UUID,
			), nil)
			req.Header.Add(authHeader, authResp.Token)
			if err != nil {
				return nil, NewParsingError(err.Error(), false)
			}
			resp, err := client.Do(req)
			if err != nil {
				return nil, NewExternalError(err.Error(), false)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, NewParsingError(err.Error(), false)
			}
			return body, nil
		},
		retry.Attempts(client.retryAttempts),
	)
	if err != nil {
		// handle error
		return nil, err
	}

	// parse response
	var resp ReceiptMessage
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, NewJsonError(err.Error(), false)
	}

	return &resp, nil
}

func (client *ATOLHttpClient) auth() (*authResponseMessage, error) {
	// request body: go struct to array or bytes
	reqBody, err := json.Marshal(&authRequestMessage{
		Login: client.login,
		Pass:  client.password,
	})
	if err != nil {
		return nil, NewJsonError(err.Error(), false)
	}

	// try request retryAttempts times
	body, err := retry.DoWithData(
		func() ([]byte, error) {
			resp, err := client.Post(fmt.Sprintf("%s/%s/%s", client.url, client.apiVersion, pathToAuth), applicationJsonContentType, bytes.NewReader(reqBody))
			if err != nil {
				return nil, NewExternalError(err.Error(), false)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, NewParsingError(err.Error(), false)
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
		return nil, NewJsonError(err.Error(), false)
	}

	// check inner error
	if resp.Error != nil {
		return nil, NewExternalError(resp.Error.Text, true)
	}

	return &resp, nil
}
