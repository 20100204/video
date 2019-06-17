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
	ErrorDbError = ErrorResponse{500,Err{"Db ops  faild","003"}}
	ErrorInternalFaults = ErrorResponse{500,Err{"internal service error","004"}}

)