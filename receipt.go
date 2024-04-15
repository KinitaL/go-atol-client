package atol_client

type Status string
type ErrorSourceType string

const (
	Done Status = "done"
	Fail Status = "fail"
	Wait Status = "wait"

	System  ErrorSourceType = "system"
	Driver  ErrorSourceType = "driver"
	Timeout ErrorSourceType = "timeout"
	Unknown ErrorSourceType = "unknown"
)

type (
	GetReceiptRequestMessage struct {
		GroupCode string
		UUID      string
	}

	ReceiptMessage struct {
		UUID        string                 `json:"uuid"`
		Timestamp   string                 `json:"timestamp"`
		CallbackURL string                 `json:"callback_url"`
		Status      Status                 `json:"status"`
		GroupCode   string                 `json:"group_code"`
		DaemonCode  string                 `json:"daemon_code"`
		DeviceCode  string                 `json:"device_code"`
		ExternalID  string                 `json:"external_id"`
		Error       *ReceiptMessageError   `json:"error"`
		Payload     *ReceiptMessagePayload `json:"payload"`
	}

	ReceiptMessageError struct {
		ErrorID string          `json:"error_id"`
		Code    int             `json:"code"`
		Text    string          `json:"text"`
		Type    ErrorSourceType `json:"type"`
	}

	ReceiptMessagePayload struct {
		FiscalReceiptNumber     int    `json:"fiscal_receipt_number"`
		ShiftNumber             int    `json:"shift_number"`
		ReceiptDatetime         string `json:"receipt_datetime"`
		Total                   int    `json:"total"`
		FNNumber                string `json:"fn_number"`
		ECRRegistrationNumber   string `json:"ecr_registration_number"`
		FiscalDocumentNumber    int    `json:"fiscal_document_number"`
		FiscalDocumentAttribute int    `json:"fiscal_document_attribute"`
		FNSSite                 string `json:"fns_site"`
		OFDINN                  string `json:"ofd_inn"`
		OFDReceiptURL           string `json:"ofd_receipt_url"`
	}

	PostReceiveRequestMessage struct {
		Timestamp  string  `json:"timestamp"`
		Service    Service `json:"service"`
		ExternalID string  `json:"external_id"`
	}

	Service struct {
		CallbackURL string `json:"callback_url"`
	}
)
