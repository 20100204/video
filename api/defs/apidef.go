package defs

type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string

}

type Comments struct {
	Id string
	VideoId string
	Author string
	Content string
}
type SimpleSession struct {
	UserName string
	TTL int64
}

type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}