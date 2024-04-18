package atol_client

type ATOLClient interface {
	PostReceipt(request *PostReceiveRequestMessage) (*PostReceiptMessageResponse, error)
	Callback(request *ReceiptMessage, f func(request *ReceiptMessage))
	PutReceipt()
	GetReceipt(request *GetReceiptRequestMessage) (*ReceiptMessage, error)
}
