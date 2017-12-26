package utils

type Reponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}
type ErrorResponse struct {
	Error  interface{} `json:"error"`
	Status string      `json:"status"`
}
