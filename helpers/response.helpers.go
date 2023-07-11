package helpers

type Response struct {
	Status     string      `json:"status"`
	HttpStatus int         `json:"http_status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
