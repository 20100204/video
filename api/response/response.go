package response

import (
	"awesomeProject/video/api/defs"
	"encoding/json"
	"io"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, msg defs.ErrorResponse) {
	w.WriteHeader(msg.HttpSc)
	res, _ := json.Marshal(&msg.Error)
	io.WriteString(w, string(res))
}

func SendNormalResponse(w http.ResponseWriter, res string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w,res)
}
