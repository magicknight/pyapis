package models

type Result struct {
	Code int `json:"code"`
	Desc string `json:"desc"`
	Detail interface{} `json:"detail"`
}
