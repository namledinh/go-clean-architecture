package model

type Filter struct {
	Status string `json:"status" form:"status"`
	Path   string `json:"path" form:"path"`
	DataType string `json:"data_type" form:"data_type"`
}