// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	Name         string `json:"name"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type DetailReq struct {
	AccessToken string `header:"accessToken"`
}

type RefreshReq struct {
	RefreshToken string `from:"RefreshToken"`
}

type Detail struct {
	Name string   `json:"Name"`
	Type string   `json:"Type"`
	Size float64  `json:"Size"`
	Id   string   `json:"Id"`
	Path []Detail `json:"Path"`
}

type DetailReply struct {
	File []Detail `json:"File"`
}

type DataReq struct {
	AccessToken string `header:"accessToken"`
	Id          string `form:"Id"`
}

type CreateDirReq struct {
	AccessToken string `header:"accessToken"`
	Id          string `form:"Id"`
	DirName     string `form:"DirName"`
}

type RegisterReq struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	VerifyId   string `json:"VerifyId"`
	VerifyCode string `json:"VerifyCode"`
}

type VerifyReq struct {
	Mail string `json:"Mail"`
}

type VerifyReply struct {
	VerifyId string `json:"VerifyId"`
}

type CreateReq struct {
	AccessToken string `header:"accessToken"`
	Id          string `path:"Id"`
}

type ShareReply struct {
	Url string `json:"url"`
}

type ShareReq struct {
	AccessToken string `header:"accessToken"`
	Password    string `json:"password"`
	Id          string `path:"Id"`
}

type DataReply struct {
	Data []byte
}

type CreateDataReq struct {
	AccessToken string `header:"accessToken"`
	Name        string `json:"Name"`
	Hash        string `json:"hash"`
	CloudPathId string `from:"CloudPathId"` //云盘路径id
	Size        []int  `json:"Size"`
}

type CreateDataReply struct {
	DataId string `json:"DataId"`
}

type UpleadReq struct {
	AccessToken string `header:"accessToken"`
	DataId      string `form:"DataId"`
	Data        []byte
}
