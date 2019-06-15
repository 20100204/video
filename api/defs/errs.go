package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSc int
	Error Err
}

var (
	ErrorRequestBodyParseFaild = ErrorResponse{400,Err{"Request body is not correct","001"}}
	ErrorNotAuthUser = ErrorResponse{401,Err{"user authentication faild","002"}}

)