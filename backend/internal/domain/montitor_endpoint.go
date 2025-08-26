package domain

type MonitorEndpoint struct {
	ID          string `json:"id"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}
