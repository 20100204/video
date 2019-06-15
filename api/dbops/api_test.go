package dbops

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"

	)

// init(dblogin,truncate tables)->run tests ->clear data(truncate tables)
var tempVid string
func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkflow(t *testing.T) {
	t.Run("Add",testAddUser)
	t.Run("Get",testGetUser)
	t.Run("Del",testDeleteUser)
	t.Run("Reget",testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("chj", "123")
	if err != nil {
		t.Errorf("error of add user %s:", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("chj")
	if err != nil {
		t.Errorf("Error of getUser %s:", err)
	}
	if pwd ==""{
		t.Errorf("use is not exists")
	}
}
func testDeleteUser(t *testing.T) {
	err := DeleteUser("chj", "123")
	if err != nil {
		t.Errorf("error of deleteUser %s:", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("chj")
	if err != nil {
		t.Errorf("error of regetuser %s:", err)
	}
	if pwd != "" {
		t.Log("pwd:", pwd)
	} else {
		t.Log("user is not exist")
	}
}

func TestVideoInfoWorkFlow(t *testing.T)  {
	//clearTables()
	t.Run("Add",testAddVideoInfo)
	t.Run("GET",testGetVideoInfo)
	//t.Run("DEL",testDeleteVideoInfo)

}

func testAddVideoInfo(t *testing.T)  {
	vd,err := AddNewVideo(1,"first_video")
	if err != nil {
		t.Errorf("error of addvideoinfo:%v",err)
	}
	tempVid = vd.Id
}
func testGetVideoInfo(t *testing.T) {
	_,err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of getVideoinfo:%v",err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
		err := DeleteVideoInfo(tempVid)
		if err != nil {
			t.Errorf("Error of deleteVideoInfo:%v",err)
		}
}
