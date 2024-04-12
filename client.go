package atol_client

type ATOLClient interface {
	PostReceipt()
	Callback()
	PutReceipt()
	GetReceipt()
}
