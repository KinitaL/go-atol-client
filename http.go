package atol_client

import "net/http"

type (
	ATOLHttpClient struct {
		http.Client
		Token string
	}
)

func NewATOLHttpClient() ATOLClient {
	return &ATOLHttpClient{}
}

func (client *ATOLHttpClient) PostReceipt() {

}

func (client *ATOLHttpClient) Callback() {

}

func (client *ATOLHttpClient) PutReceipt() {

}

func (client *ATOLHttpClient) GetReceipt() {

}
