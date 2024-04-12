package atol_client

type (
	authRequestMessage struct {
		Login string `json:"login"`
		Pass  string `pass:"pass"`
	}

	authResponseMessage struct {
		Error     *authError `json:"error"`
		Token     string     `json:"token"`
		Timestamp string     `json:"timestamp"`
	}

	authError struct {
		ErrorID string `json:"error_id"`
		Code    int64  `json:"code"`
		Text    string `json:"text"`
		Type    string `json:"type"`
	}
)
