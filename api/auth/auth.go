package auth

import (
	"awesomeProject/video/api/session"
	"net/http"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FILED_UNAME = "X-User-Name"

func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FILED_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter,r *http.Request) bool  {
	uname := r.Header.Get(HEADER_FILED_UNAME)
	if len(uname)==0{
		return  false
	}
	return true
}
