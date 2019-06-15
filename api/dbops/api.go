package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCredential(LoginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("Insert Into users(login_name,pwd) values (?,?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(LoginName, pwd)
	stmtIns.Close()
	return nil
}
func GetUserCredential(loginName string) (string, error) {
    stmtOut,err := dbConn.Prepare("select pwd from users where login_name=?")
    if err != nil {
    	log.Printf("%s",err)
    	return "" ,nil
	}
    var pwd string
    stmtOut.QueryRow(loginName).Scan(&pwd)
    stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string,pwd string) error  {
	stmtDel, err := dbConn.Prepare("Delete from users where login_name=? and pwd = ?")
	if err != nil {
		log.Printf("%s",err)
		return err
	}
	stmtDel.Exec(loginName,pwd)
	stmtDel.Close()
	return nil

}
