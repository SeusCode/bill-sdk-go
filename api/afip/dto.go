package afip

type (
	ServerStatus struct {
		DbServer   string `json:"DbServer"`
		AppServer  string `json:"AppServer"`
		AuthServer string `json:"AuthServer"`
	}

	ServerStatusResponse struct {
		ServerStatus ServerStatus `json:"server_status"`
	}

	PingResponse struct {
		Datetime  string `json:"datetime"`
		Timestamp int64  `json:"timestamp"`
	}
)
