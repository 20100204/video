package main

import (
	"awesomeProject/video/api/auth"
	"awesomeProject/video/api/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m.r
}
func (m middleWareHandler) ServerHttp(w http.ResponseWriter, r *http.Request) {
	auth.ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handlers.CreateUser)
	router.POST("/user/:user_name", handlers.Login)
	return router;
}

func main() {
	router := RegisterHandlers()
	mh := NewMiddleWareHandler(router)
	http.ListenAndServe(":8090", mh)
}
