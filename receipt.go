package atol_client

type Status string
type ErrorSourceType string
type AgentType string
type Operation string

const (
	Done Status = "done"
	Fail Status = "fail"
	Wait Status = "wait"

	System  ErrorSourceType = "system"
	Driver  ErrorSourceType = "driver"
	Timeout ErrorSourceType = "timeout"
	Unknown ErrorSourceType = "unknown"

	BankPayingAgent    AgentType = "bank_paying_agent"
	BankPayingSubagent AgentType = "bank_paying_subagent"
	PayingAgent        AgentType = "paying_agent"
	PayingSubagent     AgentType = "paying_subagent"
	Attorney           AgentType = "attorney"
	CommissionAgent    AgentType = "commission_agent"
	Another            AgentType = "another"

	Sell           Operation = "sell"
	SellRefund     Operation = "sell_refund"
	SellCorrection Operation = "sell_correction"
	BuyRefund      Operation = "buy_refund"
	Buy            Operation = "buy"
	BuyCorrection  Operation = "buy_correction"
)

type (
	GetReceiptRequestMessage struct {
		UUID string
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
		Operation  Operation  `json:"operation"` // used to build request url
		Timestamp  string     `json:"timestamp"`
		Service    Service    `json:"service"`
		ExternalID string     `json:"external_id"`
		Receipt    Receipt    `json:"receipt"`
		Correction Correction `json:"correction"` // instead of receipt or vice versa
	}

	Service struct {
		CallbackURL string `json:"callback_url"`
	}

	Correction struct {
		Company        Company        `json:"company"`
		Payments       []Payment      `json:"payments"`
		Vats           []Vat          `json:"vats"`
		DeviceNumber   string         `json:"device_number"`
		Cashier        string         `json:"cashier"`
		CorrectionInfo CorrectionInfo `json:"correction_info"`
	}

	CorrectionInfo struct {
		Type       string `json:"type"`
		BaseDate   string `json:"base_date"`
		BaseNumber string `json:"base_number"`
	}

	Receipt struct {
		Client               Client              `json:"client"`
		Company              Company             `json:"company"`
		AgentInfo            AgentInfo           `json:"agent_info"`
		SupplierInfo         SupplierInfo        `json:"supplier_info"`
		Total                int                 `json:"total"`
		AdditionalCheckProps string              `json:"additional_check_props"`
		Cashier              string              `json:"cashier"`
		DeviceNumber         string              `json:"device_number"`
		AdditionalUserProps  AdditionalUserProps `json:"additional_user_props"`
	}

	Client struct {
		Name  string `json:"name"`
		INN   string `json:"inn"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	}

	Company struct {
		Email          string `json:"email"`
		SNO            string `json:"SNO"`
		INN            string `json:"INN"`
		PaymentAddress string `json:"payment_address"`
		Location       string `json:"location"`
	}

	AgentInfo struct {
		Type                    AgentType               `json:"type"`
		PayingAgentInfo         PayingAgentInfo         `json:"paying_agent"`
		ReceivePaymentsOperator ReceivePaymentsOperator `json:"receive_payments_operator"`
		MoneyTransferObject     MoneyTransferObject     `json:"money_transfer_object"`
	}

	PayingAgentInfo struct {
		Operation string   `json:"operation"`
		Phones    []string `json:"phones"`
	}

	ReceivePaymentsOperator struct {
		Phones []string `json:"phones"`
	}

	MoneyTransferObject struct {
		Phones  []string `json:"phones"`
		Name    string   `json:"name"`
		INN     string   `json:"inn"`
		Address string   `json:"address"`
	}

	SupplierInfo struct {
		Phones   []string     `json:"phones"`
		Items    []SupplyItem `json:"items"`
		Payments []Payment    `json:"payments"`
		Vats     []Vat        `json:"vats"`
	}

	SupplyItem struct {
		Name              string            `json:"name"`
		Price             int               `json:"price"`
		Quantity          int               `json:"quantity"`
		Sum               int               `json:"sum"`
		MeasurementUnit   string            `json:"measurement_unit"`
		MeasurementCode   string            `json:"measurement_code"`
		PaymentMethod     string            `json:"payment_method"`
		PaymentObject     string            `json:"payment_object"`
		Vat               Vat               `json:"vat"`
		AgentInfo         AgentInfo         `json:"agent_info"`
		SupplierInnerInfo SupplierInnerInfo `json:"supplier_info"`
		UserData          string            `json:"user_data"`
		Excise            int               `json:"excise"`
		CountryCode       string            `json:"country_code"`
		DeclarationNumber string            `json:"declaration_number"`
	}

	Vat struct {
		Type string `json:"type"`
		Sum  int    `json:"sum"`
	}

	SupplierInnerInfo struct {
		Phones []string `json:"phones"`
		Name   string   `json:"name"`
		INN    string   `json:"inn"`
	}

	Payment struct {
		Type int `json:"type"`
		Sum  int `json:"sum"`
	}

	AdditionalUserProps struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	PostReceiptMessageResponse struct {
		UUID      string               `json:"uuid"`
		Timestamp string               `json:"timestamp"`
		Status    Status               `json:"status"`
		Error     *ReceiptMessageError `json:"error"`
	}
)
