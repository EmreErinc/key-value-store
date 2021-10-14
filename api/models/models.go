package models

type Response struct {
	Code   int                    `json:"code"`
	Error  bool                   `json:"error"`
	Errors map[string]interface{} `json:"errors"`
	Result interface{}            `json:"result"`
}
