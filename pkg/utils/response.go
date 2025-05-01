package utils

type Pagination struct {
	Total   int `json:"total"`
	Page    int `json:"page"`
	Limit   int `json:"limit"`
	Maxpage int `json:"max_page"`
	Offset  int `json:"offset,omitempty"`
}

type Meta struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type Response[T any] struct {
	Success    bool        `json:"success"`
	Meta       *Meta       `json:"meta"`
	Error      any         `json:"error,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       T           `json:"data,omitempty"`
}

func ResponseError[T any](message string, response *Response[T]) Response[T] {
	res := Response[T]{
		Success: false,
		Meta: &Meta{
			Code:    "99",
			Message: message,
		},
	}
	if response != nil {
		res.Data = response.Data
		res.Error = response.Error
		res.Pagination = response.Pagination

	}

	return res
}
func ResponseSuccess[T any](message string, response *Response[T]) Response[T] {
	res := Response[T]{
		Success: true,
		Meta: &Meta{
			Code:    "00",
			Message: message,
		},
	}

	if response != nil {
		res.Data = response.Data
		res.Pagination = response.Pagination
		res.Error = response.Error
	}
	return res
}
