package request

type SysLoginReq struct {
	Username   string `json:"username" form:"username" v:"username@required|#用户名不能为空"`
	Password   string `json:"password" form:"password" v:"password@required|#密码不能为空"`
	VerifyCode string `json:"verifyCode" form:"verifyCode"`
	VerifyKey  string `json:"verifyKey" form:"verifyKey"`
}
