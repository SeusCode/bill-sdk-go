package responses

type (
	AuthServerStatus struct {
		DbServer   string `json:"DbServer"`
		AppServer  string `json:"AppServer"`
		AuthServer string `json:"AuthServer"`
	}

	AuthResponse struct {
		JWT          string           `json:"jwt"`
		ServerStatus AuthServerStatus `json:"server_status"`
	}
)
