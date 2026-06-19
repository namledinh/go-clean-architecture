package common

type httpResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data interface{}, paging interface{}, filter interface{}) *httpResponse {
	return &httpResponse{
		Data: data,
		Paging: paging,
		Filter: filter,
	}
}

func SimpleSuccessResponse(data interface{}) *httpResponse {
	return &httpResponse{
		Data: data,
	}
}