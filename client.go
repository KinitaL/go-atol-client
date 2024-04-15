package atol_client

import "context"

type ATOLClient interface {
	PostReceipt()
	Callback()
	PutReceipt()
	GetReceipt(ctx context.Context, request *GetReceiptRequestMessage) (*ReceiptMessage, error)
}
