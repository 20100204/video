package session

import (
	"awesomeProject/video/api/dbops"
	"awesomeProject/video/api/defs"
	"awesomeProject/video/api/utils"
	"log"
	"sync"
	"time"
)



var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDb()  {
     r,err := dbops.RetrieveAllSessions()
     if err != nil {
     	log.Printf("err:",err)
	 }
     r.Range(func(key, v interface{}) bool {
		 ss := v.(*defs.SimpleSession)
		 sessionMap.Store(key,ss)
		 return true
	 })

}

func GenerateNewSessionId(UserName string) string  {
    id,_ := utils.NewUUID()
    ct := time.Now().UnixNano()/1000000
    ttl := ct +30*60*1000
    ss := &defs.SimpleSession{UserName:UserName,TTL:ttl}
    sessionMap.Store(id,ss)
    dbops.InsertSession(id,ttl,UserName)
    return  id
}

func deleteExpiredSession(sid string)  {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
func IsSessionExpired(sid string) (string,bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := time.Now().UnixNano() / 1000000
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		} else {
			return ss.(*defs.SimpleSession).UserName, false
		}
	}
	return "",false
}
