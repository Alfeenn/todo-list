package web

type WebResponse struct {
	Code    int         `json:"code,omitempty"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
