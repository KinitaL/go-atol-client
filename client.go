package atol_client

type ATOLClient interface {
	PostReceipt(request *PostReceiveRequestMessage) (*PostReceiptMessageResponse, error)
	Callback()
	PutReceipt()
	GetReceipt(request *GetReceiptRequestMessage) (*ReceiptMessage, error)
}
