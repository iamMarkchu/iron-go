package error

import "net/http"

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ZeroErrorHandler(err error) (int, interface{}) {
	return http.StatusBadRequest, Resp{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
}
