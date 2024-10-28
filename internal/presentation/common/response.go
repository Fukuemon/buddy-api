package common

type ErrorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Msg         string `json:"msg"`
}
