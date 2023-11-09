package dto

type BaseResponse struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type MetaResponse struct {
	IsSuccess  bool   `json:"is_success"`
	Message    string `json:"message"`
	StatusCode string `json:"status_code"`
}

// Meta Response Paggination
type MetaResponsePaggination struct {
	IsSuccess  bool   `json:"is_success"`
	Message    string `json:"message"`
	Page       int64  `json:"page"`
	StatusCode string `json:"status_code"`
	TotalItems int64  `json:"total_items"`
	TotalPage  int64  `json:"total_page"`
}
