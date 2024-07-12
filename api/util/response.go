package util

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func Success(data any) *Response {
	res := Response{
		Code:    0,
		Message: "操作成功",
		Result:  data,
	}
	return &res
}

func Abort(msg string) *Response {
	res := Response{
		Code:    422,
		Message: msg,
		Result:  nil,
	}
	return &res
}
func Failed(msg string) *Response {
	res := Response{
		Code:    500,
		Message: msg,
		Result:  nil,
	}
	return &res
}
