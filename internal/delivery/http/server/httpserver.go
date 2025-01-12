package httpserver

type RequestError struct {
	ErrMsg string `json:"errMsg"`
}

func (r *RequestError) Error() string {
	return r.ErrMsg
}
