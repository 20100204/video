package dbops

import (
	"awesomeProject/video/api/defs"
	"awesomeProject/video/api/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func AddUserCredential(LoginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("Insert Into users(login_name,pwd) values (?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(LoginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", nil
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("Delete from users where login_name=? and pwd = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("2006-01-02 15:04:05")
	stmtInt, err := dbConn.Prepare("Insert into video_info (id,author_id,name,display_ctime) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtInt.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	defer stmtInt.Close()
	res := &defs.VideoInfo{vid, aid, name, ctime}
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("select author_id,name,display_ctime from video_info where id = ?")
	if err != nil {
		return nil, err
	}
	var (
		name, display_ctime string
		aid                 int
	)

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &display_ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer stmtOut.Close()
	return &defs.VideoInfo{vid, aid, name, display_ctime}, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("Delete from video_info where id = ?")
	if err != nil {
		return err
	}
	if _, err = stmtDel.Exec(vid); err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil

}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("Insert Into comments (id,video_id,author_id,content) values (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string) ([]*defs.Comments, error) {
	stmtOut, err := dbConn.Prepare("select comments.id,users.login_name,comments.content from comments  left join users on comments.author_id = users.id where comments.video_id= ?")
	if err != nil {
		return nil, err
	}
	var res []*defs.Comments
	rows, err := stmtOut.Query(vid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return nil, err
		}
		temp := &defs.Comments{id, vid, name, content}
		res = append(res, temp)
	}

	return res, nil

}
