package auth

type (
	ServerStatus struct {
		DbServer   string `json:"DbServer"`
		AppServer  string `json:"AppServer"`
		AuthServer string `json:"AuthServer"`
	}

	AuthResponse struct {
		JWT          string       `json:"jwt"`
		ServerStatus ServerStatus `json:"server_status"`
	}
)
