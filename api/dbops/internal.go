package dbops

import (
	"awesomeProject/video/api/defs"
	"database/sql"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstring := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("Insert Into sessions(session_id,ttl,login_name) values (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstring, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select login_name,ttl from sessions where session_id=?")
	if err != nil {
		return ss, err
	}
	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&uname, &ttl)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.UserName = uname
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil

}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("select * from sessions")
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			break;
		}
		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			ss := &defs.SimpleSession{UserName: login_name, TTL: ttl}
			m.Store(id, ss)

		}
	}
	return m, nil
}

func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("delete from sessions where session_id=?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(sid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}
