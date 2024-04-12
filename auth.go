package atol_client

type (
	AuthRequestMessage struct {
		Login string `json:"login"`
		Pass  string `pass:"pass"`
	}

	AuthResponseMessage struct {
		Error     *AuthError `json:"error"`
		Token     string     `json:"token"`
		Timestamp string     `json:"timestamp"`
	}

	AuthError struct {
		ErrorID string `json:"error_id"`
		Code    int64  `json:"code"`
		Text    string `json:"text"`
		Type    string `json:"type"`
	}
)
