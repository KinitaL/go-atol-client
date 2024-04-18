package atol_client

// TODO json schema validation
// TODO enum types

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
		FiscalReceiptNumber     int     `json:"fiscal_receipt_number"`
		ShiftNumber             int     `json:"shift_number"`
		ReceiptDatetime         string  `json:"receipt_datetime"`
		Total                   float64 `json:"total"`
		FNNumber                string  `json:"fn_number"`
		ECRRegistrationNumber   string  `json:"ecr_registration_number"`
		FiscalDocumentNumber    int     `json:"fiscal_document_number"`
		FiscalDocumentAttribute int     `json:"fiscal_document_attribute"`
		FNSSite                 string  `json:"fns_site"`
		OFDINN                  string  `json:"ofd_inn"`
		OFDReceiptURL           string  `json:"ofd_receipt_url"`
	}

	PostReceiveRequestMessage struct {
		Operation  Operation   `json:"operation" validate:"required"` // used to build request url
		Timestamp  string      `json:"timestamp" validate:"required"`
		Service    *Service    `json:"service,omitempty"`
		ExternalID string      `json:"external_id" validate:"required"`
		Receipt    *Receipt    `json:"receipt,omitempty"`
		Correction *Correction `json:"correction,omitempty"` // instead of receipt or vice versa
	}

	Service struct {
		CallbackURL string `json:"callback_url"`
	}

	Correction struct {
		Company        *Company        `json:"company"`
		Payments       []Payment       `json:"payments"`
		Vats           []Vat           `json:"vats"`
		DeviceNumber   string          `json:"device_number"`
		Cashier        string          `json:"cashier"`
		CorrectionInfo *CorrectionInfo `json:"correction_info"`
	}

	CorrectionInfo struct {
		Type       string `json:"type"`
		BaseDate   string `json:"base_date"`
		BaseNumber string `json:"base_number"`
	}

	Receipt struct {
		Client               *Client              `json:"client" validate:"required"`
		Company              *Company             `json:"company" validate:"required"`
		AgentInfo            *AgentInfo           `json:"agent_info,omitempty"`
		SupplierInfo         *SupplierInfo        `json:"supplier_info,omitempty"`
		Items                []Item               `json:"items" validate:"required"`
		Payments             []Payment            `json:"payments" validate:"required"`
		Vats                 []Vat                `json:"vats,omitempty"`
		Total                float64              `json:"total" validate:"required"`
		AdditionalCheckProps string               `json:"additional_check_props,omitempty"`
		Cashier              string               `json:"cashier,omitempty"`
		DeviceNumber         string               `json:"device_number,omitempty"`
		AdditionalUserProps  *AdditionalUserProps `json:"additional_user_props,omitempty"`
	}

	Client struct {
		Name  string `json:"name"`
		INN   string `json:"inn"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	}

	Company struct {
		Email          string `json:"email"`
		SNO            string `json:"sno"`
		INN            string `json:"inn" validate:"required"`
		PaymentAddress string `json:"payment_address" validate:"required"`
		Location       string `json:"location"`
	}

	AgentInfo struct {
		Type                    *AgentType               `json:"type,omitempty"`
		PayingAgentInfo         *PayingAgentInfo         `json:"paying_agent,omitempty"`
		ReceivePaymentsOperator *ReceivePaymentsOperator `json:"receive_payments_operator,omitempty"`
		MoneyTransferObject     *MoneyTransferObject     `json:"money_transfer_object,omitempty"`
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
		Phones []string `json:"phones"`
	}

	Item struct {
		Name              string             `json:"name" validate:"required"`
		Price             float64            `json:"price" validate:"required"`
		Quantity          float64            `json:"quantity" validate:"required"`
		Sum               float64            `json:"sum" validate:"required"`
		MeasurementUnit   string             `json:"measurement_unit,omitempty"`
		MeasurementCode   string             `json:"measurement_code,omitempty"`
		PaymentMethod     string             `json:"payment_method,omitempty"`
		PaymentObject     string             `json:"payment_object,omitempty"`
		Vat               *Vat               `json:"vat,omitempty"`
		AgentInfo         *AgentInfo         `json:"agent_info,omitempty"`
		SupplierInnerInfo *SupplierInnerInfo `json:"supplier_info,omitempty"`
		UserData          string             `json:"user_data,omitempty"`
		Excise            int                `json:"excise,omitempty"`
		CountryCode       string             `json:"country_code,omitempty"`
		DeclarationNumber string             `json:"declaration_number,omitempty"`
	}

	Vat struct {
		Type string  `json:"type,omitempty"`
		Sum  float64 `json:"sum,omitempty"`
	}

	SupplierInnerInfo struct {
		Phones []string `json:"phones"`
		Name   string   `json:"name"`
		INN    string   `json:"inn"`
	}

	Payment struct {
		Type int     `json:"type" validate:"required"`
		Sum  float64 `json:"sum" validate:"required"`
	}

	AdditionalUserProps struct {
		Name  string `json:"name" validate:"required"`
		Value string `json:"value" validate:"required"`
	}

	PostReceiptMessageResponse struct {
		UUID      string               `json:"uuid"`
		Timestamp string               `json:"timestamp"`
		Status    Status               `json:"status"`
		Error     *ReceiptMessageError `json:"error"`
	}
)
