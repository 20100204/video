package handlers

import (
	"awesomeProject/video/api/dbops"
	"awesomeProject/video/api/defs"
	"awesomeProject/video/api/response"
	"awesomeProject/video/api/session"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)
func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params )  {
	 res,_ := ioutil.ReadAll(r.Body)
	 ubody := &defs.UserCredential{}
	 println(p)
	 if err := json.Unmarshal(res,ubody);err != nil {
	 	response.SendErrorResponse(w,defs.ErrorRequestBodyParseFaild)
	 	return
	 }
	 if err := dbops.AddUserCredential(ubody.UserName,ubody.Pwd);err != nil {
		response.SendErrorResponse(w,defs.ErrorDbError)
		 return
	 }
	 id := session.GenerateNewSessionId(ubody.UserName)
	 su := defs.SignedUp{Success:true,SessionId:id}
	 if resp,err := json.Marshal(su);err != nil {
			response.SendErrorResponse(w,defs.ErrorInternalFaults)
		 return
	 }else{
		 response.SendNormalResponse(w,string(resp),201)
	 }

}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
		userName := p.ByName("user_name")
		io.WriteString(w,"hello "+ userName)

}
