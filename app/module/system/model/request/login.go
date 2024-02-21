package request

type SysLoginReq struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	VerifyCode string `json:"verifyCode"`
	VerifyKey  string `json:"verifyKey"`
}
